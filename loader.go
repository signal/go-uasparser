package uas

import (
	"encoding/xml"
	"github.com/hashicorp/golang-lru"
	"io"
	"os"
	"regexp"
)

var regMatcher *regexp.Regexp

func init() {
	regMatcher = regexp.MustCompile("^/(?P<reg>.*)/(?P<flags>[imsU]*)\\s*$")
}

func compileReg(reg string) *regexp.Regexp {
	return regexp.MustCompile(regMatcher.ReplaceAllString(reg, "(?${flags}:${reg})"))
}

func compileBrowserRegs(regs []*BrowserReg) {
	for i, reg := range regs {
		regs[i].Reg = compileReg(reg.RegString)
	}
}

func compileOsRegs(regs []*OsReg) {
	for i, reg := range regs {
		regs[i].Reg = compileReg(reg.RegString)
	}
}

func compileDeviceRegs(regs []*DeviceReg) {
	for i, reg := range regs {
		regs[i].Reg = compileReg(reg.RegString)
	}
}

func mapBrowserTypeToBrowser(manifest *Manifest) {
	for _, browser := range manifest.Data.Browsers {
		browserType, found := manifest.GetBrowserType(browser.TypeId)
		if !found {
			browserType = manifest.otherBrowserType
		}
		browser.Type = browserType
	}
}

func mapOsToBrowser(manifest *Manifest) {
	for _, browser := range manifest.Data.Browsers {
		browser.Os, _ = manifest.GetOsForBrowser(browser.Id)
	}
}

func initOtherUnknown(manifest *Manifest) {
	manifest.otherBrowserType, _ = manifest.FindBrowserTypeByName(OtherBrowserTypeName)
	manifest.unknownOs, _ = manifest.FindOsByName(UnknownOsName)
	manifest.otherDevice, _ = manifest.FindDeviceByName(OtherDeviceName)
	manifest.unknownBrowser = &Browser{
		entity: entity{
			Name: UnknownBrowserName,
		},
	}
	manifest.unknownBrowserVersion = &BrowserVersion{
		Browser: manifest.unknownBrowser,
		Version: "",
	}
}

// Creates a new Manifest instance by processing the XML from the provided Reader.
func Load(reader io.Reader) (*Manifest, error) {
	cache, err := lru.New(5000) // TODO make configurable
	if err != nil {
		return nil, err
	}

	manifest := &Manifest{cache: cache}
	if err := xml.NewDecoder(reader).Decode(manifest); err != nil {
		return nil, err
	}
	initOtherUnknown(manifest)
	compileBrowserRegs(manifest.Data.BrowsersReg)
	compileOsRegs(manifest.Data.OperatingSystemsReg)
	compileDeviceRegs(manifest.Data.DevicesReg)
	mapBrowserTypeToBrowser(manifest)
	mapOsToBrowser(manifest)

	return manifest, nil
}

// Creates a new Manifest instance by processing the XML from the provided file.
func LoadFile(path string) (*Manifest, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Load(file)
}

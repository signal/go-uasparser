package uas

import (
	"encoding/xml"
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

func compileBrowserRegs(data *Data) {
	regs := data.BrowsersReg
	for i, reg := range regs {
		regs[i].Reg = compileReg(reg.RegString)
	}
}

func compileOsRegs(data *Data) {
	regs := data.OperatingSystemsReg
	for i, reg := range regs {
		regs[i].Reg = compileReg(reg.RegString)
	}
}

func compileDeviceRegs(data *Data) {
	regs := data.DevicesReg
	for i, reg := range regs {
		regs[i].Reg = compileReg(reg.RegString)
	}
}

func mapBrowserTypeToBrowser(manifest *Manifest) {
	for _, browser := range manifest.Data.Browsers {
		browserType, found := manifest.GetBrowserType(browser.TypeId)
		if !found {
			browserType = manifest.OtherBrowserType()
		}
		browser.Type = browserType
	}
}

func mapOsToBrowser(manifest *Manifest) {
	for _, browser := range manifest.Data.Browsers {
		browser.Os, _ = manifest.GetOsForBrowser(browser.Id)
	}
}

// Creates a new Manifest instance by processing the XML from the provided Reader.
func Load(reader io.Reader) (*Manifest, error) {
	manifest := &Manifest{}
	if err := xml.NewDecoder(reader).Decode(manifest); err != nil {
		return nil, err
	}
	compileBrowserRegs(manifest.Data)
	compileOsRegs(manifest.Data)
	compileDeviceRegs(manifest.Data)
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

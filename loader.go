// Loads up a UAS manifest from an XML file.
package uas

import (
	"encoding/xml"
	"io"
	"os"
	"regexp"
)

var regMatcher *regexp.Regexp

func init() {
	regMatcher = regexp.MustCompile("^/(?P<reg>.*)/(?P<flags>[imsU]*)$")
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

func Load(reader io.Reader) (*Manifest, error) {
	manifest := &Manifest{}
	if err := xml.NewDecoder(reader).Decode(manifest); err != nil {
		return nil, err
	}
	compileBrowserRegs(manifest.Data)
	compileOsRegs(manifest.Data)
	compileDeviceRegs(manifest.Data)
	return manifest, nil
}

func LoadFile(path string) (*Manifest, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Load(file)
}

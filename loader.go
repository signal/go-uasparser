// Loads up a UAS manifest from an XML file.
package uas

import (
  "encoding/xml"
  "io"
  "regexp"
)

var regMatcher *regexp.Regexp

func init() {
  regMatcher = regexp.MustCompile("^/(?P<reg>.*)/(?P<flags>[imsU]*)$")
}

func compileBrowserRegs(manifest *Manifest) {
  for i, reg := range manifest.Data.BrowsersReg {
    manifest.Data.BrowsersReg[i].Reg = regexp.MustCompile(
      regMatcher.ReplaceAllString(reg.RegString, "(?${flags}:${reg})"))
  }
}

func Load(reader io.Reader) (*Manifest, error) {
  var manifest Manifest
  if err := xml.NewDecoder(reader).Decode(&manifest); err != nil {
    return nil, err
  }
  compileBrowserRegs(&manifest)
  return &manifest, nil
}

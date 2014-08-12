// Loads up a UAS manifest from an XML file.
package uas

import (
  "encoding/xml"
  "io"
)

func Load(reader io.Reader) (*Manifest, error) {
  var manifest Manifest
  if err := xml.NewDecoder(reader).Decode(&manifest); err != nil {
    return nil, err
  }
  return &manifest, nil
}

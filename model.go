// Defines the UAS model.
package uas

import (
  "encoding/xml"
  "regexp"
)

type Entity struct {
  Id         int    `xml:"id"`
  Name       string `xml:"name"`
  Company    string `xml:"company"`
  URLCompany string `xml:"url_company"`
  Icon       string `xml:"icon"`
}
type Robot struct {
  Entity
  Family    string `xml:"family"`
  UserAgent string `xml:"useragent"`
  InfoURL   string `xml:"bot_info_url"`
}
type Os struct {
  Entity
  Family  string `xml:"family"`
  URL     string `xml:"url"`
  InfoURL string `xml:"os_info_url"`
}
type Browser struct {
  Entity
  Type    int    `xml:"type"`
  URL     string `xml:"url"`
  InfoURL string `xml:"browser_info_url"`
}
type Device struct {
  Id      int    `xml:"id"`
  Name    string `xml:"name"`
  Icon    string `xml:"icon"`
  InfoURL string `xml:"device_info_url"`
}
type BrowserType struct {
  Id       int    `xml:"id"`
  TypeName string `xml:"type"`
}
type BrowserOs struct {
  BrowserId int `xml:"browser_id"`
  OsId      int `xml:"os_id"`
}
type RegEntity struct {
  Order     int    `xml:"order"`
  RegString string `xml:"regstring"`
  Reg       *regexp.Regexp
}
type BrowserReg struct {
  RegEntity
  BrowserId int `xml:"browser_id"`
}
type OsReg struct {
  RegEntity
  OsId int `xml:"os_id"`
}
type DeviceReg struct {
  RegEntity
  DeviceId int `xml:"device_id"`
}
type Data struct {
  XMLName             xml.Name      `xml:"data"`
  Robots              []Robot       `xml:"robots>robot"`
  OperatingSystems    []Os          `xml:"operating_systems>os"`
  Browsers            []Browser     `xml:"browsers>browser"`
  BrowserTypes        []BrowserType `xml:"browser_types>browser_type"`
  BrowsersReg         []BrowserReg  `xml:"browsers_reg>browser_reg"`
  BrowsersOs          []BrowserOs   `xml:"browsers_os>browser_os"`
  OperatingSystemsReg []OsReg       `xml:"operating_systems_reg>operating_system_reg"`
  Devices             []Device      `xml:"devices>device"`
  DevicesReg          []DeviceReg   `xml:"devices_reg>device_reg"`
}

type Checksum struct {
  Type string `xml:"type,attr"`
  URL  string `xml:",innerxml"`
}
type Description struct {
  XMLName   xml.Name   `xml:"description"`
  Label     string     `xml:"label"`
  Version   string     `xml:"version"`
  Checksums []Checksum `xml:"checksum"`
}

type Manifest struct {
  XMLName     xml.Name `xml:"uasdata"`
  Description Description
  Data        Data
}

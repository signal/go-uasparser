// Defines the UAS model.
package uas

import (
  "encoding/xml"
  "regexp"
)

const (
  OtherBrowserTypeName = "Other"
  UnknownOsName        = "unknown"
  OtherDeviceName      = "Other"
)

type regEntity struct {
  Order     int    `xml:"order"`
  RegString string `xml:"regstring"`
  Reg       *regexp.Regexp
}
type entity struct {
  Id         int    `xml:"id"`
  Name       string `xml:"name"`
  Company    string `xml:"company"`
  URLCompany string `xml:"url_company"`
  Icon       string `xml:"icon"`
}

type Robot struct {
  entity
  Family    string `xml:"family"`
  UserAgent string `xml:"useragent"`
  InfoURL   string `xml:"bot_info_url"`
}
type Os struct {
  entity
  Family  string `xml:"family"`
  URL     string `xml:"url"`
  InfoURL string `xml:"os_info_url"`
}
type OsReg struct {
  regEntity
  OsId int `xml:"os_id"`
}
type Browser struct {
  entity
  TypeId  int    `xml:"type"`
  URL     string `xml:"url"`
  InfoURL string `xml:"browser_info_url"`
}
type BrowserType struct {
  Id   int    `xml:"id"`
  Name string `xml:"type"` // a purposeful departure from the UAS naming
}
type BrowserOs struct {
  BrowserId int `xml:"browser_id"`
  OsId      int `xml:"os_id"`
}
type BrowserReg struct {
  regEntity
  BrowserId int `xml:"browser_id"`
}
type Device struct {
  Id      int    `xml:"id"`
  Name    string `xml:"name"`
  Icon    string `xml:"icon"`
  InfoURL string `xml:"device_info_url"`
}
type DeviceReg struct {
  regEntity
  DeviceId int `xml:"device_id"`
}
type Data struct {
  XMLName             xml.Name       `xml:"data"`
  Robots              []*Robot       `xml:"robots>robot"`
  OperatingSystems    []*Os          `xml:"operating_systems>os"`
  Browsers            []*Browser     `xml:"browsers>browser"`
  BrowserTypes        []*BrowserType `xml:"browser_types>browser_type"`
  BrowsersReg         []*BrowserReg  `xml:"browsers_reg>browser_reg"`
  BrowsersOs          []*BrowserOs   `xml:"browsers_os>browser_os"`
  OperatingSystemsReg []*OsReg       `xml:"operating_systems_reg>operating_system_reg"`
  Devices             []*Device      `xml:"devices>device"`
  DevicesReg          []*DeviceReg   `xml:"devices_reg>device_reg"`
}

type Checksum struct {
  Type string `xml:"type,attr"`
  URL  string `xml:",innerxml"`
}
type Description struct {
  XMLName   xml.Name    `xml:"description"`
  Label     string      `xml:"label"`
  Version   string      `xml:"version"`
  Checksums []*Checksum `xml:"checksum"`
}

type Manifest struct {
  XMLName     xml.Name `xml:"uasdata"`
  Description *Description
  Data        *Data

  // for memoization
  otherBrowserType *BrowserType
  unknownOs        *Os
  otherDevice      *Device
}

type Agent struct {
  String  string
  Type    string
  Browser *Browser
  Os      *Os
  Device  *Device
}

func (self *Manifest) FindRobot(ua string) (*Robot, bool) {
  for _, robot := range self.Data.Robots {
    if ua == robot.UserAgent {
      return robot, true
    }
  }
  return nil, false
}

func (self *Manifest) IsRobot(ua string) bool {
  _, found := self.FindRobot(ua)
  return found
}

func (self *Manifest) GetBrowser(id int) (*Browser, bool) {
  for _, browser := range self.Data.Browsers {
    if id == browser.Id {
      return browser, true
    }
  }
  return nil, false
}

func (self *Manifest) GetBrowserType(id int) (*BrowserType, bool) {
  for _, browserType := range self.Data.BrowserTypes {
    if id == browserType.Id {
      return browserType, true
    }
  }
  return nil, false
}

func (self *Manifest) GetOs(id int) (*Os, bool) {
  for _, os := range self.Data.OperatingSystems {
    if id == os.Id {
      return os, true
    }
  }
  return nil, false
}

func (self *Manifest) GetOsForBrowser(id int) (*Os, bool) {
  for _, browserOs := range self.Data.BrowsersOs {
    if id == browserOs.BrowserId {
      return self.GetOs(browserOs.OsId)
    }
  }
  return nil, false
}

func (self *Manifest) GetDevice(id int) (*Device, bool) {
  for _, device := range self.Data.Devices {
    if id == device.Id {
      return device, true
    }
  }
  return nil, false
}

func (self *Manifest) FindBrowserTypeByName(name string) (*BrowserType, bool) {
  for _, os := range self.Data.BrowserTypes {
    if name == os.Name {
      return os, true
    }
  }
  return nil, false
}

func (self *Manifest) FindOsByName(name string) (*Os, bool) {
  for _, os := range self.Data.OperatingSystems {
    if name == os.Name {
      return os, true
    }
  }
  return nil, false
}

func (self *Manifest) FindDeviceByName(name string) (*Device, bool) {
  for _, device := range self.Data.Devices {
    if name == device.Name {
      return device, true
    }
  }
  return nil, false
}

func (self *Manifest) OtherBrowserType() *BrowserType {
  if self.otherBrowserType == nil {
    self.otherBrowserType, _ = self.FindBrowserTypeByName(OtherBrowserTypeName)
  }
  return self.otherBrowserType
}

func (self *Manifest) UnknownOs() *Os {
  if self.unknownOs == nil {
    self.unknownOs, _ = self.FindOsByName(UnknownOsName)
  }
  return self.unknownOs
}

func (self *Manifest) OtherDevice() *Device {
  if self.otherDevice == nil {
    self.otherDevice, _ = self.FindDeviceByName(OtherDeviceName)
  }
  return self.otherDevice
}

func (self *Manifest) parseBrowserFromUserAgent(ua string) *Browser {
  for _, reg := range self.Data.BrowsersReg {
    if reg.Reg.MatchString(ua) {
      browser, _ := self.GetBrowser(reg.BrowserId)
      return browser
    }
  }
  return nil
}

func (self *Manifest) parseOsFromUserAgent(ua string) *Os {
  for _, reg := range self.Data.OperatingSystemsReg {
    if reg.Reg.MatchString(ua) {
      os, _ := self.GetOs(reg.OsId)
      return os
    }
  }
  return self.UnknownOs()
}

func (self *Manifest) parseDeviceFromUserAgent(ua string) *Device {
  for _, reg := range self.Data.DevicesReg {
    if reg.Reg.MatchString(ua) {
      device, _ := self.GetDevice(reg.DeviceId)
      return device
    }
  }
  return self.OtherDevice()
}

func (self *Manifest) ParseBrowser(ua string) *Agent {
  agent := &Agent{}

  if !self.IsRobot(ua) {
    if agent.Browser = self.parseBrowserFromUserAgent(ua); agent.Browser != nil {
      agent.String = ua
      browserType, found := self.GetBrowserType(agent.Browser.TypeId)
      if !found {
        browserType = self.OtherBrowserType()
      }
      agent.Type = browserType.Name

      agent.Os, found = self.GetOsForBrowser(agent.Browser.Id)
      if !found {
        agent.Os = self.parseOsFromUserAgent(ua)
      }

      agent.Device = self.parseDeviceFromUserAgent(ua)
    }
  }

  return agent
}

package uas

import (
  _ "fmt"
)

func (self *Manifest) FindRobot(ua string) (*Robot, bool) {
  for _, robot := range self.Data.Robots {
    if ua == robot.UserAgent {
      return robot, true
    }
  }
  return nil, false
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

  _, found := self.FindRobot(ua)
  if !found {
    for _, reg := range self.Data.BrowsersReg {
      if reg.Reg.MatchString(ua) {
        agent.String = ua
        agent.Browser, found = self.GetBrowser(reg.BrowserId)
        if found {
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
        return agent
      }
    }
  }

  return agent
}

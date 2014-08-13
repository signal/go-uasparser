package uas

import (
  _ "fmt"
)

const (
  UnknownOsName   = "unknown"
  OtherDeviceName = "Other"
)

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

func (self *Manifest) GetBrowser(id int) (*Browser, bool) {
  for _, browser := range self.Data.Browsers {
    if id == browser.Id {
      return browser, true
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

func (self *Manifest) GetDevice(id int) (*Device, bool) {
  for _, device := range self.Data.Devices {
    if id == device.Id {
      return device, true
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

func (self *Manifest) parseOs(ua string) *Os {
  for _, reg := range self.Data.OperatingSystemsReg {
    if reg.Reg.MatchString(ua) {
      os, _ := self.GetOs(reg.OsId)
      return os
    }
  }
  return self.UnknownOs()
}

func (self *Manifest) parseDevice(ua string) *Device {
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
        agent.Type = "Mobile Browser" // FIXME
        agent.Browser, _ = self.GetBrowser(reg.BrowserId)
        agent.Os = self.parseOs(ua)
        agent.Device = self.parseDevice(ua)
        return agent
      }
    }

    // os

    // device
  }

  return agent
}

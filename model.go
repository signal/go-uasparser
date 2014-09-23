package uas

import (
	"encoding/xml"
	"regexp"
	"strings"
)

const (
	UnknownBrowserName     = "unknown"
	OtherBrowserTypeName   = "Other"
	UnknownBrowserTypeName = "unknown"
	UnknownOsName          = "unknown"
	OtherDeviceName        = "Other"
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

	// UnknownBrowserVersion is a BrowserVersion that represents an unknown match.
	UnknownBrowserVersion *BrowserVersion
	// UnknownBrowser is a Browser that represents an unknown browser.
	UnknownBrowser *Browser
	// OtherBrowserType is a Browser that represents the Other type.
	OtherBrowserType *BrowserType
	// UnknownOs is an Os that represents the Unknown type.
	UnknownOs *Os
	// OtherDevice is a Device that represents the Other type.
	OtherDevice *Device
}

type BrowserVersion struct {
	*Browser
	Version string
}
type Agent struct {
	String         string
	Type           string
	BrowserVersion *BrowserVersion
	Os             *Os
	Device         *Device
}

// Returns the Robot instance (along with a true value indicating success) matching the provided
// user-agent string. Otherwise a nil and false are returned.
func (self *Manifest) FindRobot(ua string) (*Robot, bool) {
	for _, robot := range self.Data.Robots {
		if ua == strings.Trim(robot.UserAgent, " ") {
			return robot, true
		}
	}
	return nil, false
}

// Returns true if the given user-agent string matches the user-agent for a Robot (bot).
func (self *Manifest) IsRobot(ua string) bool {
	_, found := self.FindRobot(ua)
	return found
}

// Finds a specific Browser by its listed ID in the manifest. Returns nil,false if
// no Browser has the given ID; otherwise it returns a Browser instance and true.
func (self *Manifest) GetBrowser(id int) (*Browser, bool) {
	for _, browser := range self.Data.Browsers {
		if id == browser.Id {
			return browser, true
		}
	}
	return nil, false
}

// Finds a specific BrowserType by its listed ID in the manifest. Returns nil,false if
// no BrowserType has the given ID; otherwise it returns a BrowserType instance and true.
func (self *Manifest) GetBrowserType(id int) (*BrowserType, bool) {
	for _, browserType := range self.Data.BrowserTypes {
		if id == browserType.Id {
			return browserType, true
		}
	}
	return nil, false
}

// Finds a specific Os by its listed ID in the manifest. Returns nil,false if
// no OS has the given ID; otherwise it returns an Os instance and true.
func (self *Manifest) GetOs(id int) (*Os, bool) {
	for _, os := range self.Data.OperatingSystems {
		if id == os.Id {
			return os, true
		}
	}
	return nil, false
}

// Finds the Os tied to a Browser given by the Browser's ID in the manifest. Returns nil,false
// if no Browser has the given ID or if no Os exists for the ID listed with the Browser (unlikely).
// Otherwise it returns an Os instance and true.
func (self *Manifest) GetOsForBrowser(id int) (*Os, bool) {
	for _, browserOs := range self.Data.BrowsersOs {
		if id == browserOs.BrowserId {
			return self.GetOs(browserOs.OsId)
		}
	}
	return nil, false
}

// Finds a specific Device by its ID in the manifest. Returns nil,false if
// no Device has the given ID; otherwise it returns a Device instance and true.
func (self *Manifest) GetDevice(id int) (*Device, bool) {
	for _, device := range self.Data.Devices {
		if id == device.Id {
			return device, true
		}
	}
	return nil, false
}

// Finds a specific BrowserType by its listed name in the manifest. Returns nil,false if
// no BrowserType has the given name; otherwise it returns a BrowserType instance and true.
func (self *Manifest) FindBrowserTypeByName(name string) (*BrowserType, bool) {
	for _, os := range self.Data.BrowserTypes {
		if name == os.Name {
			return os, true
		}
	}
	return nil, false
}

// Finds a specific Os by its listed name in the manifest. Returns nil,false if
// no OS has the given name; otherwise it returns an Os instance and true.
func (self *Manifest) FindOsByName(name string) (*Os, bool) {
	for _, os := range self.Data.OperatingSystems {
		if name == os.Name {
			return os, true
		}
	}
	return nil, false
}

// Finds a specific Device by its listed name in the manifest. Returns nil,false if
// no device has the given name; otherwise it returns a Device instance and true.
func (self *Manifest) FindDeviceByName(name string) (*Device, bool) {
	for _, device := range self.Data.Devices {
		if name == device.Name {
			return device, true
		}
	}
	return nil, false
}

// Parses out a BrowserVersion instance from a user-agent string. That is, it finds
// a matching Browser and extracts a version string if it can. Returns nil if no
// match could be found.
func (self *Manifest) ParseBrowserVersion(ua string) *BrowserVersion {
	for _, reg := range self.Data.BrowsersReg {
		if matches := reg.Reg.FindStringSubmatch(ua); matches != nil {
			browser, _ := self.GetBrowser(reg.BrowserId)
			return &BrowserVersion{browser, strings.Join(matches[1:], "/")}
		}
	}
	return nil
}

// Parses a Os from the provider user-agent string. You may get different
// results over what you might get from calling Parse as this is a less deductive
// function. Returns nil if no Device is matched.
func (self *Manifest) ParseOs(ua string) *Os {
	for _, reg := range self.Data.OperatingSystemsReg {
		if reg.Reg.MatchString(ua) {
			os, _ := self.GetOs(reg.OsId)
			return os
		}
	}
	return nil
}

// Parses a Device from the provider user-agent string. You may get different
// results over what you might get from calling Parse as this is a less deductive
// function. Returns nil if no Device is matched.
func (self *Manifest) ParseDevice(ua string) *Device {
	for _, reg := range self.Data.DevicesReg {
		if reg.Reg.MatchString(ua) {
			device, _ := self.GetDevice(reg.DeviceId)
			return device
		}
	}
	return nil
}

func (self *Manifest) deduceDevice(agentType string) *Device {
	switch agentType {
	case "Other", "Library", "Validator", "Useragent Anonymizer":
		return self.OtherDevice
	case "Mobile Browser", "Wap Browser":
		device, _ := self.FindDeviceByName("Smartphone")
		return device
	default:
		device, _ := self.FindDeviceByName("Personal computer")
		return device
	}
}

// Parses a provided user-agent string and returns an Agent instance. If the user-agent
// matches that of a robot, nil is returned. If no browser is matched, it will be listed
// as unknown, but the OS and device may still be matched.
func (self *Manifest) Parse(ua string) *Agent {
	var agent *Agent

	if len(ua) > 0 && !self.IsRobot(ua) {
		agent = &Agent{String: ua}

		if browserVer := self.ParseBrowserVersion(ua); browserVer != nil {
			agent.BrowserVersion = browserVer
			browserType, found := self.GetBrowserType(browserVer.TypeId)
			if !found {
				browserType = self.OtherBrowserType
			}
			agent.Type = browserType.Name
			agent.Os, _ = self.GetOsForBrowser(browserVer.Id)
		} else {
			agent.BrowserVersion = self.UnknownBrowserVersion
			agent.Type = UnknownBrowserTypeName
		}

		// if no OS found using the browser, find one directly
		if agent.Os == nil {
			if agent.Os = self.ParseOs(ua); agent.Os == nil {
				agent.Os = self.UnknownOs
			}
		}

		if agent.Device = self.ParseDevice(ua); agent.Device == nil {
			agent.Device = self.deduceDevice(agent.Type)
		}
	}

	return agent
}

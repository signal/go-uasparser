package uas

import (
  "fmt"
  "os"
  "path/filepath"
  "testing"
)

// setup

func loadManifest(fileName string) *Manifest {
  filePath, err := filepath.Abs(fmt.Sprintf("test/data/%s", fileName))
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  manifest, err := LoadFile(filePath)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return manifest
}

var manifest *Manifest

func init() {
  manifest = loadManifest("uas-20140812-01.xml")
}

// tests

func TestLoad_PartialFile(t *testing.T) {
  filePath, err := filepath.Abs("test/data/uas-partial-data.xml")
  if err != nil {
    t.Error(fmt.Sprintf("Unexpected error", err))
  }

  file, err := os.Open(filePath)
  if err != nil {
    t.Error("Unexpected error", err)
  }
  defer file.Close()

  _, err = Load(file)
  if err == nil {
    t.Error("Expected an error reading partial file")
  }

  AssertEquals(t, "error message reading file", err.Error(), "EOF")
}

func TestLoadValidFile_Robots(t *testing.T) {
  robots := manifest.Data.Robots
  AssertEquals(t, "length", 1387, len(robots))
  AssertEquals(t, "first id", 14157, robots[0].Id)
  AssertEquals(t, "last id", 3441, robots[1386].Id)

  robot := robots[0]
  AssertEquals(t, "name", " Scrubby/3.1", robot.Name)
  AssertEquals(t, "company", "Scrub The Web", robot.Company)
  AssertEquals(t, "url company", "http://www.scrubtheweb.com/", robot.URLCompany)
  AssertEquals(t, "icon", "bot_scrub.png", robot.Icon)
  AssertEquals(t, "family", " Scrubby", robot.Family)
  AssertEquals(t,
    "user agent",
    "Mozilla/5.0 (compatible; Scrubby/3.1; +http://www.scrubtheweb.com/help/technology.html)",
    robot.UserAgent)
  AssertEquals(t, "info url", "/list-of-ua/bot-detail?bot= Scrubby", robot.InfoURL)
}

func TestLoadValidFile_OperatingSystems(t *testing.T) {
  oses := manifest.Data.OperatingSystems
  AssertEquals(t, "length", 122, len(oses))
  AssertEquals(t, "first id", 1, oses[0].Id)
  AssertEquals(t, "last id", 146, oses[121].Id)

  os := oses[0]
  AssertEquals(t, "name", "Windows XP", os.Name)
  AssertEquals(t, "company", "Microsoft Corporation.", os.Company)
  AssertEquals(t, "url company", "http://www.microsoft.com/", os.URLCompany)
  AssertEquals(t, "icon", "windowsxp.png", os.Icon)
  AssertEquals(t, "family", "Windows", os.Family)
  AssertEquals(t, "info url", "/list-of-ua/os-detail?os=Windows XP", os.InfoURL)
}

func TestLoadValidFile_Browsers(t *testing.T) {
  browsers := manifest.Data.Browsers
  AssertEquals(t, "length", 463, len(browsers))
  AssertEquals(t, "first id", 1, browsers[0].Id)
  AssertEquals(t, "last id", 526, browsers[462].Id)

  browser := browsers[0]
  AssertEquals(t, "name", "Camino", browser.Name)
  AssertEquals(t, "company", "Mozilla Foundation", browser.Company)
  AssertEquals(t, "url company", "http://www.mozilla.org/", browser.URLCompany)
  AssertEquals(t, "icon", "camino.png", browser.Icon)
  AssertEquals(t, "url", "http://caminobrowser.org/", browser.URL)
  AssertEquals(t, "type id", 0, browser.TypeId)
  AssertEquals(t, "info url", "/list-of-ua/browser-detail?browser=Camino", browser.InfoURL)
}

func TestLoadValidFile_BrowserTypes(t *testing.T) {
  types := manifest.Data.BrowserTypes
  AssertEquals(t, "length", 11, len(types))

  AssertEquals(t, "first type id", 0, types[0].Id)
  AssertEquals(t, "first type name", "Browser", types[0].Name)

  AssertEquals(t, "last type id", 50, types[10].Id)
  AssertEquals(t, "last type name", "Useragent Anonymizer", types[10].Name)
}

func TestLoadValidFile_BrowserRegs(t *testing.T) {
  regs := manifest.Data.BrowsersReg
  AssertEquals(t, "length", 628, len(regs))
  AssertEquals(t, "first order", 1, regs[0].Order)
  AssertEquals(t, "last order", 628, regs[627].Order)

  reg := regs[627]
  AssertEquals(t, "browser id", 282, reg.BrowserId)
  AssertEquals(t, "regstring", "/WinHttp/si", reg.RegString)

  // check actual regs
  AssertEquals(t, "complex regstring",
    "(?si:^Mozilla.*Android.*AppleWebKit.*Chrome.*OPR\\/([0-9\\.]+))",
    regs[0].Reg.String())
  AssertEquals(t, "simple regstring", "(?si:WinHttp)", regs[627].Reg.String())
}

func TestLoadValidFile_BrowserOperatingSystems(t *testing.T) {
  oses := manifest.Data.BrowsersOs
  AssertEquals(t, "length", 72, len(oses))
  AssertEquals(t, "first browser id", 18, oses[0].BrowserId)
  AssertEquals(t, "first os id", 44, oses[0].OsId)

  AssertEquals(t, "last browser id", 515, oses[71].BrowserId)
  AssertEquals(t, "last os id", 87, oses[71].OsId)
}

func TestLoadValidFile_OperatingSystemRegs(t *testing.T) {
  regs := manifest.Data.OperatingSystemsReg
  AssertEquals(t, "length", 219, len(regs))
  AssertEquals(t, "first order", 1, regs[0].Order)
  AssertEquals(t, "last order", 219, regs[218].Order)

  reg := regs[0]
  AssertEquals(t, "os id", 35, reg.OsId)
  AssertEquals(t, "regstring", "/palm/si", reg.RegString)

  // check actual regs
  AssertEquals(t, "complex regstring",
    "(?si:^Mozilla\\/.*Ubuntu.*[Tablet|Mobile].*WebKit)",
    regs[22].Reg.String())
  AssertEquals(t, "simple regstring", "(?si:palm)", regs[0].Reg.String())
}

func TestLoadValidFile_Devices(t *testing.T) {
  devices := manifest.Data.Devices
  AssertEquals(t, "length", 8, len(devices))
  AssertEquals(t, "first id", 1, devices[0].Id)
  AssertEquals(t, "last id", 8, devices[7].Id)

  device := devices[0]
  AssertEquals(t, "name", "Other", device.Name)
  AssertEquals(t, "icon", "other.png", device.Icon)
  AssertEquals(t, "info url", "/list-of-ua/device-detail?device=Other", device.InfoURL)
}

func TestLoadValidFile_DeviceRegs(t *testing.T) {
  regs := manifest.Data.DevicesReg
  AssertEquals(t, "length", 108, len(regs))
  AssertEquals(t, "first order", 1, regs[0].Order)
  AssertEquals(t, "last order", 108, regs[107].Order)

  reg := regs[107]
  AssertEquals(t, "device id", 4, reg.DeviceId)
  AssertEquals(t, "regstring", "/^Mozilla.*Android.*Tablet.*AppleWebKit/si", reg.RegString)

  // check actual regs
  AssertEquals(t, "simple regstring",
    "(?si:^Mozilla.*Android.*Tablet.*AppleWebKit)",
    regs[107].Reg.String())
  AssertEquals(t, "complex regstring",
    "(?si:^Mozilla.*Android.*GT\\-("+
      "P1000|P1010|P3100|P3105|P3110|P3113|P5100|P5110|P5113|P5200|P5210|P6200|P6201|P6210|P6211"+
      "|P6800|P6810|P7110|P7300|P7310|P7320|P7500|P7510|P7511))",
    regs[89].Reg.String())
}

// description

func TestLoadValidFile_Description(t *testing.T) {
  description := manifest.Description
  AssertEquals(t, "label",
    "Data (format xml) for UASparser - http://user-agent-string.info/download/UASparser",
    description.Label)
  AssertEquals(t, "version", "20140812-01", description.Version)
  AssertEquals(t, "checksum length", 2, len(description.Checksums))

  AssertEquals(t, "first checksum type", "MD5", description.Checksums[0].Type)
  AssertEquals(t, "first checksum url",
    "http://user-agent-string.info/rpc/get_data.php?format=xml&amp;md5=y",
    description.Checksums[0].URL)

  AssertEquals(t, "second checksum type", "SHA1", description.Checksums[1].Type)
  AssertEquals(t, "second checksum url",
    "http://user-agent-string.info/rpc/get_data.php?format=xml&amp;sha1=y",
    description.Checksums[1].URL)
}

// robots

func Test_NoRobotFound(t *testing.T) {
  ua := "GonzoBotz/1.0"
  _, ok := manifest.FindRobot(ua)
  Asserts(t, "no robot found", !ok)
}

func TestFindFirstRobot(t *testing.T) {
  ua := "Mozilla/5.0 (compatible; Scrubby/3.1; +http://www.scrubtheweb.com/help/technology.html)"
  robot, ok := manifest.FindRobot(ua)
  Asserts(t, "found robot", ok)
  AssertDeepEquals(t, "robot", manifest.Data.Robots[0], robot)
}

func TestFindLastRobot(t *testing.T) {
  ua := "^Nail (http://CaretNail.com)"
  robot, ok := manifest.FindRobot(ua)
  Asserts(t, "found robot", ok)
  AssertDeepEquals(t, "robot", manifest.Data.Robots[len(manifest.Data.Robots)-1], robot)
}

// browsers

func TestParse_NoUAProvided(t *testing.T) {
  AssertDeepEquals(t, "browser agent", &Agent{}, manifest.ParseBrowser(""))
}

func TestParse_WhenRobotUAProvided(t *testing.T) {
  ua := "Mozilla/5.0 (compatible; Scrubby/3.1; +http://www.scrubtheweb.com/help/technology.html)"
  AssertDeepEquals(t, "browser agent", &Agent{}, manifest.ParseBrowser(ua))
}

func TestParse_FindOperaMobileBrowser(t *testing.T) {
  ua := "Mozilla/5.0 (Linux; Android 2.3.4; MT11i Build/4.0.2.A.0.62) AppleWebKit/537.22 " +
    "(KHTML, like Gecko) Chrome/25.0.1364.123 Mobile Safari/537.22 OPR/14.0.1025.52315"

  browser, ok := manifest.GetBrowser(321) // Opera Mobile
  Asserts(t, "browser found", ok)
  os, ok := manifest.GetOs(107) // Android, Gingerbread
  Asserts(t, "os found", ok)
  device, ok := manifest.GetDevice(1) // Other
  Asserts(t, "device found", ok)

  agent := manifest.ParseBrowser(ua)
  AssertEquals(t, "agent string", ua, agent.String)
  AssertEquals(t, "agent type", "Mobile Browser", agent.Type)
  AssertDeepEquals(t, "agent browser", browser, agent.Browser)
  AssertDeepEquals(t, "agent os", os, agent.Os)
  AssertDeepEquals(t, "agent device", device, agent.Device)
}

func TestParse_FindAgentWithBrowserOsMapping(t *testing.T) {
  ua := "Mozilla/5.0 (Macintosh; U; Intel Mac OS X; en-US) AppleWebKit/528.16 " +
    "(KHTML, like Gecko, Safari/528.16) OmniWeb/v622.8.0.112941"

  browser, ok := manifest.GetBrowser(18) // OmniWeb
  Asserts(t, "browser found", ok)
  os, ok := manifest.GetOs(44) // Mac OS
  Asserts(t, "os found", ok)
  device, ok := manifest.GetDevice(1) // Other
  Asserts(t, "device found", ok)

  agent := manifest.ParseBrowser(ua)
  AssertEquals(t, "agent string", ua, agent.String)
  AssertEquals(t, "agent type", "Browser", agent.Type)
  AssertDeepEquals(t, "agent browser", browser, agent.Browser)
  AssertDeepEquals(t, "agent os", os, agent.Os)
  AssertDeepEquals(t, "agent device", device, agent.Device)
}

func TestParse_BrowserFoundButUnknownOs(t *testing.T) {
  ua := "Mozilla/4.0 (compatible; DPlus 0.5)"

  browser, ok := manifest.GetBrowser(441) // DPlus
  Asserts(t, "browser found", ok)
  os, ok := manifest.FindOsByName(UnknownOsName)
  Asserts(t, "os found", ok)
  device, ok := manifest.GetDevice(1) // Other
  Asserts(t, "device found", ok)

  agent := manifest.ParseBrowser(ua)
  AssertEquals(t, "agent string", ua, agent.String)
  AssertEquals(t, "agent type", "Browser", agent.Type)
  AssertDeepEquals(t, "agent browser", browser, agent.Browser)
  AssertDeepEquals(t, "agent os", os, agent.Os)
  AssertDeepEquals(t, "agent device", device, agent.Device)
}

package uas

import (
  "fmt"
  "os"
  "path/filepath"
  "testing"
)

func loadManifest(fileName string) *Manifest {
  filePath, err := filepath.Abs(fmt.Sprintf("test/data/%s", fileName))
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  file, err := os.Open(filePath)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  defer file.Close()

  manifest, err := Load(file)
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

// helpers

func assertEquals(t *testing.T, what string, expected interface{}, actual interface{}) {
  if expected != actual {
    t.Error("Expected", what, "to be", expected, "but instead was", actual, "instead")
  }
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

  assertEquals(t, "error message reading file", err.Error(), "EOF")
}

func TestLoadValidFile_Robots(t *testing.T) {
  robots := manifest.Data.Robots
  assertEquals(t, "length", 1387, len(robots))
  assertEquals(t, "first id", 14157, robots[0].Id)
  assertEquals(t, "last id", 3441, robots[1386].Id)

  robot := robots[0]
  assertEquals(t, "name", " Scrubby/3.1", robot.Name)
  assertEquals(t, "company", "Scrub The Web", robot.Company)
  assertEquals(t, "url company", "http://www.scrubtheweb.com/", robot.URLCompany)
  assertEquals(t, "icon", "bot_scrub.png", robot.Icon)
  assertEquals(t, "family", " Scrubby", robot.Family)
  assertEquals(t,
    "user agent",
    "Mozilla/5.0 (compatible; Scrubby/3.1; +http://www.scrubtheweb.com/help/technology.html)",
    robot.UserAgent)
  assertEquals(t, "info url", "/list-of-ua/bot-detail?bot= Scrubby", robot.InfoURL)
}

func TestLoadValidFile_OperatingSystems(t *testing.T) {
  oses := manifest.Data.OperatingSystems
  assertEquals(t, "length", 122, len(oses))
  assertEquals(t, "first id", 1, oses[0].Id)
  assertEquals(t, "last id", 146, oses[121].Id)

  os := oses[0]
  assertEquals(t, "name", "Windows XP", os.Name)
  assertEquals(t, "company", "Microsoft Corporation.", os.Company)
  assertEquals(t, "url company", "http://www.microsoft.com/", os.URLCompany)
  assertEquals(t, "icon", "windowsxp.png", os.Icon)
  assertEquals(t, "family", "Windows", os.Family)
  assertEquals(t, "info url", "/list-of-ua/os-detail?os=Windows XP", os.InfoURL)
}

func TestLoadValidFile_Browsers(t *testing.T) {
  browsers := manifest.Data.Browsers
  assertEquals(t, "length", 463, len(browsers))
  assertEquals(t, "first id", 1, browsers[0].Id)
  assertEquals(t, "last id", 526, browsers[462].Id)

  browser := browsers[0]
  assertEquals(t, "name", "Camino", browser.Name)
  assertEquals(t, "company", "Mozilla Foundation", browser.Company)
  assertEquals(t, "url company", "http://www.mozilla.org/", browser.URLCompany)
  assertEquals(t, "icon", "camino.png", browser.Icon)
  assertEquals(t, "url", "http://caminobrowser.org/", browser.URL)
  assertEquals(t, "type", 0, browser.Type)
  assertEquals(t, "info url", "/list-of-ua/browser-detail?browser=Camino", browser.InfoURL)
}

func TestLoadValidFile_BrowserRegs(t *testing.T) {
  regs := manifest.Data.BrowsersReg
  assertEquals(t, "length", 628, len(regs))
  assertEquals(t, "first order", 1, regs[0].Order)
  assertEquals(t, "last order", 628, regs[627].Order)

  reg := regs[627]
  assertEquals(t, "browser id", 282, reg.BrowserId)
  assertEquals(t, "regstring", "/WinHttp/si", reg.RegString)

  // check actual regs
  assertEquals(t, "complex regstring",
    "(?si:^Mozilla.*Android.*AppleWebKit.*Chrome.*OPR\\/([0-9\\.]+))",
    regs[0].Reg.String())
  assertEquals(t, "simple regstring", "(?si:WinHttp)", regs[627].Reg.String())
}

func TestLoadValidFile_BrowserOperatingSystems(t *testing.T) {
  oses := manifest.Data.BrowsersOs
  assertEquals(t, "length", 72, len(oses))
  assertEquals(t, "first browser id", 18, oses[0].BrowserId)
  assertEquals(t, "first os id", 44, oses[0].OsId)

  assertEquals(t, "last browser id", 515, oses[71].BrowserId)
  assertEquals(t, "last os id", 87, oses[71].OsId)
}

func TestLoadValidFile_OperatingSystemRegs(t *testing.T) {
  regs := manifest.Data.OperatingSystemsReg
  assertEquals(t, "length", 219, len(regs))
  assertEquals(t, "first order", 1, regs[0].Order)
  assertEquals(t, "last order", 219, regs[218].Order)

  reg := regs[0]
  assertEquals(t, "os id", 35, reg.OsId)
  assertEquals(t, "regstring", "/palm/si", reg.RegString)

  // check actual regs
  assertEquals(t, "complex regstring",
    "(?si:^Mozilla\\/.*Ubuntu.*[Tablet|Mobile].*WebKit)",
    regs[22].Reg.String())
  assertEquals(t, "simple regstring", "(?si:palm)", regs[0].Reg.String())
}

package uas

import (
  "fmt"
  "os"
  "path/filepath"
  "testing"
)

// helpers

func assertEquals(t *testing.T, what string, expected interface{}, actual interface{}) {
  if expected != actual {
    t.Error("Expected", what, "to be", expected, "but instead was", actual, "instead")
  }
}

func openDataFile(t *testing.T, name string) *os.File {
  filePath, err := filepath.Abs(fmt.Sprintf("test/data/%s", name))
  if err != nil {
    t.Error(fmt.Sprintf("Unexpected error", err))
  }

  file, err := os.Open(filePath)
  if err != nil {
    t.Error("Unexpected error", err)
  }
  return file
}

func loadManifest(t *testing.T, fileName string) *Manifest {
  file := openDataFile(t, fileName)
  defer file.Close()

  manifest, err := Load(file)
  if err != nil {
    t.Error("Unexpected error", err)
  }
  return manifest
}

// tests

func TestLoad_PartialFile(t *testing.T) {
  file := openDataFile(t, "uas-partial-data.xml")
  defer file.Close()

  _, err := Load(file)
  if err == nil {
    t.Error("Expected an error reading partial file")
  }

  assertEquals(t, "error message reading file", err.Error(), "EOF")
}

func TestLoadValidFile_Robots(t *testing.T) {
  manifest := loadManifest(t, "uas-20140812-01.xml")
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
  manifest := loadManifest(t, "uas-20140812-01.xml")
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
  manifest := loadManifest(t, "uas-20140812-01.xml")
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

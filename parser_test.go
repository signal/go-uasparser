package uas

import (
  "testing"
)

// robots

func Test_NoRobotFound(t *testing.T) {
  manifest := LoadManifest("uas-20140812-01.xml")
  ua := "GonzoBotz/1.0"
  _, ok := manifest.FindRobot(ua)
  Asserts(t, "no robot found", !ok)
}

func TestFindFirstRobot(t *testing.T) {
  manifest := LoadManifest("uas-20140812-01.xml")
  ua := "Mozilla/5.0 (compatible; Scrubby/3.1; +http://www.scrubtheweb.com/help/technology.html)"
  robot, ok := manifest.FindRobot(ua)
  Asserts(t, "found robot", ok)
  AssertDeepEquals(t, "robot", manifest.Data.Robots[0], robot)
}

func TestFindLastRobot(t *testing.T) {
  manifest := LoadManifest("uas-20140812-01.xml")
  ua := "^Nail (http://CaretNail.com)"
  robot, ok := manifest.FindRobot(ua)
  Asserts(t, "found robot", ok)
  AssertDeepEquals(t, "robot", manifest.Data.Robots[len(manifest.Data.Robots)-1], robot)
}

// browsers

func TestParse_NoUAProvided(t *testing.T) {
  manifest := LoadManifest("uas-20140812-01.xml")
  AssertDeepEquals(t, "browser agent", &Agent{}, manifest.ParseBrowser(""))
}

func TestParse_WhenRobotUAProvided(t *testing.T) {
  manifest := LoadManifest("uas-20140812-01.xml")
  ua := "Mozilla/5.0 (compatible; Scrubby/3.1; +http://www.scrubtheweb.com/help/technology.html)"
  AssertDeepEquals(t, "browser agent", &Agent{}, manifest.ParseBrowser(ua))
}

func TestParse_FindOperaMobileBrowser(t *testing.T) {
  manifest := LoadManifest("uas-20140812-01.xml")
  ua := "Mozilla/5.0 (Linux; Android 2.3.4; MT11i Build/4.0.2.A.0.62) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.123 Mobile Safari/537.22 OPR/14.0.1025.52315"

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

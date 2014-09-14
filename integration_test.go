package uas

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"testing"
)

func loadCsvFile(path string) *os.File {
	testFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return testFile
}

func TestIntegration_Browsers(t *testing.T) {
	manifest, err := LoadFile("tmp/uas-manifest.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	testFile := loadCsvFile("tmp/uas-browser-tests.csv")
	defer testFile.Close()
	reader := csv.NewReader(testFile)
	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		typeName, browserName, uastr := line[0], line[1], strings.Trim(line[2], " ")
		if typeName == "robot" {
			Asserts(t, fmt.Sprintf("[%s] to be a robot", uastr), manifest.IsRobot(uastr))
		} else {
			agent := manifest.Parse(uastr)

			if agent != nil {
				AssertEquals(t, fmt.Sprintf("type for [%s]", uastr), typeName, agent.Type)
				AssertEquals(t, fmt.Sprintf("name for [%s]", uastr), browserName, agent.BrowserVersion.Name)
			} else {
				t.Error("Expected to find agent with type:", typeName, "and browser:",
					browserName, "for", uastr, "but found nothing")
			}
		}
	}
}

func TestIntegration_OperatingSystems(t *testing.T) {
	manifest, err := LoadFile("tmp/uas-manifest.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	testFile := loadCsvFile("tmp/uas-os-tests.csv")
	defer testFile.Close()
	reader := csv.NewReader(testFile)
	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		osName, uastr := line[0], strings.Trim(line[1], " ")
		os := manifest.ParseOs(uastr)

		if os != nil {
			AssertEquals(t, fmt.Sprintf("name for [%s]", uastr), osName, os.Name)
		} else {
			// if no device directly found, perhaps we can deduce it from the agent
			agent := manifest.Parse(uastr)
			if agent != nil {
				AssertEquals(t, fmt.Sprintf("name for [%s]", uastr), osName, agent.Os.Name)
			} else if !manifest.IsRobot(uastr) {
				// whelp, it wasn't a robot either
				t.Error("Expected to find agent with os:", osName, "for", uastr)
			}
		}
	}
}

func TestIntegration_Devices(t *testing.T) {
	manifest, err := LoadFile("tmp/uas-manifest.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	testFile := loadCsvFile("tmp/uas-device-tests.csv")
	defer testFile.Close()
	reader := csv.NewReader(testFile)
	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		deviceName, uastr := line[0], strings.Trim(line[1], " ")
		device := manifest.ParseDevice(uastr)

		if device != nil {
			AssertEquals(t, fmt.Sprintf("name for [%s]", uastr), deviceName, device.Name)
		} else {
			agent := manifest.Parse(uastr)
			// if no device directly found, perhaps we can deduce it from the agent
			if agent != nil {
				AssertEquals(t, fmt.Sprintf("name for [%s]", uastr), deviceName, agent.Device.Name)
			} else if !manifest.IsRobot(uastr) {
				// whelp, it wasn't a robot either
				t.Error("Expected to find device:", deviceName, "for", uastr, "but found nothing")
			}
		}
	}
}

// From issue#2

type agentTest struct {
	ua          string
	agentType   string
	browserName string
	osName      string
	deviceName  string
}

var specificAgentsExpected = [...]agentTest{
	{"MT6572/V1 Linux/3.4.5 Android/4.3 Release/12.28.2013 Browser/AppleWebKit534.30 Profile/MIDP-2.0 Configuration/CLDC-1.1 Mobile Safari/534.30 Android 4.3;",
		"unknown", "unknown", "Android 4.3 Jelly Bean", "Personal computer"},
	{"Bmobile_AX530 Linux/2.6.35.7 Android/2.3.6 Release/10.10.2012 Browser/AppleWebKit533.1 Profile/MIDP-2.0 Configuration/CLDC-1.1 Mobile Safari/533.1",
		"unknown", "unknown", "Android", "Personal computer"},
	{"TBW975112_9300_V0701 Linux/3.4.5 Android/4.2.2 Release/03.26.2013 Browser/AppleWebKit534.30 Mobile Safari/534.30 System/Android 4.2.2",
		"unknown", "unknown", "Android 4.2 Jelly Bean", "Personal computer"},
	{"MT6572_TD/V1 Linux/3.4.5 Android/4.2.2 Release/03.26.2013 Browser/AppleWebKit534.30 Mobile Safari/534.30 MBBMS/2.2;",
		"unknown", "unknown", "Android", "Personal computer"},
	{"Lenovo-A889/S100 Linux/3.4.5 Android/4.2 Release/08.07.2013 Browser/AppleWebKit 534.30 Profile/ Configuration;",
		"unknown", "unknown", "Android", "Personal computer"},
	{"sprd-F950/1.0 Linux/2.6.35.7 Android/4.0.4 Release/06.03.2013 Browser/AppleWebKit533.1 (KHTML, like Gecko) Mozilla/5.0 Mobile",
		"unknown", "unknown", "Android", "Personal computer"},
	{"Karbonn A50s/V1 Linux/3.4.5 Android/4.2.2 Release/03.26.2013 Browser/AppleWebKit534.30 Mobile Safari/534.30 MBBMS/2.2;",
		"unknown", "unknown", "Android", "Personal computer"},
	{"ultrafone 105(Linux; U;) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0.3 Mobile Safari/534.30",
		"unknown", "unknown", "Linux", "Personal computer"},
}

func TestSpecificAgents(t *testing.T) {
	manifest, err := LoadFile("tmp/uas-manifest.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, test := range specificAgentsExpected {
		//test := specificAgents
		ua := test.ua
		agent := manifest.Parse(test.ua)
		AssertEquals(t, fmt.Sprintf("agent type for [%s]", ua), test.agentType, agent.Type)
		AssertEquals(t, fmt.Sprintf("browser name for [%s]", ua), test.browserName, agent.BrowserVersion.Name)
		AssertEquals(t, fmt.Sprintf("os name for [%s]", ua), test.osName, agent.Os.Name)
		AssertEquals(t, fmt.Sprintf("device name for [%s]", ua), test.deviceName, agent.Device.Name)
	}
}

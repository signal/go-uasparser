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

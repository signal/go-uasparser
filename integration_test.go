package uas

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestIntegration_Browsers(t *testing.T) {
	manifest, err := LoadFile("tmp/uas-manifest.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	testFile, err := os.Open("tmp/uas-browser-tests.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer testFile.Close()

	reader := csv.NewReader(testFile)
	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		typeName, browserName, uastr := line[0], line[1], strings.Trim(line[2], " ")
		if typeName == "robot" {
			Asserts(t, "is a robot", manifest.IsRobot(uastr))
		} else {
			agent := manifest.ParseBrowser(uastr)

			if agent != nil {
				AssertEquals(t, fmt.Sprintf("type for [%s]", uastr), typeName, agent.Type)
				AssertEquals(t, fmt.Sprintf("name for [%s]", uastr), browserName, agent.Browser.Name)
			} else {
				t.Error("Expected to find agent with type:", typeName, "and browser:",
					browserName, "for", uastr)
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

	testFile, err := os.Open("tmp/uas-os-tests.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer testFile.Close()

	reader := csv.NewReader(testFile)
	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		osName, uastr := line[0], strings.Trim(line[1], " ")
		agent := manifest.ParseBrowser(uastr)

		if agent != nil {
			AssertEquals(t, fmt.Sprintf("name for [%s]", uastr), osName, agent.Os.Name)
		} else {
			t.Error("Expected to find agent with os:", osName, "for", uastr)
		}
	}
}

func TestIntegration_Devices(t *testing.T) {
	manifest, err := LoadFile("tmp/uas-manifest.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	testFile, err := os.Open("tmp/uas-device-tests.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer testFile.Close()

	reader := csv.NewReader(testFile)
	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		deviceName, uastr := line[0], strings.Trim(line[1], " ")
		agent := manifest.ParseBrowser(uastr)

		if agent != nil {
			AssertEquals(t, fmt.Sprintf("name for [%s]", uastr), deviceName, agent.Device.Name)
		} else {
			t.Error("Expected to find agent with device:", deviceName, "for", uastr)
		}
	}
}

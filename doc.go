/*
Package uas provides a go implementation of the http://user-agent-string.info/ processor.
Standard usage is to provide a user-agent string to the Parse method of a Manifest
instance and retrieve an Agent instance in return. From the Agent, you can obtain:
browser details, operating system details, and device details.

You must create a Manifest instance by providing an XML file from the UAS.info site
(http://user-agent-string.info/rpc/get_data.php?key=free&format=xml&download=y) to the
LoadFile function; or you can provide a Reader of similar ilk to the Load function.
This package currently doesn't support downloading Manifests automatically, but you can
also easily create new instances of different Manifests; i.e. a Manifest is not a global
object. This package also does not yet support the .ini format; mostly this was to
make processing easier by using the built-in XML unmarshalling capabilities of Go.

	import (
		"fmt"
		"os"
		"github.com/signal/go-uasparser"
	)
	var manifest, err := uas.LoadFile("/tmp/uas_YYYYMMDD-01.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

Given a Manifest, you can now easily parse an Agent like so:

	var agent := manifest.Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) ...")
	if agent != nil {
		fmt.Println("Agent type:", agent.Type)
		fmt.Println("Browser name:", agent.BrowserVersion.Name)
		fmt.Println("Browser Version:", agent.BrowserVersion.Version)
		fmt.Println("OS name:", agent.Os.Name)
		fmt.Println("Device name:", agent.Device.Name)
	}

You can check out the model structure to figure out what other values are available.
Unlike other implementations, the values are not simply returned as a flat map.

Currently, robots are treated differently in that any agent recognized as one is
returned from Parse as a nil value. You can check to see if the agent is indeed a
robot by asking if it's so:

	if manifest.IsRobot("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)") {
	  fmt.Println("I AM A ROBOT")
	}

In all cases, when an Agent is found it will be cached in a Manifest-specific LRU that
can hold 5000 entries. This is not configurable at the moment.
*/
package uas

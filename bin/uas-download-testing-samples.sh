#!/bin/bash

wget -q \
  -O tmp/uas-data/uas-manifest.xml \
  "http://user-agent-string.info/rpc/get_data.php?key=free&format=xml&download=y"

wget -q \
  -O tmp/uas-data/uas-browser-tests.csv \
  "http://user-agent-string.info/rpc/get_data.php?uaslist=csv"
wget -q \
  -O tmp/uas-data/uas-os-tests.csv \
  "http://user-agent-string.info/rpc/get_data.php?uasOSlist=csv"
wget -q \
  -O tmp/uas-data/uas-device-tests.csv \
  "http://user-agent-string.info/rpc/get_data.php?uasDEVICElist=csv"


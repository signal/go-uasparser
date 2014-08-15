#!/bin/bash

wget -nv \
  -O tmp/uas-manifest.xml \
  "http://user-agent-string.info/rpc/get_data.php?key=free&format=xml&download=y"

wget -nv \
  -O tmp/uas-browser-tests.csv \
  "http://user-agent-string.info/rpc/get_data.php?uaslist=csv"
wget -nv \
  -O tmp/uas-os-tests.csv \
  "http://user-agent-string.info/rpc/get_data.php?uasOSlist=csv"
wget -nv \
  -O tmp/uas-device-tests.csv \
  "http://user-agent-string.info/rpc/get_data.php?uasDEVICElist=csv"


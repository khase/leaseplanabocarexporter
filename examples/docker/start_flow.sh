#!/bin/bash

echo "get new cars.json"
../LeaseplanAbocarExporter get -a -o cars.json --poll 360 --exec "../useTemplate.sh"
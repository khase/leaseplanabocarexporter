#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

## get new cars.json
./LeaseplanAbocarExporter get -a -o cars.json

## use FillTemplate to compute .csv
./FillTemplate -t $SCRIPT_DIR/cars.csv.tpl -o ./cars.csv fromfile -f json ./cars.json

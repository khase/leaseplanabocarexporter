#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

## save last result
cp cars.json cars.last.json

## get new cars.json
./LeaseplanAbocarExporter get -a -o cars.json

## use FillTemplate to compute difference
./FillTemplate -t $SCRIPT_DIR/changed.txt.tpl -o ./changed.txt fromfile -f json ./cars.json ./cars.last.json
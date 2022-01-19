#!/bin/bash

echo "save last result"
cp cars.current.json cars.last.json || true ## will fail on first run because there is no current json yet
cp cars.json cars.current.json

if test -f "./cars.last.json"; then
    echo "use FillTemplate to compute difference"
    ../FillTemplate -t ../changed.txt.tpl -o ./changed.txt --exec "../sendNotification.sh" fromfile -f json ./cars.current.json ./cars.last.json
else
    echo "first run detected."
    echo "skipping changed handler."
fi

echo "use FillTemplate to compute csv"
../FillTemplate -t ../cars.csv.tpl -o ./cars.csv fromfile -f json ./cars.current.json

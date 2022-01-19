## save last result
Copy-Item cars.json cars.last.json

## get new cars.json
.\LeaseplanAbocarExporter.exe get -a -o cars.json

## use FillTemplate to compute difference
.\FillTemplate.exe -t "$PSScriptRoot\changed.txt.tpl" -o .\changed.txt fromfile -f json ./cars.json ./cars.last.json
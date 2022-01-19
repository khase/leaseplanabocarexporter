## get new cars.json
.\LeaseplanAbocarExporter.exe get -a -o cars.json

## use FillTemplate to compute .csv
.\FillTemplate.exe -t "$PSScriptRoot\cars.csv.tpl" -o .\cars.csv fromfile -f json .\cars.json
# Requirements

all these examples require the following tools to be available in the executing directory:
- LeaseplanAbocarExporter (https://github.com/khase/LeaseplanAbocarExporter)
- FillTemplate (https://github.com/khase/FillTemplate)

Additionally the `LeaseplanAbocarExporter login` command has to be executed beforehand. 

# Info

## generateExcel
generateExcel will generate a simple excel file (actually csv) with the following informations

| Name | Marke | Model | PS | BLP | Fahrzeugtyp | Farbe | Farbfamilie | Schaltung | Treibstoff | Antrieb | TürAnz | SitzAnz |
|------|-------|-------|----|-----|-------------|-------|-------------|-----------|------------|---------|--------|---------|

## changeSinceLastRun
changeSinceLastRun will generate a simple overview with newly added and just removed cars from the list.

**NOTE**: this script requires a cars.json from a previous export to be present in the current directory.

The output is in a simple Text Format
```
added: 
- Car 1 (url to the car)
- Car 2 (url to the car)
- Car 3 (url to the car)
removed: 
- Car 1 (url to the car - possible 404)
- Car 2 (url to the car - possible 404)
- Car 3 (url to the car - possible 404)
```

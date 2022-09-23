# Overview

`LeaseplanAbocarExporter` is a small go tool which exports a json containing all car offers available to you.

# Usage:
  LeaseplanAbocarExporter [flags]

  LeaseplanAbocarExporter [command]

## Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         get list of cars
  help        Help about any command
  login       login to the leaseplan abocar platform and safe the token to the current config

## Flags:

| short command | long command | description |
|---------------|--------------|-------------|
|      | --config string  | config file (default is $HOME/.leaseplanabocar.yaml) |
|  -h, | --help           | help for LeaseplanAbocarExporter |
|  -t, | --token string   | token to be used for auth (can be retrieved using LeaseplanAbocarExporter login) |
|      | --viper          | use Viper for configuration (default true) |

# Usage (Login):
  LeaseplanAbocarExporter login [flags]

## Flags:
| short command | long command | description |
|---------------|--------------|-------------|
|      | --email string     | email to be used for login |
|  -h, | --help             | help for login |
|      | --password string  | password to be used for login |

## Global Flags:
| short command | long command | description |
|---------------|--------------|-------------|
|      | --config string  | config file (default is $HOME/.leaseplanabocar.yaml) |
|  -t, | --token string   | token to be used for auth (can be retrieved using LeaseplanAbocarExporter login) |
|      | --viper          | use Viper for configuration (default true)|

# Usage (get):
  LeaseplanAbocarExporter get [flags]

## Flags:
| short command | long command | description |
|---------------|--------------|-------------|
|  -a, | --all            | indicates if the programm should page through all cars. |
|  -n, | --count int      | number of cars to get for a single page. (default 10) |
|      | --exec string    | command to execute when output file changed (requires output filename to be set) |
|  -h, | --help           | help for get 
|  -o, | --output string  | filename where the json should be written to. |
|  -p, | --page int       | page to get. (default 1) 
|      | --poll int       | polling cycle in minutes. |

## Global Flags:
| short command | long command | description |
|---------------|--------------|-------------|
|      | --config string  | config file (default is $HOME/.leaseplanabocar.yaml) |
|  -t, | --token string   | token to be used for auth (can be retrieved using LeaseplanAbocarExporter login) |
|      | --viper          | use Viper for configuration (default true) |

# Example

## Login
When first using the exporter you need to get a new Auth-Token.
With the command `LeaseplanAbocarExporter login` you can interactively login with your personal data and the exporter will print your personal Token. Additionally the the token will be saved to a personalized config file in your home directory `.leaseplanabbocarexporter`
Future executions of the exporter will try to retrieve the token from this config.

## Get visible offers
With the command `LeaseplanAbocarExporter get` the tool will query the api for the first page and print the available car offers.
```
got 10 cars
1) SEAT ATECA 2.0 TDI 4Drive DSG - 150HP (47250€)
2) SEAT ATECA 2.0 TDI 4Drive DSG - 150HP (47250€)
3) VOLVO Xc40 T3 Geartronic Momentum Pro - 163HP (43350€)
4) SEAT ATECA 1,5 TSI Style DSG - 150HP (37650€)
5) SEAT ATECA 1,5 TSI Style DSG - 150HP (37650€)
6) SEAT ATECA 1,5 TSI Style DSG - 150HP (37650€)
7) SEAT ATECA 1,5 TSI Style DSG - 150HP (37650€)
8) SEAT ATECA 1,5 TSI Style DSG - 150HP (37650€)
9) SEAT ATECA 1,5 TSI Style DSG - 150HP (37650€)
10) SEAT ATECA 1,5 TSI Style DSG - 150HP (37650€)
```

With the flags `count` and `page` you can manipulate which page to retrieve or how many cars should be retrieved for a page.

The flag `all` will tell the exporter to page through all pages and return a single result with all the offers available to you.

Using the `output` flag will write the complete result json object to the disc instead of printing it to the console. This way the data can be further processed with external tools.

## Car Object (anonymized)
```json
{
    "ReadonlyProperties": {
        "State": "Online"
    },
    "RentalObject": {
        "ReadonlyProperties": {
            "ContractState": null,
            "HasFollowupBooking": false,
            "State": null,
            "OfferTypeTarifgroupText": null,
            "VehicleProviderName": null,
            "OfferTypeTarifgroupName": null,
            "OfferTypeExternalTarifgroup": null
        },
        "ChassisNo": "",
        "DistMark": null,
        "ObjGroup": "",
        "Tarifgroup": "",
        "CarLabel": "",
        "CarModell": "",
        "CarModellspec": "",
        "DateLicense": null,
        "DateRegistration": "",
        "DateUnregister": null,
        "DateBuy": null,
        "KindOfGear": "",
        "KindOfFuel": "",
        "KindOfDrive": "",
        "Mileage": 0,
        "PowerHp": 0,
        "PowerKw": 0,
        "Co2emission": 0,
        "Environmentalbadge": 0,
        "EnergyEfficiencyClass": "0",
        "NoOfSeats": 5,
        "NoOfDoors": 5,
        "NoOfGears": 1,
        "CapacityCylinder": 0,
        "Equipment": [
            "Tires",
            "Seats",
            "",
            ""
        ],
        "Color": "",
        "ColorFamily": null,
        "ColorCode": null,
        "CarNavi": "",
        "Clutch": false,
        "Consumption": null,
        "CarStatus": null,
        "CarType": "",
        "Uniquenumber": 0,
        "CarUsed": false,
        "CurrentLocation": null,
        "Station": null,
        "TenantIdent": null,
        "ObjNo": null,
        "ExternalSourceId": null,
        "PriceProducer1": 0.0,
        "PriceProducer1Net": 0.0,
        "OfferTypeTarifgroup": null,
        "InsuranceNo": null,
        "VehicleProviderIdent": null,
        "PriceTemplateTarifgroup": null,
        "ExtNumber": null,
        "LastGreatinspDate": null,
        "LastSmallinspDate": null,
        "ExternalTariffGroup": null,
        "Tyretype": "",
        "PurchasingConditionsPriceNet": null,
        "BaseMarkupPrice": null,
        "MarkupPrices": null,
        "DefaultRuntimeFlexibility": false,
        "Ident": ""
    },
    "Bookmarked": false,
    "MobilityBudget": null,
    "CO2Budget": null,
    "AdditionalPayment": 0.0,
    "IsPayback": false,
    "MobilityAdditionalPayment": 0.0,
    "IsMobilityPayback": false,
    "CO2AdditionalPayment": 0.0,
    "IsCO2Payback": false,
    "MinMileage": null,
    "MaxMileage": null,
    "CreatedDate": "",
    "ExternalTariffGroup": null,
    "TarifElementIdents": null,
    "ImageIdents": [
        "",
        "",
        ""
    ],
    "ImageLinks": [
        "https://imgl.krone.at/scaled/2347804/v0780ce/og_image.jpg",
        "https://imgl.krone.at/scaled/2347804/v0780ce/og_image.jpg",
        "https://imgl.krone.at/scaled/2347804/v0780ce/og_image.jpg"
    ],
    "OfferTypeToken": "",
    "OfferTypeName": "",
    "Active": true,
    "Archived": false,
    "Tyretype": "",
    "Color": "",
    "KindOfGear": "",
    "KindOfFuel": "",
    "Price": 0.0,
    "PriceGr": 0.0,
    "UserdefinedOutfit2": "",
    "MinimumAge": 18,
    "ShopInsuranceIdent": null,
    "ShopServiceRegistrationIdent": null,
    "MileagePerMonth": 0,
    "RuntimeMonths": 0,
    "TextblockDeliveryAdmission": null,
    "TextblockMiscTerms": "",
    "TextblockInsuranceTerms": "",
    "MileageExtraPrice": 0.0,
    "MileageExtraPriceGr": 0.0,
    "ActiveInLzs": null,
    "TextblockInsurance": null,
    "TextblockDelivery": null,
    "ColorCode": null,
    "MileageCount": null,
    "MaturityDays": null,
    "Maturity": null,
    "Mileage": null,
    "MileageLevel2Count": null,
    "MileageExtraLevel2Price": null,
    "MileageExtraLevel2PriceGr": null,
    "UserdefinedOutfit1": null,
    "Engine": null,
    "NaviIncluded": null,
    "CarSale": null,
    "PriceUnit": null,
    "AdversityLimit": null,
    "Topoffer": null,
    "DeliveryBeginLocation": null,
    "DeliveryEndLocation": null,
    "ExceedenceLevel1Days": null,
    "ExceedenceLevel1Price": null,
    "ExceedenceLevel1PriceGr": null,
    "ExceedenceLevel2Days": null,
    "ExceedenceLevel2Price": null,
    "ExceedenceLevel2PriceGr": null,
    "InsurancePriceFixedRate": null,
    "InsurancePrice": null,
    "InsurancePriceGr": null,
    "PriceInformation": null,
    "DeliveryPrice": null,
    "DeliveryPriceGr": null,
    "DeliveryPriceFixedRate": null,
    "CostSharingTK": null,
    "CostSharingVK": null,
    "ProductGroup": null,
    "CostPlace": null,
    "ExternalSourceId": null,
    "UseOfferTypePriceTemplate": true,
    "MappedTarifgroup": null,
    "SalaryWaiver": 0.0,
    "IsFlexOffer": false,
    "Ident": ""
  }
```

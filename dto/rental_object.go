package dto

import "time"

type RentalObject struct {
	ReadonlyProperties      ReadonlyProperties `json:"ReadonlyProperties"`
	ChassisNo               interface{}        `json:"ChassisNo"`
	DistMark                interface{}        `json:"DistMark"`
	ObjGroup                interface{}        `json:"ObjGroup"`
	Tarifgroup              interface{}        `json:"Tarifgroup"`
	CarLabel                string             `json:"CarLabel"`
	CarModell               string             `json:"CarModell"`
	CarModellspec           string             `json:"CarModellspec"`
	DateLicense             interface{}        `json:"DateLicense"`
	DateRegistration        time.Time          `json:"DateRegistration"`
	DateUnregister          time.Time          `json:"DateUnregister"`
	DateBuy                 interface{}        `json:"DateBuy"`
	KindOfGear              string             `json:"KindOfGear"`
	KindOfFuel              string             `json:"KindOfFuel"`
	KindOfDrive             string             `json:"KindOfDrive"`
	Mileage                 int                `json:"Mileage"`
	PowerHp                 int                `json:"PowerHp"`
	PowerKw                 int                `json:"PowerKw"`
	Co2Emission             int                `json:"Co2emission"`
	Environmentalbadge      int                `json:"Environmentalbadge"`
	EnergyEfficiencyClass   interface{}        `json:"EnergyEfficiencyClass"`
	NoOfSeats               int                `json:"NoOfSeats"`
	NoOfDoors               int                `json:"NoOfDoors"`
	NoOfGears               interface{}        `json:"NoOfGears"`
	CapacityCylinder        CapacityCylinder   `json:"CapacityCylinder"`
	Equipment               []string           `json:"Equipment"`
	Color                   string             `json:"Color"`
	ColorFamily             string             `json:"ColorFamily"`
	ColorCode               interface{}        `json:"ColorCode"`
	CarNavi                 string             `json:"CarNavi"`
	Consumption             interface{}        `json:"Consumption"`
	CarStatus               interface{}        `json:"CarStatus"`
	CarType                 string             `json:"CarType"`
	Uniquenumber            int                `json:"Uniquenumber"`
	CarUsed                 bool               `json:"CarUsed"`
	CurrentLocation         interface{}        `json:"CurrentLocation"`
	Station                 interface{}        `json:"Station"`
	TenantIdent             interface{}        `json:"TenantIdent"`
	ObjNo                   interface{}        `json:"ObjNo"`
	ExternalSourceID        interface{}        `json:"ExternalSourceId"`
	PriceProducer1          float64            `json:"PriceProducer1"`
	PriceProducer1Net       float64            `json:"PriceProducer1Net"`
	OfferTypeTarifgroup     interface{}        `json:"OfferTypeTarifgroup"`
	InsuranceNo             interface{}        `json:"InsuranceNo"`
	VehicleProviderIdent    interface{}        `json:"VehicleProviderIdent"`
	PriceTemplateTarifgroup interface{}        `json:"PriceTemplateTarifgroup"`
	ExtNumber               interface{}        `json:"ExtNumber"`
	LastGreatinspDate       interface{}        `json:"LastGreatinspDate"`
	LastSmallinspDate       interface{}        `json:"LastSmallinspDate"`
	WinterTyres             interface{}        `json:"WinterTyres"`
	ExternalTariffGroup     interface{}        `json:"ExternalTariffGroup"`
	Ident                   string             `json:"Ident"`
}

package dto

type GroupedResponse struct {
	Filters    Filters `json:"Filters"`
	FilterText string  `json:"FilterText"`
	Items      []Item  `json:"Items"`
	ItemCount  int64   `json:"ItemCount"`
}

type Filters struct {
	Power                AvailableMonths `json:"power"`
	Price                AvailableMonths `json:"price"`
	MileagePerMonth      AvailableMonths `json:"mileagePerMonth"`
	Colors               []Brand         `json:"colors"`
	ColorFamilies        []Brand         `json:"colorFamilies"`
	NoOfDoors            []NoOfDoor      `json:"noOfDoors"`
	CapacityCylinder     AvailableMonths `json:"capacityCylinder"`
	Co2Emission          AvailableMonths `json:"co2Emission"`
	PriceProducer        AvailableMonths `json:"priceProducer"`
	KindOfDrives         []Brand         `json:"kindOfDrives"`
	KindOfFuels          []Brand         `json:"kindOfFuels"`
	KindOfGears          []Brand         `json:"kindOfGears"`
	MaxAdditionalPayment int64           `json:"maxAdditionalPayment"`
	MobilityBudget       int64           `json:"mobilityBudget"`
	ElectricBudget       int64           `json:"electricBudget"`
	Co2Budget            int64           `json:"co2Budget"`
	AvailableMonths      AvailableMonths `json:"availableMonths"`
	Runtime              AvailableMonths `json:"runtime"`
	CarTypes             []Brand         `json:"carTypes"`
	Brands               []Brand         `json:"brands"`
	Models               []Brand         `json:"models"`
	ModelSpecs           []Brand         `json:"modelSpecs"`
	CarUsed              []CarUsed       `json:"carUsed"`
	IsFlexOffer          []CarUsed       `json:"isFlexOffer"`
	Clutch               []CarUsed       `json:"clutch"`
	SalaryWaiver         AvailableMonths `json:"salaryWaiver"`
	TyreTypes            []Brand         `json:"tyreTypes"`
	Stations             interface{}     `json:"stations"`
}

type AvailableMonths struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

type Brand struct {
	Ident interface{} `json:"ident"`
	Value string      `json:"value"`
	Count int64       `json:"count"`
}

type CarUsed struct {
	Ident *bool `json:"ident"`
	Value bool  `json:"value"`
	Count int64 `json:"count"`
}

type NoOfDoor struct {
	Ident int64 `json:"ident"`
	Value int64 `json:"value"`
	Count int64 `json:"count"`
}

type Item struct {
	VehiclesCount                int          `json:"VehiclesCount"`
	RentalObject                 RentalObject `json:"RentalObject"`
	Bookmarked                   bool         `json:"Bookmarked"`
	MobilityBudget               interface{}  `json:"MobilityBudget"`
	ElectricBudget               interface{}  `json:"ElectricBudget"`
	CO2Budget                    interface{}  `json:"CO2Budget"`
	AdditionalPayment            int64        `json:"AdditionalPayment"`
	IsPayback                    bool         `json:"IsPayback"`
	MobilityAdditionalPayment    int64        `json:"MobilityAdditionalPayment"`
	IsMobilityPayback            bool         `json:"IsMobilityPayback"`
	CO2AdditionalPayment         int64        `json:"CO2AdditionalPayment"`
	IsCO2Payback                 bool         `json:"IsCO2Payback"`
	FuelTypes                    []KindOfFuel `json:"FuelTypes"`
	Stations                     []Station    `json:"Stations"`
	MinMileage                   *int64       `json:"MinMileage"`
	MaxMileage                   *int64       `json:"MaxMileage"`
	CreatedDate                  MyTime       `json:"CreatedDate"`
	ExternalTariffGroup          interface{}  `json:"ExternalTariffGroup"`
	TarifElementIdents           interface{}  `json:"TarifElementIdents"`
	ImageIdents                  []string     `json:"ImageIdents"`
	ImageLinks                   []string     `json:"ImageLinks"`
	Active                       bool         `json:"Active"`
	Archived                     bool         `json:"Archived"`
	Topoffer                     interface{}  `json:"Topoffer"`
	InsurancePriceFixedRate      interface{}  `json:"InsurancePriceFixedRate"`
	DeliveryPriceFixedRate       interface{}  `json:"DeliveryPriceFixedRate"`
	UseOfferTypePriceTemplate    bool         `json:"UseOfferTypePriceTemplate"`
	MappedTarifgroup             interface{}  `json:"MappedTarifgroup"`
	SalaryWaiver                 int64        `json:"SalaryWaiver"`
	IsFlexOffer                  bool         `json:"IsFlexOffer"`
	OfferTypeToken               string       `json:"OfferTypeToken"`
	OfferTypeName                string       `json:"OfferTypeName"`
	Tyretype                     Tyretype     `json:"Tyretype"`
	Color                        string       `json:"Color"`
	KindOfGear                   KindOfGear   `json:"KindOfGear"`
	KindOfFuel                   KindOfFuel   `json:"KindOfFuel"`
	Price                        int64        `json:"Price"`
	PriceGr                      float64      `json:"PriceGr"`
	UserdefinedOutfit2           string       `json:"UserdefinedOutfit2"`
	MinimumAge                   *int64       `json:"MinimumAge"`
	ShopInsuranceIdent           interface{}  `json:"ShopInsuranceIdent"`
	ShopServiceRegistrationIdent interface{}  `json:"ShopServiceRegistrationIdent"`
	MileagePerMonth              int64        `json:"MileagePerMonth"`
	RuntimeMonths                int64        `json:"RuntimeMonths"`
	TextblockDeliveryAdmission   interface{}  `json:"TextblockDeliveryAdmission"`
	TextblockMiscTerms           string       `json:"TextblockMiscTerms"`
	TextblockInsuranceTerms      string       `json:"TextblockInsuranceTerms"`
	MileageExtraPrice            float64      `json:"MileageExtraPrice"`
	MileageExtraPriceGr          float64      `json:"MileageExtraPriceGr"`
	ActiveInLzs                  interface{}  `json:"ActiveInLzs"`
	TextblockInsurance           interface{}  `json:"TextblockInsurance"`
	TextblockDelivery            interface{}  `json:"TextblockDelivery"`
	ColorCode                    interface{}  `json:"ColorCode"`
	MetaColor                    interface{}  `json:"MetaColor"`
	MileageCount                 interface{}  `json:"MileageCount"`
	MaturityDays                 interface{}  `json:"MaturityDays"`
	Maturity                     interface{}  `json:"Maturity"`
	Mileage                      interface{}  `json:"Mileage"`
	MileageLevel2Count           interface{}  `json:"MileageLevel2Count"`
	MileageExtraLevel2Price      interface{}  `json:"MileageExtraLevel2Price"`
	MileageExtraLevel2PriceGr    interface{}  `json:"MileageExtraLevel2PriceGr"`
	UserdefinedOutfit1           interface{}  `json:"UserdefinedOutfit1"`
	Engine                       interface{}  `json:"Engine"`
	NaviIncluded                 interface{}  `json:"NaviIncluded"`
	CarSale                      interface{}  `json:"CarSale"`
	PriceUnit                    interface{}  `json:"PriceUnit"`
	AdversityLimit               interface{}  `json:"AdversityLimit"`
	DeliveryBeginLocation        interface{}  `json:"DeliveryBeginLocation"`
	DeliveryEndLocation          interface{}  `json:"DeliveryEndLocation"`
	ExceedenceLevel1Days         interface{}  `json:"ExceedenceLevel1Days"`
	ExceedenceLevel1Price        interface{}  `json:"ExceedenceLevel1Price"`
	ExceedenceLevel1PriceGr      interface{}  `json:"ExceedenceLevel1PriceGr"`
	ExceedenceLevel2Days         interface{}  `json:"ExceedenceLevel2Days"`
	ExceedenceLevel2Price        interface{}  `json:"ExceedenceLevel2Price"`
	ExceedenceLevel2PriceGr      interface{}  `json:"ExceedenceLevel2PriceGr"`
	InsurancePrice               interface{}  `json:"InsurancePrice"`
	InsurancePriceGr             interface{}  `json:"InsurancePriceGr"`
	PriceInformation             interface{}  `json:"PriceInformation"`
	DeliveryPrice                interface{}  `json:"DeliveryPrice"`
	DeliveryPriceGr              interface{}  `json:"DeliveryPriceGr"`
	CostSharingTK                interface{}  `json:"CostSharingTK"`
	CostSharingVK                interface{}  `json:"CostSharingVK"`
	ProductGroup                 interface{}  `json:"ProductGroup"`
	CostPlace                    interface{}  `json:"CostPlace"`
	ExternalSourceID             interface{}  `json:"ExternalSourceId"`
	Ident                        string       `json:"Ident"`
}

type RentalObject struct {
	ReadonlyProperties           ReadonlyProperties `json:"ReadonlyProperties"`
	ChassisNo                    *string            `json:"ChassisNo"`
	DistMark                     interface{}        `json:"DistMark"`
	ObjGroup                     *ObjGroup          `json:"ObjGroup"`
	Tarifgroup                   *Tarifgroup        `json:"Tarifgroup"`
	CarLabel                     CarLabel           `json:"CarLabel"`
	CarModell                    string             `json:"CarModell"`
	CarModellspec                *string            `json:"CarModellspec"`
	DateLicense                  interface{}        `json:"DateLicense"`
	DateRegistration             MyTime             `json:"DateRegistration"`
	DateUnregister               MyTime             `json:"DateUnregister"`
	DateBuy                      interface{}        `json:"DateBuy"`
	KindOfGear                   KindOfGear         `json:"KindOfGear"`
	KindOfFuel                   KindOfFuel         `json:"KindOfFuel"`
	KindOfDrive                  *KindOfDrive       `json:"KindOfDrive"`
	Mileage                      int64              `json:"Mileage"`
	PowerHP                      int64              `json:"PowerHp"`
	PowerKw                      int64              `json:"PowerKw"`
	PrimaryPower                 int64              `json:"PrimaryPower"`
	Co2Emission                  int64              `json:"Co2emission"`
	Environmentalbadge           *int64             `json:"Environmentalbadge"`
	EnergyEfficiencyClass        *string            `json:"EnergyEfficiencyClass"`
	NoOfSeats                    int64              `json:"NoOfSeats"`
	NoOfDoors                    int64              `json:"NoOfDoors"`
	NoOfGears                    *int64             `json:"NoOfGears"`
	CapacityCylinder             *int64             `json:"CapacityCylinder"`
	Equipment                    []string           `json:"Equipment"`
	Color                        string             `json:"Color"`
	ColorFamily                  *ColorFamily       `json:"ColorFamily"`
	ColorCode                    interface{}        `json:"ColorCode"`
	CarNavi                      *CarNavi           `json:"CarNavi"`
	Clutch                       bool               `json:"Clutch"`
	Consumption                  *int64             `json:"Consumption"`
	CarStatus                    interface{}        `json:"CarStatus"`
	CarType                      CarType            `json:"CarType"`
	Uniquenumber                 int64              `json:"Uniquenumber"`
	CarUsed                      bool               `json:"CarUsed"`
	CurrentLocation              interface{}        `json:"CurrentLocation"`
	Station                      interface{}        `json:"Station"`
	StationIdent                 StationIdent       `json:"StationIdent"`
	ObjNo                        interface{}        `json:"ObjNo"`
	ExternalSourceID             interface{}        `json:"ExternalSourceId"`
	PriceProducer1               float64            `json:"PriceProducer1"`
	PriceProducer1Net            *float64           `json:"PriceProducer1Net"`
	OfferTypeTarifgroup          interface{}        `json:"OfferTypeTarifgroup"`
	InsuranceNo                  interface{}        `json:"InsuranceNo"`
	VehicleProviderIdent         interface{}        `json:"VehicleProviderIdent"`
	PriceTemplateTarifgroup      interface{}        `json:"PriceTemplateTarifgroup"`
	LastGreatinspDate            interface{}        `json:"LastGreatinspDate"`
	LastSmallinspDate            interface{}        `json:"LastSmallinspDate"`
	Tyretype                     Tyretype           `json:"Tyretype"`
	PurchasingConditionsPriceNet interface{}        `json:"PurchasingConditionsPriceNet"`
	BaseMarkupPrice              interface{}        `json:"BaseMarkupPrice"`
	MarkupPrices                 interface{}        `json:"MarkupPrices"`
	TenantIdent                  interface{}        `json:"TenantIdent"`
	NextTechcontrol              interface{}        `json:"NextTechcontrol"`
	NextGreatinsp                interface{}        `json:"NextGreatinsp"`
	NextGreatinspKM              interface{}        `json:"NextGreatinspKm"`
	MaxMileage                   int64              `json:"MaxMileage"`
	FormerMileage                int64              `json:"FormerMileage"`
	EXTNumber                    interface{}        `json:"ExtNumber"`
	CarTorent                    interface{}        `json:"CarTorent"`
	DefaultRuntimeFlexibility    interface{}        `json:"DefaultRuntimeFlexibility"`
	Ident                        string             `json:"Ident"`
}

type ReadonlyProperties struct {
	ContractState               interface{} `json:"ContractState"`
	HasFollowupBooking          bool        `json:"HasFollowupBooking"`
	State                       interface{} `json:"State"`
	OfferTypeTarifgroupText     interface{} `json:"OfferTypeTarifgroupText"`
	VehicleProviderName         interface{} `json:"VehicleProviderName"`
	OfferTypeTarifgroupName     interface{} `json:"OfferTypeTarifgroupName"`
	OfferTypeExternalTarifgroup interface{} `json:"OfferTypeExternalTarifgroup"`
}

type Station struct {
	StationIdent      StationIdent `json:"StationIdent"`
	RentalObjectIdent string       `json:"RentalObjectIdent"`
	OfferIdent        string       `json:"OfferIdent"`
}

type KindOfFuel string

const (
	Benzin       KindOfFuel = "Benzin"
	Diesel       KindOfFuel = "Diesel"
	Elektro      KindOfFuel = "Elektro"
	PlugInHybrid KindOfFuel = "Plug-in-Hybrid"
)

type KindOfGear string

const (
	Automatik KindOfGear = "Automatik"
	Schaltung KindOfGear = "Schaltung"
)

type CarLabel string

const (
	Audi      CarLabel = "Audi"
	Bmw       CarLabel = "BMW"
	Cupra     CarLabel = "CUPRA"
	Hyundai   CarLabel = "Hyundai"
	LandRover CarLabel = "Land Rover"
	Seat      CarLabel = "Seat"
	Vw        CarLabel = "VW"
	VwNfz     CarLabel = "VW NFZ"
)

type CarNavi string

const (
	CarNaviJa CarNavi = "ja"
	Ja        CarNavi = "Ja"
	Nein      CarNavi = "Nein"
)

type CarType string

const (
	GeländewagenPickup CarType = "Geländewagen/pickup"
	Kombi              CarType = "Kombi"
	Kompaktwagen       CarType = "Kompaktwagen"
	Limousine          CarType = "Limousine"
	VanKleinbus        CarType = "Van/kleinbus"
)

type ColorFamily string

const (
	Beige           ColorFamily = "Beige"
	Blau            ColorFamily = "Blau"
	ColorFamilyGrau ColorFamily = "grau"
	Grau            ColorFamily = "Grau"
	Grün            ColorFamily = "Grün"
	Orange          ColorFamily = "Orange"
	Rot             ColorFamily = "Rot"
	Schwarz         ColorFamily = "Schwarz"
	Silber          ColorFamily = "Silber"
	Weiß            ColorFamily = "Weiß"
)

type KindOfDrive string

const (
	Allrad       KindOfDrive = "Allrad"
	Frontantrieb KindOfDrive = "Frontantrieb"
	Heckantrieb  KindOfDrive = "Heckantrieb"
)

type ObjGroup string

const (
	Pkw ObjGroup = "PKW"
	Tra ObjGroup = "TRA"
)

type StationIdent string

const (
	Ccunzele2021011217115535345A31 StationIdent = "CCUNZELE2021011217115535345A31"
)

type Tarifgroup string

const (
	E  Tarifgroup = "E"
	T1 Tarifgroup = "T1"
)

type Tyretype string

const (
	Allwetterbereifung Tyretype = "Allwetterbereifung"
	Sommerreifen       Tyretype = "Sommerreifen"
	The8FachBereifung  Tyretype = "8-Fach Bereifung"
	Winterreifen       Tyretype = "Winterreifen"
)

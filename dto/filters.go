package dto

type Filters struct {
	Power                Power            `json:"power"`
	Price                Price            `json:"price"`
	MileagePerMonth      MileagePerMonth  `json:"mileagePerMonth"`
	CarsFilter           CarsFilter       `json:"carsFilter"`
	Colors               []StringFilter         `json:"colors"`
	ColorFamilies        []StringFilter         `json:"colorFamilies"`
	NoOfDoors            []IntFilter            `json:"noOfDoors"`
	CapacityCylinder     CapacityCylinder `json:"capacityCylinder"`
	Co2Emission          Co2Emission      `json:"co2Emission"`
	PriceProducer        PriceProducer    `json:"priceProducer"`
	KindOfDrives         []StringFilter         `json:"kindOfDrives"`
	KindOfFuels          []StringFilter         `json:"kindOfFuels"`
	KindOfGears          []StringFilter         `json:"kindOfGears"`
	MaxAdditionalPayment int              `json:"maxAdditionalPayment"`
	MobilityBudget       float64          `json:"mobilityBudget"`
	Co2Budget            float64          `json:"co2Budget"`
	AvailableMonths      AvailableMonths  `json:"availableMonths"`
	Runtime              Runtime          `json:"runtime"`
}

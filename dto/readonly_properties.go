package dto

type ReadonlyProperties struct {
	ContractState               interface{} `json:"ContractState"`
	HasFollowupBooking          bool        `json:"HasFollowupBooking"`
	State                       interface{} `json:"State"`
	OfferTypeTarifgroupText     interface{} `json:"OfferTypeTarifgroupText"`
	VehicleProviderName         interface{} `json:"VehicleProviderName"`
	OfferTypeTarifgroupName     interface{} `json:"OfferTypeTarifgroupName"`
	OfferTypeExternalTarifgroup interface{} `json:"OfferTypeExternalTarifgroup"`
}

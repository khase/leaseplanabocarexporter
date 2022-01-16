package dto

type CarsFilter struct {
	UsedCars bool     `json:"usedCars"`
	NewCars  bool     `json:"newCars"`
	CarTypes []string `json:"carTypes"`
	Brands   []string `json:"brands"`
	Models   []string `json:"models"`
}

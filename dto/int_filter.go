package dto

type IntFilter struct {
	Ident  int `json:"Ident"`
	Value int `json:"Value"`
	Count       int `json:"Count"`
}

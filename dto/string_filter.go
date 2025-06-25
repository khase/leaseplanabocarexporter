package dto

type StringFilter struct {
	Ident  string `json:"Ident"`
	Value string `json:"Value"`
	Count       int `json:"Count"`
}

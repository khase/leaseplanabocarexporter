package dto

type CarResponse struct {
	Filters    Filters `json:"Filters"`
	FilterText string  `json:"FilterText"`
	Items      []Item  `json:"Items"`
	ItemCount  int     `json:"ItemCount"`
}

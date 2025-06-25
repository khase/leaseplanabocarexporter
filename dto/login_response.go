package dto

type LoginResponse struct {
	Address              Address       `json:"Address"`
	Newsletter           bool          `json:"Newsletter"`
	AddressToken         string        `json:"AddressToken"`
	PreferredLanguage    string        `json:"PreferredLanguage"`
	StoredPaymentMethods []interface{} `json:"StoredPaymentMethods"`
}

type Address struct {
	Salut         string      `json:"Salut"`
	FirstName     string      `json:"FirstName"`
	LastName      string      `json:"LastName"`
	DateOfBirth   MyTime      `json:"DateOfBirth"`
	Email         string      `json:"Email"`
	Phone         string      `json:"Phone"`
	Street        string      `json:"Street"`
	Zip           string      `json:"Zip"`
	City          string      `json:"City"`
	Nationality   string      `json:"Nationality"`
	Ident         string      `json:"Ident"`
	Company       interface{} `json:"Company"`
	AddressExtras interface{} `json:"AddressExtras"`
}

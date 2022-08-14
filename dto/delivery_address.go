package dto

type DeliveryAddress struct {
	Name1                 string      `json:"Name1"`
	Name2                 interface{} `json:"Name2"`
	Firstname             interface{} `json:"Firstname"`
	Street                string      `json:"Street"`
	Zip                   string      `json:"Zip"`
	City                  string      `json:"City"`
	Phone1                string      `json:"Phone1"`
	Phone2                interface{} `json:"Phone2"`
	Email                 interface{} `json:"Email"`
	Interlocutor          interface{} `json:"Interlocutor"`
	Phoneinterlocutor     interface{} `json:"Phoneinterlocutor"`
	InterlocutorEmail     interface{} `json:"InterlocutorEmail"`
	InterlocutorSalut     interface{} `json:"InterlocutorSalut"`
	InterlocutorFirstname interface{} `json:"InterlocutorFirstname"`
	InterlocutorNotes     interface{} `json:"InterlocutorNotes"`
	Salut                 interface{} `json:"Salut"`
	Employer              string      `json:"Employer"`
	Nationality           interface{} `json:"Nationality"`
	GeoLat                float64     `json:"GeoLat"`
	GeoLong               float64     `json:"GeoLong"`
	Ident                 string      `json:"Ident"`
	ToBeDeleted           interface{} `json:"ToBeDeleted"`
}

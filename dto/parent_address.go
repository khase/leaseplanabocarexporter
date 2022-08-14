package dto

type ParentAddress struct {
	Ident                         string          `json:"Ident"`
	Name1                         string          `json:"Name1"`
	Street                        string          `json:"Street"`
	Zip                           string          `json:"Zip"`
	City                          string          `json:"City"`
	InterlocutorFirstname         string          `json:"InterlocutorFirstname"`
	Interlocutor                  string          `json:"Interlocutor"`
	Phoneinterlocutor             string          `json:"Phoneinterlocutor"`
	InterlocutorEmail             string          `json:"InterlocutorEmail"`
	InterlocutorNotes             string          `json:"InterlocutorNotes"`
	ExtraProperties               ExtraProperties `json:"ExtraProperties"`
	Employer                      interface{}     `json:"Employer"`
	Phone1                        interface{}     `json:"Phone1"`
	HideCompanyRentalPrices       bool            `json:"HideCompanyRentalPrices"`
	UseSalaryWaiver               bool            `json:"UseSalaryWaiver"`
	CompanyPolicyDocumentsEnabled bool            `json:"CompanyPolicyDocumentsEnabled"`
}

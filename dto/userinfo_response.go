package dto

type UserInfo struct {
	PreferredLanguage          string                  `json:"PreferredLanguage"`
	Inactive                   bool                    `json:"Inactive"`
	AddressToken               string                  `json:"AddressToken"`
	CreditratingDate           interface{}             `json:"CreditratingDate"`
	CreditratingIndex          interface{}             `json:"CreditratingIndex"`
	Delivery                   bool                    `json:"Delivery"`
	Extnumber                  string                  `json:"Extnumber"`
	CO2PricePerGramOverpayment int64                   `json:"CO2PricePerGramOverpayment"`
	CO2PricePerGramRepayment   int64                   `json:"CO2PricePerGramRepayment"`
	CO2MaxOverpayment          int64                   `json:"CO2MaxOverpayment"`
	CO2MaxRepayment            int64                   `json:"CO2MaxRepayment"`
	InvoiceReceiver            interface{}             `json:"InvoiceReceiver"`
	ExtraProperties            UserInfoExtraProperties `json:"ExtraProperties"`
	DeliveryAddresses          []DeliveryAddress       `json:"DeliveryAddresses"`
	ParentAddress              ParentAddress           `json:"ParentAddress"`
	Policy                     interface{}             `json:"Policy"`
	OwnedPolicies              interface{}             `json:"OwnedPolicies"`
	AddressRole                AddressRole             `json:"AddressRole"`
	StoredPaymentMethods       []interface{}           `json:"StoredPaymentMethods"`
	Station                    string                  `json:"Station"`
	AdrNo                      string                  `json:"AdrNo"`
	AdrGroup                   interface{}             `json:"AdrGroup"`
	Newsletter                 bool                    `json:"Newsletter"`
	Birth                      string                  `json:"Birth"`
	Licno                      string                  `json:"Licno"`
	Licdate                    MyTime                  `json:"Licdate"`
	Licauthor                  string                  `json:"Licauthor"`
	Licclasses                 string                  `json:"Licclasses"`
	Passno                     interface{}             `json:"Passno"`
	Passdate                   interface{}             `json:"Passdate"`
	Passauthor                 interface{}             `json:"Passauthor"`
	ExternalSourceID           interface{}             `json:"ExternalSourceId"`
	Taxno                      interface{}             `json:"Taxno"`
	Taxid                      interface{}             `json:"Taxid"`
	Ident                      string                  `json:"Ident"`
	ToBeDeleted                interface{}             `json:"ToBeDeleted"`
	Salut                      string                  `json:"Salut"`
	Phone1                     string                  `json:"Phone1"`
	Nationality                string                  `json:"Nationality"`
	Phone2                     interface{}             `json:"Phone2"`
	Email                      string                  `json:"Email"`
	InterlocutorFirstname      interface{}             `json:"InterlocutorFirstname"`
	InterlocutorSalut          interface{}             `json:"InterlocutorSalut"`
	InterlocutorEmail          interface{}             `json:"InterlocutorEmail"`
	Phoneinterlocutor          interface{}             `json:"Phoneinterlocutor"`
	Interlocutor               interface{}             `json:"Interlocutor"`
	InterlocutorNotes          interface{}             `json:"InterlocutorNotes"`
	Employer                   interface{}             `json:"Employer"`
	GeoLat                     float64                 `json:"GeoLat"`
	GeoLong                    float64                 `json:"GeoLong"`
	Name1                      string                  `json:"Name1"`
	Name2                      interface{}             `json:"Name2"`
	Firstname                  string                  `json:"Firstname"`
	Street                     string                  `json:"Street"`
	Zip                        string                  `json:"Zip"`
	City                       string                  `json:"City"`
}

type AddressRole struct {
	RoleName       string      `json:"RoleName"`
	OwnerIdent     interface{} `json:"OwnerIdent"`
	AssignedPolicy interface{} `json:"AssignedPolicy"`
	Ident          string      `json:"Ident"`
}

type DeliveryAddress struct {
	Ident                 string      `json:"Ident"`
	ToBeDeleted           interface{} `json:"ToBeDeleted"`
	Salut                 *string     `json:"Salut"`
	Phone1                string      `json:"Phone1"`
	Nationality           *string     `json:"Nationality"`
	Phone2                interface{} `json:"Phone2"`
	Email                 interface{} `json:"Email"`
	InterlocutorFirstname interface{} `json:"InterlocutorFirstname"`
	InterlocutorSalut     interface{} `json:"InterlocutorSalut"`
	InterlocutorEmail     interface{} `json:"InterlocutorEmail"`
	Phoneinterlocutor     interface{} `json:"Phoneinterlocutor"`
	Interlocutor          interface{} `json:"Interlocutor"`
	InterlocutorNotes     interface{} `json:"InterlocutorNotes"`
	Employer              string      `json:"Employer"`
	GeoLat                float64     `json:"GeoLat"`
	GeoLong               float64     `json:"GeoLong"`
	Name1                 string      `json:"Name1"`
	Name2                 interface{} `json:"Name2"`
	Firstname             *string     `json:"Firstname"`
	Street                string      `json:"Street"`
	Zip                   string      `json:"Zip"`
	City                  string      `json:"City"`
}

type UserInfoExtraProperties struct {
	EmailAddressVerified bool          `json:"EmailAddressVerified"`
	ReservationAllowed   bool          `json:"ReservationAllowed"`
	StateInfoText        interface{}   `json:"StateInfoText"`
	PassCheckedDate      interface{}   `json:"PassCheckedDate"`
	PassValidTo          interface{}   `json:"PassValidTo"`
	PassCheckedBy        interface{}   `json:"PassCheckedBy"`
	PassCheckStatus      string        `json:"PassCheckStatus"`
	LicenseCheckedDate   interface{}   `json:"LicenseCheckedDate"`
	LicenseCheckedBy     interface{}   `json:"LicenseCheckedBy"`
	LicenseCheckStatus   string        `json:"LicenseCheckStatus"`
	CreditratingLabel    interface{}   `json:"CreditratingLabel"`
	CreditratingSource   interface{}   `json:"CreditratingSource"`
	CreditratingDateTime interface{}   `json:"CreditratingDateTime"`
	CreditratingUser     interface{}   `json:"CreditratingUser"`
	CreditratingStatus   interface{}   `json:"CreditratingStatus"`
	ImageLinksPass       []interface{} `json:"ImageLinksPass"`
	MobilityBudget       interface{}   `json:"MobilityBudget"`
	ElectricBudget       interface{}   `json:"ElectricBudget"`
	CO2Budget            interface{}   `json:"CO2Budget"`
}

type ParentAddress struct {
	Ident                         string                       `json:"Ident"`
	Name1                         string                       `json:"Name1"`
	Street                        string                       `json:"Street"`
	Zip                           string                       `json:"Zip"`
	City                          string                       `json:"City"`
	InterlocutorFirstname         string                       `json:"InterlocutorFirstname"`
	Interlocutor                  string                       `json:"Interlocutor"`
	Phoneinterlocutor             string                       `json:"Phoneinterlocutor"`
	InterlocutorEmail             string                       `json:"InterlocutorEmail"`
	ShowCompanyRentalPrices       bool                         `json:"ShowCompanyRentalPrices"`
	UseSalaryWaiver               bool                         `json:"UseSalaryWaiver"`
	CompanyPolicyDocumentsEnabled bool                         `json:"CompanyPolicyDocumentsEnabled"`
	ExtraProperties               ParentAddressExtraProperties `json:"ExtraProperties"`
	Employer                      interface{}                  `json:"Employer"`
	Phone1                        interface{}                  `json:"Phone1"`
}

type ParentAddressExtraProperties struct {
	EmailAddressVerified           bool             `json:"EmailAddressVerified"`
	ReservationAllowed             bool             `json:"ReservationAllowed"`
	CustomerAreaAllowed            bool             `json:"CustomerAreaAllowed"`
	CustomerAreaMode               interface{}      `json:"CustomerAreaMode"`
	SimpleOnlinereservationAllowed bool             `json:"SimpleOnlinereservationAllowed"`
	PassCheckedDate                interface{}      `json:"PassCheckedDate"`
	PassValidTo                    interface{}      `json:"PassValidTo"`
	PassCheckedBy                  interface{}      `json:"PassCheckedBy"`
	PassCheckStatus                interface{}      `json:"PassCheckStatus"`
	LicenseCheckedDate             interface{}      `json:"LicenseCheckedDate"`
	LicenseValidTo                 interface{}      `json:"LicenseValidTo"`
	LicenseCheckedBy               interface{}      `json:"LicenseCheckedBy"`
	LicenseCheckStatus             interface{}      `json:"LicenseCheckStatus"`
	CreditratingLabel              interface{}      `json:"CreditratingLabel"`
	CreditratingSource             interface{}      `json:"CreditratingSource"`
	CreditratingDateTime           interface{}      `json:"CreditratingDateTime"`
	CreditratingUser               interface{}      `json:"CreditratingUser"`
	CreditratingStatus             interface{}      `json:"CreditratingStatus"`
	CreditratingScoreValue         interface{}      `json:"CreditratingScoreValue"`
	CreditratingRisk               interface{}      `json:"CreditratingRisk"`
	CreditratingScoreText          interface{}      `json:"CreditratingScoreText"`
	CreditratingScoreInfoText      interface{}      `json:"CreditratingScoreInfoText"`
	Note                           interface{}      `json:"Note"`
	PaymentCondition               PaymentCondition `json:"PaymentCondition"`
	ImageLinksLicense              []interface{}    `json:"ImageLinksLicense"`
	ImageLinksPass                 []interface{}    `json:"ImageLinksPass"`
	CompanyLogo                    CompanyLogo      `json:"CompanyLogo"`
}

type CompanyLogo struct {
	ImageLink  string `json:"ImageLink"`
	ImageIdent string `json:"ImageIdent"`
	Text       string `json:"Text"`
	Filename   string `json:"Filename"`
	UploadDate MyTime `json:"UploadDate"`
}

type PaymentCondition struct {
	Ident                     interface{} `json:"Ident"`
	Name1                     interface{} `json:"Name1"`
	TextRG                    interface{} `json:"TextRG"`
	Text0                     interface{} `json:"Text0"`
	TextGS                    interface{} `json:"TextGS"`
	Paymentkind               interface{} `json:"Paymentkind"`
	Netdays                   int64       `json:"Netdays"`
	Disdays                   int64       `json:"Disdays"`
	Dispercent                int64       `json:"Dispercent"`
	Warndays1                 int64       `json:"Warndays1"`
	Warndays2                 int64       `json:"Warndays2"`
	Warndays3                 int64       `json:"Warndays3"`
	Tolawyerdays              int64       `json:"Tolawyerdays"`
	Warnextra1                int64       `json:"Warnextra1"`
	Warnextra2                int64       `json:"Warnextra2"`
	Warnextra3                int64       `json:"Warnextra3"`
	BankRate                  int64       `json:"BankRate"`
	Nowarning                 bool        `json:"Nowarning"`
	NoAutoPrintCollectInvoice bool        `json:"NoAutoPrintCollectInvoice"`
	Limit                     int64       `json:"Limit"`
	DeductionInvoice          bool        `json:"DeductionInvoice"`
	InvoiceDay                interface{} `json:"InvoiceDay"`
	CollectInvoices           bool        `json:"CollectInvoices"`
	Dtein                     bool        `json:"Dtein"`
	NopaymentOnAccount        bool        `json:"NopaymentOnAccount"`
	Nosecurity                bool        `json:"Nosecurity"`
}

package dto

type ExtraProperties struct {
	EmailAddressVerified           bool              `json:"EmailAddressVerified"`
	ReservationAllowed             bool              `json:"ReservationAllowed"`
	StateInfoText                  interface{}       `json:"StateInfoText"`
	PassCheckedDate                interface{}       `json:"PassCheckedDate"`
	PassValidTo                    interface{}       `json:"PassValidTo"`
	PassCheckedBy                  interface{}       `json:"PassCheckedBy"`
	PassCheckStatus                *string           `json:"PassCheckStatus"`
	LicenseCheckedDate             interface{}       `json:"LicenseCheckedDate"`
	LicenseValidTo                 interface{}       `json:"LicenseValidTo"`
	LicenseCheckedBy               interface{}       `json:"LicenseCheckedBy"`
	LicenseCheckStatus             *string           `json:"LicenseCheckStatus"`
	CreditratingLabel              interface{}       `json:"CreditratingLabel"`
	CreditratingSource             interface{}       `json:"CreditratingSource"`
	CreditratingDateTime           interface{}       `json:"CreditratingDateTime"`
	CreditratingUser               interface{}       `json:"CreditratingUser"`
	CreditratingStatus             interface{}       `json:"CreditratingStatus"`
	CreditratingScoreValue         interface{}       `json:"CreditratingScoreValue"`
	CreditratingRisk               interface{}       `json:"CreditratingRisk"`
	CreditratingScoreText          interface{}       `json:"CreditratingScoreText"`
	CreditratingScoreInfoText      interface{}       `json:"CreditratingScoreInfoText"`
	ImageLinksLicense              []interface{}     `json:"ImageLinksLicense"`
	ImageLinksPass                 []interface{}     `json:"ImageLinksPass"`
	CompanyLogo                    *CompanyLogo      `json:"CompanyLogo"`
	Budget                         interface{}       `json:"Budget"`
	CO2Budget                      interface{}       `json:"CO2Budget"`
	CustomerAreaAllowed            *bool             `json:"CustomerAreaAllowed,omitempty"`
	CustomerAreaMode               interface{}       `json:"CustomerAreaMode"`
	SimpleOnlinereservationAllowed *bool             `json:"SimpleOnlinereservationAllowed,omitempty"`
	Note                           interface{}       `json:"Note"`
	PaymentCondition               *PaymentCondition `json:"PaymentCondition,omitempty"`
}

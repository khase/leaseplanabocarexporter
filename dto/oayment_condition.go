package dto

type PaymentCondition struct {
	Ident                     interface{} `json:"Ident"`
	Name1                     interface{} `json:"Name1"`
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
	TextRG                    interface{} `json:"TextRG"`
	Text0                     interface{} `json:"Text0"`
	TextGS                    interface{} `json:"TextGS"`
	NoAutoPrintCollectInvoice bool        `json:"NoAutoPrintCollectInvoice"`
	Limit                     int64       `json:"Limit"`
	DeductionInvoice          bool        `json:"DeductionInvoice"`
	InvoiceDay                interface{} `json:"InvoiceDay"`
	CollectInvoices           bool        `json:"CollectInvoices"`
	Dtein                     bool        `json:"Dtein"`
	NopaymentOnAccount        bool        `json:"NopaymentOnAccount"`
	Nosecurity                bool        `json:"Nosecurity"`
}

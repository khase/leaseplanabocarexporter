package dto

type AddressRole struct {
	RoleName       string      `json:"RoleName"`
	OwnerIdent     interface{} `json:"OwnerIdent"`
	AssignedPolicy interface{} `json:"AssignedPolicy"`
	Ident          string      `json:"Ident"`
}

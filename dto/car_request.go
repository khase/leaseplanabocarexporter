package dto

type CarRequest struct {
	Page         int    `json:"page"`
	ItemsPerPage int    `json:"itemsPerPage"`
	OrderBy      string `json:"orderBy"`
	OrderAsc     bool   `json:"orderAsc"`
	Bookmark     bool   `json:"bookmark"`
}

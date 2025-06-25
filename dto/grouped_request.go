package dto

type GroupedRequest struct {
	Page      int    `json:"page"`
	OrderBy   string `json:"orderBy"`
	OrderAsc  bool   `json:"orderAsc"`
	Bookmark  bool   `json:"bookmark"`
	DateField string `json:"dateField"`
}

package models

type Status struct {
	Status StatusResponse `json:"status"`
}

type StatusResponse struct {
	StatusCode      int    `json:"status_code"`
	DescriptionCode string `json:"description_code"`
	Description     string `json:"description"`
}
type TransactionByUser struct {
	User string `json:"from"`
}

package dtos

type AccountRequest struct {
	DocumentNumber string `json:"document_number" binding:"required"`
}

type AccountBalanceResponse struct {
	ID             int64   `json:"id"`
	DocumentNumber string  `json:"document_number"`
	Balance        float64 `json:"balance"`
}

type AccountResponse struct {
	ID             int64 `json:"id"`
	DocumentNumber int64 `json:"document_number"`
}

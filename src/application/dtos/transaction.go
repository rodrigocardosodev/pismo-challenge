package dtos

type TransactionRequest struct {
	AccountID       int64   `json:"account_id" binding:"required"`
	OperationTypeID uint8   `json:"operation_type_id" binding:"required"`
	Amount          float64 `json:"amount" binding:"-"`
}

type TransactionResponse struct {
	ID              int64   `json:"id"`
	AccountID       int64   `json:"account_id"`
	OperationTypeID uint8   `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
	EventDate       string  `json:"event_date"`
}

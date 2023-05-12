package dtos

type TransactionRequest struct {
	AccountID       int64  `json:"account_id" binding:"required"`
	OperationTypeID int8   `json:"operation_type_id" binding:"required"`
	Amount          uint64 `json:"amount" binding:"-"`
}

type TransactionResponse struct {
	ID              int64  `json:"id"`
	AccountID       int64  `json:"account_id"`
	OperationTypeID int8   `json:"operation_type_id"`
	Amount          uint64 `json:"amount"`
	EventDate       string `json:"event_date"`
}

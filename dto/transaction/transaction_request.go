package transactiondto

type TransactionRequest struct {
	OrderID     int    `json:"order_id"`
	UserOrderID int    `json:"user_order_id"`
	Status      string `json:"status" gorm:"type: varchar(255)"`
}

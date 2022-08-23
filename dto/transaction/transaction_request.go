package transactiondto

type TransactionRequest struct {
	ProductId int `gorm:"type: int" json:"product_id" validate:"required"`
	Price     int `gorm:"type: int" json:"price" validate:"required"`
	BuyerId   int `gorm:"type: int" json:"buyer_id" validate:"required"`
}

type UpdateTransactionRequest struct {
	Status    string `json:"status" gorm:"type: varchar(255)"`
	BuyerID   int    `json:"buyer_id"`
	ProductID int    `json:"product_id"`
}

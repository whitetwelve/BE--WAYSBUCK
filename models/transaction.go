package models

import "time"

// User model struct
type Transaction struct {
	ID      int                  `json:"id" gorm:"primary_key:auto_increment"`
	Status  string               `json:"status" gorm:"type: varchar(255)"`
	BuyerID int                  `json:"buyer_id"`
	Buyer   UsersProfileResponse `json:"buyer"`
	// SellerID  int                        `json:"seller_id"`
	// Seller    UsersProfileResponse       `json:"seller"`
	ProductID int                        `json:"product_id"`
	Product   ProductTransactionResponse `json:"product"`
	Price     int                        `json:"price"`
	CreatedAt time.Time                  `json:"-"`
	UpdatedAt time.Time                  `json:"-"`
}

type TransactionResponse struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Status    string               `json:"status" gorm:"type: varchar(255)"`
	BuyerID   int                  `json:"buyer_id"`
	Buyer     UsersProfileResponse `json:"buyer"`
	ProductID int                  `json:"product_id"`
	Product   ProductResponse      `json:"product"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}

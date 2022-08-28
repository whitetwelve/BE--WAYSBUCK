package models

import "time"

// User model struct
type Transaction struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	BuyerID   int       `json:"buyer_id"`
	Buyer     User      `json:"buyer" `
	Price     int       `json:"price"`
	Carts     []Cart    `json:"cart"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
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

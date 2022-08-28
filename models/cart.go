package models

import "time"

// User model struct
type Cart struct {
	ID            int                 `json:"id" gorm:"PRIMARY_KEY"`
	UserID        int                 `json:"user_id"`
	User          User                `json:"user"`
	ProductID     int                 `json:"product_id"`
	Product       ProductUserResponse `json:"product"`
	TransactionID *int                `json:"transaction_id"`
	Transaction   Transaction         `json:"-"`
	ToppingID     []int               `json:"topping_id" form:"topping_id" gorm:"-"`
	Topping       []Toping            `json:"topping" gorm:"many2many:cart_topping"`
	SubTotal      int                 `json:"sub_total"`
	QTY           int                 `json:"qty"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"-"`
}

type CartResponse struct {
	ID        int       `json:"id"`
	SubTotal  int       `json:"sub_total"`
	CreatedAt time.Time `json:"created_at"`
}

func (CartResponse) TableName() string {
	return "carts"
}

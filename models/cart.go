package models

import "time"

// User model struct
type Cart struct {
	ID          int                 `json:"id" gorm:"PRIMARY_KEY"`
	UserID      int                 `json:"user_id"`
	ProductID   int                 `json:"product_id"`
	Product     ProductUserResponse `json:"product"`
	Qty         int                 `json:"qty" form:"qty"`
	Image       string              `json:"image"`
	ProductName string              `json:"product_name"`
	SubAmount   int                 `json:"subamount"`
	TopingID    []int               `json:"toping_id" form:"toping_id" gorm:"-"`
	Toping      []Toping            `json:"toping" gorm:"many2many:cart_topings"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"-"`
}

type CartResponse struct {
	ID        int                 `json:"id"`
	UserID    int                 `json:"user_id"`
	ProductID int                 `json:"product_id"`
	ToppingID int                 `json:"-"`
	Product   ProductUserResponse `json:"product"`
	Toping    []Toping            `json:"toping"`
	Qty       int                 `json:"qty" form:"qty"`
	SubAmount int                 `json:"subamount"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"-"`
}

func (CartResponse) TableName() string {
	return "carts"
}

package models

import "time"

// User model struct
type Transaction struct {
	ID          int                        `json:"id" gorm:"primary_key:auto_increment"`
	Status      string                     `json:"status" gorm:"type: varchar(255)"`
	UserOrderID int                        `json:"user_order_id"`
	UserOrder   UsersProfileResponse       `json:"user_order"`
	OrderID     int                        `json:"order_id"`
	Order       ProductTransactionResponse `json:"order"`
	CreatedAt   time.Time                  `json:"-"`
	UpdatedAt   time.Time                  `json:"-"`
}

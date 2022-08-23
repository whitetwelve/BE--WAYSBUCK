package models

import "time"

type User struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullname" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	Status    string    `json:"status" gorm:"type: varchar(50)"`
	Address   string    `json:"address"`
	PostCode  string    `json:"post_code"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Image    string `json:"image"`
}

type UsersTransactionResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	PostCode string `json:"post_code"`
}
type UserBuyerResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}

func (UserBuyerResponse) TableName() string {
	return "users"
}
func (UsersTransactionResponse) TableName() string {
	return "users"
}

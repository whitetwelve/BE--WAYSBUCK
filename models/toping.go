package models

import "time"

type Toping struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Title     string    `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price     int       `json:"price" form:"price" gorm:"type: int"`
	Image     string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	ProductID int       `json:"product_id" form:"product_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type TopingResponse struct {
	ID        int             `json:"id"`
	Title     string          `json:"title"`
	Price     int             `json:"price"`
	Image     string          `json:"image"`
	ProductID int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
}

type TopingProductResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Price     int    `json:"price"`
	Image     string `json:"image"`
	ProductID int    `json:"product_id"`
}

func (TopingResponse) TableName() string {
	return "toping"
}

func (TopingProductResponse) TableName() string {
	return "toping"
}

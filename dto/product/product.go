package productdto

type CreateProduct struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Image string `json:"image" form:"image"`
}

type UpdateProduct struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" form:"price"`
	Image string `json:"image" form:"id"`
}

type ProductResponse struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" form:"price"`
	Image string `json:"image" form:"image"`
}

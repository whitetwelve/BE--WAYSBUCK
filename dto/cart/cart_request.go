package cartdto

type CartRequest struct {
	ID            int    `json:"id"`
	ProductID     int    `json:"product_id"`
	TransactionID int    `json:"transaction_id"`
	Qty           int    `json:"qty" form:"qty"`
	Image         string `json:"image" form:"image"`
	ProductName   string `json:"product_name" form:"product_name"`
	SubTotal      int    `json:"sub_total"`
	ToppingID     []int  `json:"topping_id" form:"topping_id" gorm:"-"`
	UserID        int    `json:"user_id" gorm:"type: int"`
}

type CreateCart struct {
	ProductID int   `json:"product_id" form:"product_id"`
	ToppingID []int `json:"topping_id" form:"topping_id"`
	SubTotal  int   `json:"sub_total" form:"sub_total"`
	// Qty       int    `json:"qty"`
	// Status string `json:"status"`
}

type UpdateCart struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	ToppingID int `json:"topping_id"`
}

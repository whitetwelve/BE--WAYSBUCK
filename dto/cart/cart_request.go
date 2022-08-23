package cartdto

type CartRequest struct {
	ID            int    `json:"id"`
	ProductID     int    `json:"product_id"`
	TransactionID int    `json:"transaction_id"`
	Qty           int    `json:"qty" form:"qty"`
	Image         string `json:"image" form:"image"`
	ProductName   string `json:"product_name" form:"product_name"`
	SubAmount     int    `json:"subamount"`
	TopingID      []int  `json:"toping_id" form:"toping_id" gorm:"-"`
	UserID        int    `json:"user_id" gorm:"type: int"`
}

type CreateCart struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	TopingID  []int  `json:"toping_id"`
	QTY       int    `json:"qty"`
	SubAmount int    `json:"subamount"`
	Status    string `jsom:"status"`
}

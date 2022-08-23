package authdto

type LoginResponse struct {
	ID       int    `gorm:"type: int" json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	PostCode string `gorm:"type: varchar(255)" json:"post_code"`
	Address  string `gorm:"type: varchar(255)" json:"address"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Status   string `gorm:"type: varchar(50)"  json:"status"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
}

type RegisterResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
}

type CheckAuthResponse struct {
	Id       int    `gorm:"type: int" json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Status   string `gorm:"type: varchar(50)"  json:"status"`
}

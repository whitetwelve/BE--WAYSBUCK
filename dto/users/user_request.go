package usersdto

type CreateUserRequest struct {
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
	Status   string `json:"status" form:"status" validate:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Image    string `json:"image" form:"image"`
	Address  string `json:"address" form:"address"`
	PostCode string `json:"post_code" form:"post_code"`
	Status   string `json:"status" form:"status"`
}

package payload

type CreateUserRequest struct {
	Username string `json:"username" form:"username" validate:"required,min=5,max=15"`
	Name     string `json:"name" form:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" form:"password" validate:"required,min=5,max=15"`
}

type CreateOrLoginUserResponse struct {
	Username   string `json:"username"`
	Name       string `json:"name"`
	AccesToken string `json:"acces_token"`
}

type LoginUserRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

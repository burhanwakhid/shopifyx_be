package request

type RegisterRequest struct {
	Username string `json:"username" validate:"required,string,min=5,max=15"`
	Name     string `json:"name" validate:"required,string,min=5,max=50"`
	Password string `json:"password" validate:"required,string,min=5,max=15"`
	// `validate:"required,min=4,max=15"`
}

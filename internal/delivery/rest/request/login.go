package request

type LoginRequest struct {
	Username string `json:"username" validate:"required,string,min=5,max=15"`
	Password string `json:"password" validate:"required,string,min=5,max=15"`
}

// "username": "seseorang", // not null, minLength 5, maxLength 15
// 	"password": "" // not null, minLength 5, maxLength 15

// - `404` if user not found
// - `400` if password is wrong
// - `400` if password / username is too short or long
// - `500` if server error

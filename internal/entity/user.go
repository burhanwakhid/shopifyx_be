package entity

import (
	"github.com/burhanwakhid/shopifyx_backend/pkg/token"
)

type User struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Name     string `db:"name"`
}

type LoginData struct {
	ID       string
	Username string
	Name     string
	Token    string
}

func (u *User) ToLoginData() *LoginData {
	return &LoginData{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
		Token:    u.generateJwtToken(),
	}
}

func (u *User) generateJwtToken() string {
	t, err := token.GenerateJwt(u.Username, u.ID)

	if err != nil {
		return ""
	}

	return t
}

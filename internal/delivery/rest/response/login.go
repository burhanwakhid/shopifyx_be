package response

type LoginResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"accessToken"`
}

package requests

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

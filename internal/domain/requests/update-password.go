package requests

type UpdatePassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

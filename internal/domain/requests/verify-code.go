package requests

type VerifyCode struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

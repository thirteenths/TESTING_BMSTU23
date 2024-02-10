package domain

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Hash  string `json:"hash"`
	Code  string `json:"code"`
}

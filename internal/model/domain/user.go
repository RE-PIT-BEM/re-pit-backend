package domain

type User struct {
	ID         int    `json:"string"`
	NIM        string `json:"NIM"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Department string `json:"department"`
	Role       string `json:"role"`
}

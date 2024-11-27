package dto

type LoginRequestDTO struct {
	NIM      string `json:"nim" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	Token string               `json:"token"`
	User  AuthorizeResponseDTO `json:"user"`
}

type AuthorizeResponseDTO struct {
	NIM        string `json:"nim"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Major      string `json:"major"`
	Role       string `json:"role"`
}

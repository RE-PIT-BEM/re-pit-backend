package dto

type LoginRequestDTO struct {
	NIM      string `json:"NIM" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// type LoginResponseDTO struct {
// 	JWT      string `json:"signing_key" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

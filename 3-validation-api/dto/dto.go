package dto

type SendDto struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyDto struct {
	Email string `json:"email" validate:"required,email"`
	Hash  string `json:"hash" validate:"required"`
}

package dto

type Checkout struct {
	Id      uint32 `json:"id"`
	Address string `json:"address" validate:"required"`
	Email   string `json:"email" validate:"requried,email"`
}

package dto

type Checkout struct {
	Id      uint32 `json:"id"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

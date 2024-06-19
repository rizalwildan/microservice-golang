package types

type MerchantDTO struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description"`
	City        string `json:"city" validate:"required"`
	UserID      uint   `json:"user_id" validate:"required"`
}

type MerchantResponse struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	City        string     `json:"city"`
	UserID      uint       `json:"user_id"`
	Owner       *User      `json:"owner"`
	Products    []*Product `json:"products"`
}

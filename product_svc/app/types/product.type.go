package types

type ProductDTO struct {
	Name        string  `json:"name" validate:"required,min=3"`
	Description string  `json:"description"`
	Quantity    float32 `json:"quantity"`
	Price       float32 `json:"price" validate:"required"`
	MerchantID  uint    `json:"merchant_id" validate:"required"`
}

type ProductResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    float32 `json:"quantity"`
	Price       float32 `json:"price"`
	MerchantID  uint    `json:"merchant_id"`
}

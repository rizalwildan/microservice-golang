package types

type Merchant struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type MerchantApiResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    Merchant `json:"data"`
}

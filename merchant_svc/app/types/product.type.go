package types

type Product struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
}

type ProductResponseAPI struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}

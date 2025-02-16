package dto

type CreateProductDTO struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
}

type UpdateProductDTO struct {
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}

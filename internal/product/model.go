package product

import "time"

type Product struct {
	ID              string    `json:"id"`
	StoreID         string    `json:"store_id"`
	Name            string    `json:"name"`
	Description     *string   `json:"description"`
	BasePrice       float64   `json:"base_price"`
	PricingStrategy string    `json:"pricing_strategy"`
	Attributes      string    `json:"attributes"`
	Images          string    `json:"images"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Name            string  `json:"name" binding:"required"`
	Description     string  `json:"description"`
	BasePrice       float64 `json:"base_price" binding:"required,gt=0"`
	PricingStrategy string  `json:"pricing_strategy" binding:"required,oneof=fixed weight_based variant_based bundle"`
	Attributes      string  `json:"attributes"`
	Images          string  `json:"images"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	BasePrice   float64 `json:"base_price"`
	IsActive    bool    `json:"is_active"`
}
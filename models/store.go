package models

import "time"

type Store struct {
	ID           string     `json:"id"`
	SellerID     string     `json:"seller_id"`
	Name         string     `json:"name"`
	Slug         string     `json:"slug"`
	Description  *string    `json:"description"`
	LogoURL      *string    `json:"logo_url"`
	BannerURL    *string    `json:"banner_url"`
	VerticalType string     `json:"vertical_type"`
	Settings     string     `json:"settings"`
	ThemeID      *string    `json:"theme_id"`
	ThemeConfig  string     `json:"theme_config"`
	IsActive     bool       `json:"is_active"`
	IsPublished  bool       `json:"is_published"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type CreateStoreRequest struct {
	Name         string  `json:"name" binding:"required"`
	Slug         string  `json:"slug" binding:"required"`
	Description  string  `json:"description"`
	VerticalType string  `json:"vertical_type" binding:"required,oneof=pod restaurant grocery tech kit"`
}

type UpdateStoreRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	LogoURL     string  `json:"logo_url"`
	BannerURL   string  `json:"banner_url"`
}
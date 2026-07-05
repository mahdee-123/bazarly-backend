package models

import "time"

type Seller struct {
	ID                  string     `json:"id"`
	Name                string     `json:"name"`
	Email               string     `json:"email"`
	Phone               *string    `json:"phone"`
	AvatarURL           *string    `json:"avatar_url"`
	PasswordHash        *string    `json:"-"` // json e kabho dekhabe na
	EmailVerified       bool       `json:"email_verified"`
	VerificationToken   *string    `json:"-"`
	TokenExpiresAt      *time.Time `json:"-"`
	ResetToken          *string    `json:"-"`
	ResetTokenExpires   *time.Time `json:"-"`
	SubscriptionTier    string     `json:"subscription_tier"`
	SubscriptionExpires *time.Time `json:"subscription_expires"`
	LastLoginAt         *time.Time `json:"last_login_at"`
	IsActive            bool       `json:"is_active"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

// Signup request
type SellerSignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required,min=6"`
}

// Login request
type SellerLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Response (password hash chhara)
type SellerResponse struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	Email            string     `json:"email"`
	Phone            *string    `json:"phone"`
	AvatarURL        *string    `json:"avatar_url"`
	EmailVerified    bool       `json:"email_verified"`
	SubscriptionTier string     `json:"subscription_tier"`
	CreatedAt        time.Time  `json:"created_at"`
}
package sellers

import (
	"errors"
	"time"

	"github.com/mahdee-123/bazarly-backend/internal/services"
	"github.com/mahdee-123/bazarly-backend/internal/utils"
)

func SignupSeller(req SellerSignupRequest) (*SellerResponse, error) {
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("password hashing failed")
	}

	newSeller, err := CreateSellerRepo(req, hash)
	if err != nil {
		return nil, errors.New("email already exists")
	}

	token, err := utils.GenerateRandomToken()
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	expiresAt := time.Now().Add(24 * time.Hour)
	if err := SaveVerificationTokenRepo(newSeller.ID, token, expiresAt); err != nil {
		return nil, errors.New("failed to save verification token")
	}

	go services.SendVerificationEmail(newSeller.Email, newSeller.Name, token)

	return newSeller, nil
}

func LoginSeller(req SellerLoginRequest) (*SellerResponse, string, error) {
	s, err := GetSellerByEmailRepo(req.Email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	if !s.IsActive {
		return nil, "", errors.New("account is disabled")
	}

	if !s.EmailVerified {
		return nil, "", errors.New("please verify your email first")
	}

	if s.PasswordHash == nil {
		return nil, "", errors.New("invalid email or password")
	}

	if !utils.CheckPassword(req.Password, *s.PasswordHash) {
		return nil, "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(s.ID)
	if err != nil {
		return nil, "", errors.New("token generation failed")
	}

	if err := UpdateLastLoginRepo(s.ID); err != nil {
		return nil, "", errors.New("failed to update last login")
	}

	response := &SellerResponse{
		ID:               s.ID,
		Name:             s.Name,
		Email:            s.Email,
		Phone:            s.Phone,
		EmailVerified:    s.EmailVerified,
		SubscriptionTier: s.SubscriptionTier,
		CreatedAt:        s.CreatedAt,
	}

	return response, token, nil
}

func GetSellerProfile(sellerID string) (*SellerResponse, error) {
	s, err := GetSellerByIDRepo(sellerID)
	if err != nil {
		return nil, errors.New("seller not found")
	}
	return s, nil
}

func VerifySellerEmail(token string) error {
	return VerifyEmailRepo(token)
}
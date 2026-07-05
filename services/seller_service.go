package services

import (
	"errors"

	"github.com/mahdee-123/bazarly-backend/models"
	"github.com/mahdee-123/bazarly-backend/repository"
	"github.com/mahdee-123/bazarly-backend/utils"
)

func SignupSeller(req models.SellerSignupRequest) (*models.SellerResponse, error) {
	// Password hash
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("password hashing failed")
	}

	// Database e insert
	seller, err := repository.CreateSeller(req, hash)
	if err != nil {
		return nil, errors.New("email already exists")
	}

	return seller, nil
}

func LoginSeller(req models.SellerLoginRequest) (*models.SellerResponse, string, error) {
	// Email diye seller khojo
	seller, err := repository.GetSellerByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// Account active?
	if !seller.IsActive {
		return nil, "", errors.New("account is disabled")
	}

	// Password check
	if seller.PasswordHash == nil {
		return nil, "", errors.New("invalid email or password")
	}

	if !utils.CheckPassword(req.Password, *seller.PasswordHash) {
		return nil, "", errors.New("invalid email or password")
	}

	// JWT token generate
	token, err := utils.GenerateToken(seller.ID)
	if err != nil {
		return nil, "", errors.New("token generation failed")
	}

	// Last login update
	repository.UpdateLastLogin(seller.ID)

	// Response
	response := &models.SellerResponse{
		ID:               seller.ID,
		Name:             seller.Name,
		Email:            seller.Email,
		Phone:            seller.Phone,
		EmailVerified:    seller.EmailVerified,
		SubscriptionTier: seller.SubscriptionTier,
		CreatedAt:        seller.CreatedAt,
	}

	return response, token, nil
}

func GetSellerProfile(sellerID string) (*models.SellerResponse, error) {
	seller, err := repository.GetSellerByID(sellerID)
	if err != nil {
		return nil, errors.New("seller not found")
	}
	return seller, nil
}
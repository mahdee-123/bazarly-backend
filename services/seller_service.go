package services

import (
	"errors"
	"time"

	"github.com/mahdee-123/bazarly-backend/models"
	"github.com/mahdee-123/bazarly-backend/repository"
	"github.com/mahdee-123/bazarly-backend/utils"
)

func SignupSeller(req models.SellerSignupRequest) (*models.SellerResponse, error) {
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("password hashing failed")
	}

	seller, err := repository.CreateSeller(req, hash)
	if err != nil {
		return nil, errors.New("email already exists")
	}

	token, err := utils.GenerateRandomToken()
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	expiresAt := time.Now().Add(24 * time.Hour)
	err = repository.SaveVerificationToken(seller.ID, token, expiresAt)
	if err != nil {
		return nil, errors.New("failed to save verification token")
	}

	go SendVerificationEmail(seller.Email, seller.Name, token)

	return seller, nil
}

func LoginSeller(req models.SellerLoginRequest) (*models.SellerResponse, string, error) {
	seller, err := repository.GetSellerByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	if !seller.IsActive {
		return nil, "", errors.New("account is disabled")
	}

	if !seller.EmailVerified {
		return nil, "", errors.New("please verify your email first")
	}

	if seller.PasswordHash == nil {
		return nil, "", errors.New("invalid email or password")
	}

	if !utils.CheckPassword(req.Password, *seller.PasswordHash) {
		return nil, "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(seller.ID)
	if err != nil {
		return nil, "", errors.New("token generation failed")
	}

	repository.UpdateLastLogin(seller.ID)

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

func VerifySellerEmail(token string) error {
	return repository.VerifyEmail(token)
}
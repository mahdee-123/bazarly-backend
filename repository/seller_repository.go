package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mahdee-123/bazarly-backend/db"
	"github.com/mahdee-123/bazarly-backend/models"
)

func CreateSeller(req models.SellerSignupRequest, passwordHash string) (*models.SellerResponse, error) {
	var seller models.SellerResponse

	err := db.DB.QueryRow(`
		INSERT INTO sellers (name, email, phone, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, phone, email_verified, subscription_tier, created_at
	`, req.Name, req.Email, req.Phone, passwordHash).Scan(
		&seller.ID,
		&seller.Name,
		&seller.Email,
		&seller.Phone,
		&seller.EmailVerified,
		&seller.SubscriptionTier,
		&seller.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &seller, nil
}

func GetSellerByEmail(email string) (*models.Seller, error) {
	var seller models.Seller

	err := db.DB.QueryRow(`
		SELECT id, name, email, phone, password_hash, 
		       email_verified, subscription_tier, 
		       is_active, created_at, updated_at
		FROM sellers
		WHERE email = $1
		AND deleted_at IS NULL
	`, email).Scan(
		&seller.ID,
		&seller.Name,
		&seller.Email,
		&seller.Phone,
		&seller.PasswordHash,
		&seller.EmailVerified,
		&seller.SubscriptionTier,
		&seller.IsActive,
		&seller.CreatedAt,
		&seller.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("seller not found")
	}

	if err != nil {
		return nil, err
	}

	return &seller, nil
}

func GetSellerByID(id string) (*models.SellerResponse, error) {
	var seller models.SellerResponse

	err := db.DB.QueryRow(`
		SELECT id, name, email, phone, avatar_url,
		       email_verified, subscription_tier, created_at
		FROM sellers
		WHERE id = $1
		AND deleted_at IS NULL
	`, id).Scan(
		&seller.ID,
		&seller.Name,
		&seller.Email,
		&seller.Phone,
		&seller.AvatarURL,
		&seller.EmailVerified,
		&seller.SubscriptionTier,
		&seller.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("seller not found")
	}

	if err != nil {
		return nil, err
	}

	return &seller, nil
}

func UpdateLastLogin(id string) error {
	_, err := db.DB.Exec(`
		UPDATE sellers 
		SET last_login_at = now(), updated_at = now()
		WHERE id = $1
	`, id)
	return err
}




func SaveVerificationToken(sellerID, token string, expiresAt time.Time) error {
	_, err := db.DB.Exec(`
		UPDATE sellers
		SET verification_token = $1,
		    token_expires_at = $2,
		    updated_at = now()
		WHERE id = $3
	`, token, expiresAt, sellerID)
	return err
}

func VerifyEmail(token string) error {
	result, err := db.DB.Exec(`
		UPDATE sellers
		SET email_verified = true,
		    verification_token = NULL,
		    token_expires_at = NULL,
		    updated_at = now()
		WHERE verification_token = $1
		AND token_expires_at > now()
		AND deleted_at IS NULL
	`, token)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("invalid or expired token")
	}

	return nil
}
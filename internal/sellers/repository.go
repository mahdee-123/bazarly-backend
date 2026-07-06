package sellers

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mahdee-123/bazarly-backend/internal/db"
)

func CreateSellerRepo(req SellerSignupRequest, passwordHash string) (*SellerResponse, error) {
	var s SellerResponse

	err := db.DB.QueryRow(`
		INSERT INTO sellers (name, email, phone, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, phone, email_verified, subscription_tier, created_at
	`, req.Name, req.Email, req.Phone, passwordHash).Scan(
		&s.ID, &s.Name, &s.Email, &s.Phone, &s.EmailVerified, &s.SubscriptionTier, &s.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func GetSellerByEmailRepo(email string) (*Seller, error) {
	var s Seller

	err := db.DB.QueryRow(`
		SELECT id, name, email, phone, password_hash,
		       email_verified, subscription_tier,
		       is_active, created_at, updated_at
		FROM sellers
		WHERE email = $1
		AND deleted_at IS NULL
	`, email).Scan(
		&s.ID, &s.Name, &s.Email, &s.Phone, &s.PasswordHash,
		&s.EmailVerified, &s.SubscriptionTier, &s.IsActive, &s.CreatedAt, &s.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("seller not found")
	}
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func GetSellerByIDRepo(id string) (*SellerResponse, error) {
	var s SellerResponse

	err := db.DB.QueryRow(`
		SELECT id, name, email, phone, avatar_url,
		       email_verified, subscription_tier, created_at
		FROM sellers
		WHERE id = $1
		AND deleted_at IS NULL
	`, id).Scan(
		&s.ID, &s.Name, &s.Email, &s.Phone, &s.AvatarURL,
		&s.EmailVerified, &s.SubscriptionTier, &s.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("seller not found")
	}
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func UpdateLastLoginRepo(id string) error {
	_, err := db.DB.Exec(`
		UPDATE sellers
		SET last_login_at = now(), updated_at = now()
		WHERE id = $1
	`, id)
	return err
}

func SaveVerificationTokenRepo(sellerID, token string, expiresAt time.Time) error {
	_, err := db.DB.Exec(`
		UPDATE sellers
		SET verification_token = $1,
		    token_expires_at = $2,
		    updated_at = now()
		WHERE id = $3
	`, token, expiresAt, sellerID)
	return err
}

func VerifyEmailRepo(token string) error {
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
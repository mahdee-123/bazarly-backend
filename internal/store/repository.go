package store

import (
	"database/sql"
	"errors"

	"github.com/mahdee-123/bazarly-backend/internal/db"
)

func CreateStoreRepo(sellerID string, req CreateStoreRequest) (*Store, error) {
	var s Store

	err := db.DB.QueryRow(`
		INSERT INTO stores (seller_id, name, slug, description, vertical_type)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, seller_id, name, slug, description, vertical_type,
		          is_active, is_published, created_at, updated_at
	`, sellerID, req.Name, req.Slug, req.Description, req.VerticalType).Scan(
		&s.ID, &s.SellerID, &s.Name, &s.Slug, &s.Description, &s.VerticalType,
		&s.IsActive, &s.IsPublished, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func GetStoresBySellerID(sellerID string) ([]Store, error) {
	rows, err := db.DB.Query(`
		SELECT id, seller_id, name, slug, description, vertical_type,
		       is_active, is_published, created_at, updated_at
		FROM stores
		WHERE seller_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []Store
	for rows.Next() {
		var s Store
		err := rows.Scan(
			&s.ID, &s.SellerID, &s.Name, &s.Slug, &s.Description, &s.VerticalType,
			&s.IsActive, &s.IsPublished, &s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		stores = append(stores, s)
	}

	return stores, nil
}

func GetStoreByID(storeID, sellerID string) (*Store, error) {
	var s Store

	err := db.DB.QueryRow(`
		SELECT id, seller_id, name, slug, description, vertical_type,
		       is_active, is_published, created_at, updated_at
		FROM stores
		WHERE id = $1 AND seller_id = $2 AND deleted_at IS NULL
	`, storeID, sellerID).Scan(
		&s.ID, &s.SellerID, &s.Name, &s.Slug, &s.Description, &s.VerticalType,
		&s.IsActive, &s.IsPublished, &s.CreatedAt, &s.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("store not found")
	}
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func DeleteStoreRepo(storeID, sellerID string) error {
	result, err := db.DB.Exec(`
		UPDATE stores
		SET deleted_at = now()
		WHERE id = $1 AND seller_id = $2 AND deleted_at IS NULL
	`, storeID, sellerID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("store not found")
	}

	return nil
}
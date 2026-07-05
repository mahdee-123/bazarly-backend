package repository

import (
	"database/sql"
	"errors"

	"github.com/mahdee-123/bazarly-backend/db"
	"github.com/mahdee-123/bazarly-backend/models"
)

func CreateStore(sellerID string, req models.CreateStoreRequest) (*models.Store, error) {
	var store models.Store

	err := db.DB.QueryRow(`
		INSERT INTO stores (seller_id, name, slug, description, vertical_type)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, seller_id, name, slug, description, vertical_type, 
		          is_active, is_published, created_at, updated_at
	`, sellerID, req.Name, req.Slug, req.Description, req.VerticalType).Scan(
		&store.ID,
		&store.SellerID,
		&store.Name,
		&store.Slug,
		&store.Description,
		&store.VerticalType,
		&store.IsActive,
		&store.IsPublished,
		&store.CreatedAt,
		&store.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &store, nil
}

func GetStoresBySellerID(sellerID string) ([]models.Store, error) {
	rows, err := db.DB.Query(`
		SELECT id, seller_id, name, slug, description, vertical_type,
		       is_active, is_published, created_at, updated_at
		FROM stores
		WHERE seller_id = $1
		AND deleted_at IS NULL
		ORDER BY created_at DESC
	`, sellerID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []models.Store
	for rows.Next() {
		var store models.Store
		err := rows.Scan(
			&store.ID,
			&store.SellerID,
			&store.Name,
			&store.Slug,
			&store.Description,
			&store.VerticalType,
			&store.IsActive,
			&store.IsPublished,
			&store.CreatedAt,
			&store.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	return stores, nil
}

func GetStoreByID(storeID, sellerID string) (*models.Store, error) {
	var store models.Store

	err := db.DB.QueryRow(`
		SELECT id, seller_id, name, slug, description, vertical_type,
		       is_active, is_published, created_at, updated_at
		FROM stores
		WHERE id = $1
		AND seller_id = $2
		AND deleted_at IS NULL
	`, storeID, sellerID).Scan(
		&store.ID,
		&store.SellerID,
		&store.Name,
		&store.Slug,
		&store.Description,
		&store.VerticalType,
		&store.IsActive,
		&store.IsPublished,
		&store.CreatedAt,
		&store.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("store not found")
	}

	if err != nil {
		return nil, err
	}

	return &store, nil
}

func DeleteStore(storeID, sellerID string) error {
	result, err := db.DB.Exec(`
		UPDATE stores
		SET deleted_at = now()
		WHERE id = $1
		AND seller_id = $2
		AND deleted_at IS NULL
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
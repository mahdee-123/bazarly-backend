package product

import (
	"database/sql"
	"errors"

	"github.com/mahdee-123/bazarly-backend/internal/db"
)

func Create(storeID string, req CreateProductRequest) (*Product, error) {
	var p Product

	attributes := req.Attributes
	if attributes == "" {
		attributes = "{}"
	}

		images := req.Images
		if images == "" {
			images = "[]"
		}

		err := db.DB.QueryRow(`
			INSERT INTO products (store_id, name, description, base_price, pricing_strategy, attributes, images)
			VALUES ($1, $2, $3, $4, $5, $6::jsonb, $7::jsonb)
			RETURNING id, store_id, name, description, base_price, pricing_strategy,
						attributes::text, images::text, is_active, created_at, updated_at
		`, storeID, req.Name, req.Description, req.BasePrice, req.PricingStrategy, attributes, images).Scan(
			&p.ID, &p.StoreID, &p.Name, &p.Description,
			&p.BasePrice, &p.PricingStrategy, &p.Attributes,
			&p.Images, &p.IsActive, &p.CreatedAt, &p.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		return &p, nil
	}

	func GetByStoreID(storeID string) ([]Product, error) {
		rows, err := db.DB.Query(`
			SELECT id, store_id, name, description, base_price, pricing_strategy,
					attributes::text, images::text, is_active, created_at, updated_at
			FROM products
			WHERE store_id = $1
			AND deleted_at IS NULL
			ORDER BY created_at DESC
		`, storeID)

		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var products []Product
		for rows.Next() {
			var p Product
			err := rows.Scan(
				&p.ID, &p.StoreID, &p.Name, &p.Description,
				&p.BasePrice, &p.PricingStrategy, &p.Attributes,
				&p.Images, &p.IsActive, &p.CreatedAt, &p.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}
			products = append(products, p)
		}

		return products, nil
	}

func GetByID(productID, storeID string) (*Product, error) {
	var p Product

	err := db.DB.QueryRow(`
		SELECT id, store_id, name, description, base_price, pricing_strategy,
		       attributes::text, images::text, is_active, created_at, updated_at
		FROM products
		WHERE id = $1
		AND store_id = $2
		AND deleted_at IS NULL
	`, productID, storeID).Scan(
		&p.ID, &p.StoreID, &p.Name, &p.Description,
		&p.BasePrice, &p.PricingStrategy, &p.Attributes,
		&p.Images, &p.IsActive, &p.CreatedAt, &p.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func Update(productID, storeID string, req UpdateProductRequest) (*Product, error) {
	var p Product

	err := db.DB.QueryRow(`
		UPDATE products
		SET name = COALESCE(NULLIF($1, ''), name),
		    description = COALESCE(NULLIF($2, ''), description),
		    base_price = CASE WHEN $3 > 0 THEN $3 ELSE base_price END,
		    is_active = $4,
		    updated_at = now()
		WHERE id = $5
		AND store_id = $6
		AND deleted_at IS NULL
		RETURNING id, store_id, name, description, base_price, pricing_strategy,
		          attributes::text, images::text, is_active, created_at, updated_at
	`, req.Name, req.Description, req.BasePrice, req.IsActive, productID, storeID).Scan(
		&p.ID, &p.StoreID, &p.Name, &p.Description,
		&p.BasePrice, &p.PricingStrategy, &p.Attributes,
		&p.Images, &p.IsActive, &p.CreatedAt, &p.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func Delete(productID, storeID string) error {
	result, err := db.DB.Exec(`
		UPDATE products
		SET deleted_at = now()
		WHERE id = $1
		AND store_id = $2
		AND deleted_at IS NULL
	`, productID, storeID)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}
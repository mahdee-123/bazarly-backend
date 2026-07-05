package services

import (
	"errors"
	"regexp"
	"strings"

	"github.com/mahdee-123/bazarly-backend/models"
	"github.com/mahdee-123/bazarly-backend/repository"
)

func CreateStore(sellerID string, req models.CreateStoreRequest) (*models.Store, error) {
	// Slug validate — only lowercase letters, numbers, hyphens
	slug := strings.ToLower(req.Slug)
	matched, _ := regexp.MatchString(`^[a-z0-9-]+$`, slug)
	if !matched {
		return nil, errors.New("slug can only contain lowercase letters, numbers, and hyphens")
	}

	req.Slug = slug

	store, err := repository.CreateStore(sellerID, req)
	if err != nil {
		return nil, errors.New("slug already taken or invalid data")
	}

	return store, nil
}

func GetMyStores(sellerID string) ([]models.Store, error) {
	stores, err := repository.GetStoresBySellerID(sellerID)
	if err != nil {
		return nil, errors.New("failed to fetch stores")
	}
	return stores, nil
}

func GetStore(storeID, sellerID string) (*models.Store, error) {
	store, err := repository.GetStoreByID(storeID, sellerID)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func DeleteStore(storeID, sellerID string) error {
	return repository.DeleteStore(storeID, sellerID)
}
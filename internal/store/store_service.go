package store

import (
	"errors"
	"regexp"
	"strings"
)

func CreateStore(sellerID string, req CreateStoreRequest) (*Store, error) {
	slug := strings.ToLower(req.Slug)
	matched, _ := regexp.MatchString(`^[a-z0-9-]+$`, slug)
	if !matched {
		return nil, errors.New("slug can only contain lowercase letters, numbers, and hyphens")
	}

	req.Slug = slug

	newStore, err := CreateStoreRepo(sellerID, req)
	if err != nil {
		return nil, errors.New("slug already taken or invalid data")
	}

	return newStore, nil
}

func GetMyStores(sellerID string) ([]Store, error) {
	stores, err := GetStoresBySellerID(sellerID)
	if err != nil {
		return nil, errors.New("failed to fetch stores")
	}
	return stores, nil
}

func GetStore(storeID, sellerID string) (*Store, error) {
	return GetStoreByID(storeID, sellerID)
}

func DeleteStore(storeID, sellerID string) error {
	return DeleteStoreRepo(storeID, sellerID)
}
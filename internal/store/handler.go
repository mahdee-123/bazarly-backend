package store

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdee-123/bazarly-backend/internal/utils"
)

func CreateStoreHandler(c *gin.Context) {
	sellerID := c.GetString("seller_id")

	var req CreateStoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	newStore, err := CreateStore(sellerID, req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusCreated, "Store created successfully", newStore)
}

func GetMyStoresHandler(c *gin.Context) {
	sellerID := c.GetString("seller_id")

	stores, err := GetMyStores(sellerID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Stores fetched", stores)
}

func GetStoreHandler(c *gin.Context) {
	sellerID := c.GetString("seller_id")
	storeID := c.Param("storeId") // "id" → "storeId"

	foundStore, err := GetStore(storeID, sellerID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Store fetched", foundStore)
}

func DeleteStoreHandler(c *gin.Context) {
	sellerID := c.GetString("seller_id")
	storeID := c.Param("storeId") // "id" → "storeId"

	if err := DeleteStore(storeID, sellerID); err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Store deleted", nil)
}
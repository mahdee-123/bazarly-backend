package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdee-123/bazarly-backend/models"
	"github.com/mahdee-123/bazarly-backend/services"
	"github.com/mahdee-123/bazarly-backend/utils"
)

func CreateStore(c *gin.Context) {
	sellerID := c.GetString("seller_id")

	var req models.CreateStoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	store, err := services.CreateStore(sellerID, req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusCreated, "Store created successfully", store)
}

func GetMyStores(c *gin.Context) {
	sellerID := c.GetString("seller_id")

	stores, err := services.GetMyStores(sellerID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Stores fetched", stores)
}

func GetStore(c *gin.Context) {
	sellerID := c.GetString("seller_id")
	storeID := c.Param("id")

	store, err := services.GetStore(storeID, sellerID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Store fetched", store)
}

func DeleteStore(c *gin.Context) {
	sellerID := c.GetString("seller_id")
	storeID := c.Param("id")

	err := services.DeleteStore(storeID, sellerID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Store deleted", nil)
}
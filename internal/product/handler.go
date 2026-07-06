package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdee-123/bazarly-backend/internal/store"
	"github.com/mahdee-123/bazarly-backend/internal/utils"
)

func CreateHandler(c *gin.Context) {
	storeID := c.Param("storeId")
	sellerID := c.GetString("seller_id")

	// Store ownership verify
	_, err := store.GetStoreByID(storeID, sellerID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Store not found")
		return
	}

	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	p, err := Create(storeID, req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusCreated, "Product created successfully", p)
}

func ListHandler(c *gin.Context) {
	storeID := c.Param("storeId")
	sellerID := c.GetString("seller_id")

	_, err := store.GetStoreByID(storeID, sellerID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Store not found")
		return
	}

	products, err := GetByStoreID(storeID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Products fetched", products)
}

func GetHandler(c *gin.Context) {
	storeID := c.Param("storeId")
	productID := c.Param("productId")
	sellerID := c.GetString("seller_id")

	_, err := store.GetStoreByID(storeID, sellerID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Store not found")
		return
	}

	p, err := GetByID(productID, storeID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Product fetched", p)
}

func UpdateHandler(c *gin.Context) {
	storeID := c.Param("storeId")
	productID := c.Param("productId")
	sellerID := c.GetString("seller_id")

	_, err := store.GetStoreByID(storeID, sellerID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Store not found")
		return
	}

	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	p, err := Update(productID, storeID, req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Product updated", p)
}

func DeleteHandler(c *gin.Context) {
	storeID := c.Param("storeId")
	productID := c.Param("productId")
	sellerID := c.GetString("seller_id")

	_, err := store.GetStoreByID(storeID, sellerID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Store not found")
		return
	}

	err = Delete(productID, storeID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Product deleted", nil)
}
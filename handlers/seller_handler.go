package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdee-123/bazarly-backend/models"
	"github.com/mahdee-123/bazarly-backend/services"
	"github.com/mahdee-123/bazarly-backend/utils"
)

func SellerSignup(c *gin.Context) {
	var req models.SellerSignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	seller, err := services.SignupSeller(req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusCreated, "Signup successful", seller)
}

func SellerLogin(c *gin.Context) {
	var req models.SellerLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	seller, token, err := services.LoginSeller(req)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Login successful", gin.H{
		"seller": seller,
		"token":  token,
	})
}

func GetSellerProfile(c *gin.Context) {
	sellerID := c.GetString("seller_id")

	seller, err := services.GetSellerProfile(sellerID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Profile fetched", seller)
}
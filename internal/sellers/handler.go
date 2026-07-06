package sellers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdee-123/bazarly-backend/internal/utils"
)

func SellerSignupHandler(c *gin.Context) {
	var req SellerSignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	newSeller, err := SignupSeller(req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusCreated, "Signup successful", newSeller)
}

func SellerLoginHandler(c *gin.Context) {
	var req SellerLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	loggedInSeller, token, err := LoginSeller(req)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Login successful", gin.H{
		"seller": loggedInSeller,
		"token":  token,
	})
}

func GetSellerProfileHandler(c *gin.Context) {
	sellerID := c.GetString("seller_id")

	profile, err := GetSellerProfile(sellerID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Profile fetched", profile)
}

func VerifyEmailHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		utils.Error(c, http.StatusBadRequest, "Token is required")
		return
	}

	if err := VerifySellerEmail(token); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	

	utils.Success(c, http.StatusOK, "Email verified successfully", nil)
}
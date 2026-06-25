package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"notes_api/internal/dto"
	"notes_api/internal/service"
	"notes_api/internal/utils"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	user, err := h.authService.Register(req)

	if err != nil {

		if err.Error() == "email already exists" {
			utils.BadRequest(c, err.Error())
			return
		}

		utils.InternalServerError(c)
		return
	}

	utils.Success(
		c,
		http.StatusCreated,
		"User registered successfully",
		user,
	)
}

func (h *AuthHandler) Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	result, err := h.authService.Login(req)

	if err != nil {

		if err.Error() == "invalid email or password" {
			utils.Unauthorized(c, err.Error())
			return
		}

		utils.InternalServerError(c)
		return
	}

	utils.Success(
		c,
		http.StatusOK,
		"Login successful",
		result,
	)
}

func (h *AuthHandler) Logout(c *gin.Context) {

	utils.Success(
		c,
		http.StatusOK,
		"Logout successful",
		nil,
	)

}

package handler

import (
	"net/http"
	"workshop-be/auth"
	"workshop-be/helper"
	"workshop-be/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService}
}

func (h *userHandler) Register(c *gin.Context) {
	var input user.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ApiResponse("Harap isi semua input", http.StatusUnprocessableEntity, "error", helper.FormatBindError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Gagal menambahkan user", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := h.authService.GenerateToken(int(newUser.ID))
	if err != nil {
		response := helper.ApiResponse("Gagal membuat token", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Akun berhasil dibuat", http.StatusOK, "success", user.FormatterUser(newUser, token))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ApiResponse("Harap isi semua input", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var userLoggedin user.User
	userLoggedin, err := h.userService.LoginUser(input)
	if err != nil {
		errorRes := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login gagal", http.StatusUnprocessableEntity, "gagal", errorRes)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(int(userLoggedin.ID))
	if err != nil {
		response := helper.ApiResponse("Terjadi kesalahan dengan token", http.StatusInternalServerError, "gagal", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Berhasil login", http.StatusOK, "sukses", user.FormatterUser(userLoggedin, token))
	c.JSON(http.StatusOK, response)

}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadzhuhry/bwastartup/helper"
	"github.com/muhammadzhuhry/bwastartup/user"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	// passing input ke service
	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}

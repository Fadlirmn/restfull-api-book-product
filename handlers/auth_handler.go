package handlers

import (
	"go-roadmap/models"
	"go-roadmap/services"
	
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{
	service *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler  {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	
	var user models.User

	if err:= c.ShouldBindJSON(&user) ;err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	err := h.service.Register(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"message":"user registred",
	})
}

func (h *AuthHandler) Login(c *gin.Context)  {
	var req struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.ShouldBindJSON(&req)

	token,err := h.service.Login(req.Username,req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":"invalid login",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"token":token,
	})
}

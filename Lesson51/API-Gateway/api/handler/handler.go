package handler

// import (
// 	"api-service/api/token"
// 	"api-service/models"
// 	"api-service/service"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Handler struct {
// 	User service.UserService
// }

// func (h *Handler) Register(c *gin.Context) {
// 	var req models.RegisterReq

// 	if err := c.BindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	exists, err := h.User.EmailExists(req.Email)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 	}

// 	if exists {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "email alreafy registred",
// 		})
// 	}

// 	req.Password = "pass"

// 	tokens := token.GeneratedJWTToken(req.ID, req.Email, req.Username)
// 	c.JSON(http.StatusCreated, tokens)
// }

import (
	"auth-service/service"
)

type HTTPHandler struct {
	US     *service.UserService
	Logger logger.Logger
}

func NewHandler(us *service.UserService, l logger.Logger) *HTTPHandler {
	return &HTTPHandler{US: us, Logger: l}
}
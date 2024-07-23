package handler

import (
	"gin/api/models"
	"gin/api/token"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var req models.AuthReq

	err := c.BindJSON(&req)
	if err != nil {
		log.Println("error while binding json : ", err)
	}

	err = h.auth.Register(&req)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "error while registering",
		})
		log.Println("error while registering : ", err)
	}

	c.IndentedJSON(201, gin.H{"message": "succesfully created, you can login !!!"})

}

func (h *Handler) Login(c *gin.Context) {
	var req models.LoginReq

	err := c.BindJSON(&req)
	if err != nil {
		log.Println("error while binding json : ", err)
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.auth.Login(&req)
	if err != nil {
		if err.Error() == "user_name or password incorrect" {
			c.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(500, gin.H{
			"error": "user not found",
		})
		log.Println("error while checking data on login : ", err)
		return
	}

	token := token.GenToken(&models.TokenReq{
		UserId:   res.Id,
		UserName: res.UserName,
		Email:    res.Email,
	})

	c.IndentedJSON(200, models.LoginRes{
		User:   *res,
		Tokens: *token,
	})

}

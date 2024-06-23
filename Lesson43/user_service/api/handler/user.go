package handler

import (
	"fmt"
	"net/http"
	"user_service/models"

	"github.com/gin-gonic/gin"
)

func (u *handler) CreateUser(ctx *gin.Context) {
	user := models.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = u.User.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "Created")
}

func (u *handler) GetUserById(ctx *gin.Context) {
	UserId := ctx.Param("id")

	user, err := u.User.GetById(UserId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", UserId)})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *handler) GetAllUser(ctx *gin.Context) {
	users, err := u.User.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (u *handler) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUser, err := u.User.Update(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updateUser)
}

func (u *handler) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := u.User.Delete(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete user: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "User deleted succesfully"})
}

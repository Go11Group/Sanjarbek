package handler

import (
	s "gin/storage/postgres"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

type Handler struct {
	storage *s.UserRepo
	auth *s.NewAuth

}

func NewHandler(s *s.UserRepo) *Handler {
	return &Handler{
		storage: s,
	}
}

type UserReq struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func (h *Handler) CreateUser(ctx *gin.Context) {

	newUser := UserReq{}

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		log.Println("error while binding body : ", err)
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	pp.Println(newUser)

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Successfully created",
	})

}

func (h *Handler) GetUser(ctx *gin.Context) {

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Successfull",
	})

}

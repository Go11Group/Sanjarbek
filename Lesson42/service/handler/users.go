package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(u *gin.Context) {
	url := "http://localhost:8080" + u.Request.URL.String()

	resp, err := http.Get(url)
	if err != nil {
		u.JSON(http.StatusBadRequest, gin.H{"Message": err})
		log.Println("Bad request ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		u.JSON(http.StatusBadRequest, gin.H{"Message": err})
		log.Println("Bad request ", err)
		return
	}
	u.JSON(http.StatusOK, string(body))
}

func (h *Handler) CreateUsers(u *gin.Context) {
	url := "http://localhost:8080" + u.Request.URL.String()
	req, err := http.NewRequest("POST", url, u.Request.Body)
	if err != nil {
		u.JSON(http.StatusBadRequest, gin.H{"Message": err})
		log.Println("Error while making rerquest ", err)
		return
	}
	resp, err := h.Client.Do(req)
	if err != nil {
		u.JSON(http.StatusBadRequest, gin.H{"Message": err})
		log.Println("Error while getting rerquest ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		u.JSON(http.StatusBadRequest, gin.H{"Message": err})
		log.Println("Bad request ", err)
		return
	}
	u.JSON(http.StatusOK, string(body))
}

func (h *Handler) UpdateUsers(u *gin.Context) {

	url := "http://localhost:8080" + u.Request.URL.String()
	req, err := http.NewRequest("PUT", url, u.Request.Body)
	if err != nil {
		u.JSON(http.StatusBadRequest, gin.H{"Message": err})
		log.Println("Error while rerquest ", err)
		return
	}
	resp, err := h.Client.Do(req)
	if err != nil {
		u.JSON(http.StatusBadRequest, gin.H{"Message": err})
		log.Println("Error while rerquest ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		u.JSON(http.StatusBadRequest, gin.H{"Message": err})
		log.Println("Bad request ", err)
		return
	}
	u.JSON(http.StatusOK, string(body))
}

func (h *Handler) DeleteUser(u *gin.Context) {

	url := "http://localhost:8080" + u.Request.URL.String()
	req, err := http.NewRequest("DELETE", url, u.Request.Body)
	if err != nil {
		u.JSON(http.StatusBadRequest, gin.H{"Message": err})
		log.Println("Error while making rerquest ", err)
		return
	}
	resp, err := h.Client.Do(req)
	if err != nil {
		u.JSON(http.StatusBadRequest, gin.H{"Message": err})
		log.Println("Error while getting rerquest ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		u.JSON(http.StatusBadRequest, gin.H{"Message": err})
		log.Println("Bad request ", err)
		return
	}
	u.JSON(http.StatusOK, string(body))
}
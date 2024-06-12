package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func Handler(h HandlerRepo) *http.Server{

	gin := gin.Default()
	user := gin.Group("/user")

	user.POST("", h.UserCreate)
	user.GET("/:id", h.UserGetByID)

	return &http.Server{Addr: ":8080", Handler: gin}
}
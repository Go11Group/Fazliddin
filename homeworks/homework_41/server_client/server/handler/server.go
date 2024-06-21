package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server(handler HandlerRepo) *http.Server {
	gin := gin.Default()

	user := gin.Group("/user")

	user.POST("", handler.UserCreate)
	user.GET("/:id", handler.UserGetByID)
	user.GET("", handler.UserGet)
	user.PUT("/:id", handler.UserUpdate)
	user.DELETE("/:id", handler.UserDelete)

	return &http.Server{Addr: ":8080", Handler: gin}
}

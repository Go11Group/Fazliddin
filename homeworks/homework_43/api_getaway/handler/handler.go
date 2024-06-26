package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler() *http.Server {

	gin := gin.Default()

	station := gin.Group("/station")
	terminal := gin.Group("/terminal")
	user := gin.Group("/user")
	card := gin.Group("/card")
	transaction := gin.Group("/transaction")

	station.POST("", Create)
	station.GET("", Get)
	station.GET("/:id", GetByID)
	station.PUT("/:id", Update)
	station.DELETE("/:id", Delete)

	terminal.POST("", Create)
	terminal.GET("", Get)
	terminal.GET("/:id", GetByID)
	terminal.PUT("/:id", Update)
	terminal.DELETE("/:id", Delete)

	user.POST("", Create)
	user.GET("", Get)
	user.GET("/:id", GetByID)
	user.PUT("/:id", Update)
	user.DELETE("/:id", Delete)

	card.POST("", Create)
	card.GET("", Get)
	card.GET("/:id", GetByID)
	card.PUT("/:id", Update)
	card.DELETE("/:id", Delete)

	transaction.POST("", Create)
	transaction.GET("", Get)
	transaction.GET("/:id", GetByID)
	transaction.PUT("/:id", Update)
	transaction.DELETE("/:id", Delete)

	return &http.Server{Addr: ":8081", Handler: gin}
}

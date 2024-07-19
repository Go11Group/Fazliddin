package api

import (
	"github.com/gin-gonic/gin"
	"homework_62/api/handler"
)

func Router(hd handler.Handler) {
	router := gin.Default()

	router.POST("/book")
	router.GET("/book/:name")
	router.PUT("/book/:name")
	router.DELETE("/book/:name")
	router.GET("/book")
}

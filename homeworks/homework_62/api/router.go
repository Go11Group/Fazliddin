package api

import (
	"github.com/gin-gonic/gin"
	"homework_62/api/handler"
)

func Router(hd handler.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/book", hd.CreateBook)
	router.GET("/book/:name", hd.GetBook)
	router.DELETE("/book/:name", hd.DeleteBook)
	router.GET("/book", hd.GetBooks)

	return router
}

package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"homework_62/api/handler"
	"homework_62/api/middleware"
)

func Router(hd handler.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/users", hd.CreataUser)

	router.Use(middleware.CheckPermissionMiddleware(&casbin.Enforcer{}))

	router.POST("/book", hd.CreateBook)
	router.GET("/book/:name", hd.GetBook)
	router.DELETE("/book/:name", hd.DeleteBook)
	router.GET("/book", hd.GetBooks)

	return router
}

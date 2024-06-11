package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Handler(handler HandlerRepo) *http.Server{
	gin := gin.Default()

	user := gin.Group("/user")

	user.POST("", handler.UserCreate)
	user.GET("/:id", handler.UserGetByID)
	user.GET("", handler.UserGet)
	user.PUT("", handler.UserUpdate)
	user.DELETE("/:id", handler.UserDelete)

	problems := gin.Group("/problems")
	problems.POST("", handler.ProblemCreate)
	problems.GET("/:id}", handler.ProblemGetByID)
	problems.GET("", handler.ProblemGet)
	problems.PUT("", handler.ProblemUpdate)
	problems.DELETE("/:id}", handler.ProblemDelete)

	

	return &http.Server{Addr: ":8080", Handler: gin}
}
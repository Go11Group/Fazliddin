package handler

import (
	"fmt"
	"model/user"
	"github.com/gin-gonic/gin"
)

func (h *HandlerRepo) UserCreate(c *gin.Context){
	user := user.User{}
	c.JSON(200, user)
	err := h.User.Create(user)
	c.Sho
} 

func (h *HandlerRepo) UserGetByID(c *gin.Context){
	id := c.Param("id")
	h.User.Create()
}
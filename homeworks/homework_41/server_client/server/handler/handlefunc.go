package handler

import (
	"fmt"
	user "homework_41/server_client/server/modul"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlerRepo) UserCreate(c *gin.Context) {
	u := user.User{}
	err := c.ShouldBind(&u)
	if err != nil {
		fmt.Println("ERROR JSON : ", err)
	}
	err = h.User.Create(u)
	if err != nil {
		fmt.Println("ERROR JSON: ", err)
		return
	}
	c.JSON(200, "user created")
}

func (h *HandlerRepo) UserGetByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.User.GetByID(id)
	if err != nil {
		fmt.Println("ERROR IN GET_ID : ", err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *HandlerRepo) UserGet(c *gin.Context) {
	users, err := h.User.Get()
	if err != nil {
		fmt.Println("ERROR IN GET : ", err)
	}
	c.JSON(http.StatusOK, users)
}

func (h *HandlerRepo) UserUpdate(c *gin.Context) {
	id := c.Param("id")
	u := user.User{}
	err := c.ShouldBindJSON(&u)
	fmt.Println(u)
	fmt.Println(u)
	fmt.Println(u)
	if err != nil {
		fmt.Println("ERROR JSON : ", err)
	}
	err = h.User.Update(u, id)
	if err != nil {
		fmt.Println("DATABASE ERROR : ", err)
		return
	}
}

func (h *HandlerRepo) UserDelete(c *gin.Context) {
	id := c.Param("id")

	err := h.User.Delete(id)
	if err != nil {
		fmt.Println("ERROR DELETE : ", err)
		return
	}
	c.JSON(200, "User Deleted")
}

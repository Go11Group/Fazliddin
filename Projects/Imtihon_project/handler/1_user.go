package handler

import (
	"fmt"
	"model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerRepo struct for create user
func (h *HandlerRepo) UserCreate(c *gin.Context) {
	user := user.User{}
	err := c.ShouldBindJSON(&user) //for read from body
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(200, gin.H{"ERROR IN READ USERS": err})
		return
	}
	err = h.User.Create(user) //insert to db
	if err != nil {
		c.JSON(200, gin.H{"ERROR IN CREATE USERS": err})
		fmt.Println(err)
		return
	}
}

// HandlerRepo struct for Get users with filter
func (h *HandlerRepo) UserGet(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from users where deleted_at=0`
	// filter
	if len(c.Query("name")) > 0 {
		params["name"] = c.Query("name")
		query += " and name = :name"
	}

	if len(c.Query("email")) > 0 {
		params["email"] = c.Query("email")
		query += " and email = :emailf "
	}

	if len(c.Query("birthday")) > 0 {
		params["birthday"] = c.Query("birthday")
		query += " and birthday = :birthday "
	}

	if len(c.Query("limit")) > 0 {
		params["limit"] = c.Query("limit")
		limit = ` LIMIT :limit`
	}

	if len(c.Query("offset")) > 0 {
		params["offset"] = c.Query("offset")
		offset = ` OFFSET :offset`
	}

	query = query + limit + offset

	query, arr = ReplaceQueryParams(query, params)

	users, err := h.User.Get(query, arr) // get with filter from db
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"ERROR IN GET": err})
		return
	}
	c.JSON(200, users)
}

// HandlerRepo struct for Get By ID
func (h *HandlerRepo) UserGetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.User.GetById(id) // get by id from db
	if err != nil {
		fmt.Println("Error get by ID: ", err)
		c.JSON(200, gin.H{"ERROR IN Get By ID USERS": err})
		return
	}
	c.JSON(200, user)
}

// HandlerRepo struct for Update
func (h *HandlerRepo) UserUpdate(c *gin.Context) {
	user := user.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN READ USERS": err})
		return
	}
	id := c.Param("id")
	err = h.User.Update(id, user) //update user from db
	if err != nil {
		fmt.Println("ERROR In Update Users: ", err)
		c.JSON(200, gin.H{"ERROR IN UPDATE USERS": err})
		return
	}
	c.JSON(200, gin.H{"User": "Successfully Updated"})
}

// HandlerRepo struct for soft delete
func (h *HandlerRepo) UserDelete(c *gin.Context) {
	id := c.Param("id")
	err := h.User.Delete(id) //soft delete from db
	if err != nil {
		fmt.Println("ERROR In Delete Users: ", err)
		c.JSON(200, gin.H{"ERROR in delete": err})
		return
	}
	c.JSON(200, gin.H{"User": "Successfully deleted"})
}

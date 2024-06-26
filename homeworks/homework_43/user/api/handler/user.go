package handler

import (
	"fmt"
	"homeworks/homework_43/user/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("Scan user ", err)
		return
	}

	err = h.User.Create(user)
	if err != nil {
		fmt.Println("Create user ", err)
	}
	c.JSON(http.StatusOK, gin.H{"succesfull": "user created"})
}

func (h *Handler) UserGet(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from users where deleted_at=0`
	// filter
	if len(c.Query("age")) > 0 {
		age, err := strconv.Atoi(c.Query("age"))
		if err != nil {
			fmt.Println("Atoi err ", err)
		}
		params["age"] = age
		query += " and age = :age"
	}

	if len(c.Query("name")) > 0 {
		params["name"] = c.Query("name")
		query += " and name = :name"
	}

	if len(c.Query("phone")) > 0 {
		params["phone"] = c.Query("phone")
		query += " and phone = :phone "
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
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET USERS": err})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) UserGetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.User.GetById(id) // get by id from db
	if err != nil {
		fmt.Println("Error get by ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET BY ID USERS": err})
		return
	}
	c.JSON(200, user)
}

// Handler struct for Update
func (h *Handler) UserUpdate(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN READ FROM BODY USERS": err})
		return
	}
	id := c.Param("id")
	err = h.User.Update(user, id) //update user from db
	if err != nil {
		fmt.Println("ERROR In Update Users: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN UPDATE USERS": err})
		return
	}
	c.JSON(200, gin.H{"User": "Successfully Updated"})
}

// Handler struct for soft delete
func (h *Handler) UserDelete(c *gin.Context) {
	id := c.Param("id")
	err := h.User.Delete(id) //soft delete from db
	if err != nil {
		fmt.Println("ERROR In Delete Users: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN SOFT DELETE USERS": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfully": "User deleted"})
}

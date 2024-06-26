package handler

import (
	"fmt"
	"homeworks/homework_43/user/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CardCreate(c *gin.Context) {
	card := models.Card{}
	err := c.ShouldBindJSON(&card)
	if err != nil {
		fmt.Println("Scan card ", err)
		return
	}

	err = h.Card.Create(card)
	if err != nil {
		fmt.Println("Create card ", err)
	}
	c.JSON(http.StatusOK, gin.H{"succesfull": "card created"})
}

func (h *Handler) CardGet(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from card where deleted_at=0`
	// filter
	if len(c.Query("number")) > 0 {
		number, err := strconv.Atoi(c.Query("number"))
		if err != nil {
			fmt.Println("Atoi err ", err)
		}
		params["number"] = number
		query += " and number = :number"
	}

	if len(c.Query("user_id")) > 0 {
		params["user_id"] = c.Query("user_id")
		query += " and user_id = :user_id "
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
	fmt.Println(query, arr)
	fmt.Println(query, arr)
	fmt.Println(query, arr[0]	)


	users, err := h.Card.Get(query, arr) // get with filter from db
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET Card": err})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) CardGetByID(c *gin.Context) {
	id := c.Param("id")
	card, err := h.Card.GetById(id) // get by id from db
	if err != nil {
		fmt.Println("Error get by ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET BY ID Card": err})
		return
	}
	c.JSON(200, card)
}

// Handler struct for Update
func (h *Handler) CardUpdate(c *gin.Context) {
	card := models.Card{}
	err := c.ShouldBindJSON(&card)
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN READ FROM BODY Card": err})
		return
	}
	id := c.Param("id")
	err = h.Card.Update(card, id) //update card from db
	if err != nil {
		fmt.Println("ERROR In Update Users: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN UPDATE Card": err})
		return
	}
	c.JSON(200, gin.H{"Card": "Successfully Updated"})
}

// Handler struct for soft delete
func (h *Handler) CardDelete(c *gin.Context) {
	id := c.Param("id")
	err := h.Card.Delete(id) //soft delete from db
	if err != nil {
		fmt.Println("ERROR In Delete Users: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN SOFT DELETE Card": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfully": "Card deleted"})
}

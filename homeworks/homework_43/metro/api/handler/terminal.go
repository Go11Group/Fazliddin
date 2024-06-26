package handler

import (
	"fmt"
	"homeworks/homework_43/metro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlerRepo) TerminalCreate(c *gin.Context) {
	terminal := models.Terminal{}
	err := c.ShouldBindJSON(&terminal)
	fmt.Println(terminal)
	if err != nil {
		fmt.Println("Scan terminal ", err)
		return
	}

	err = h.Terminal.Create(terminal)
	if err != nil {
		fmt.Println("Create terminal ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"succesfull": "terminal created"})
}

func (h *HandlerRepo) TerminalGet(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from terminal where deleted_at=0`

	if len(c.Query("station_id")) > 0 {
		params["station_id"] = c.Query("station_id")
		query += " and station_id = :station_id "
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

	users, err := h.Terminal.Get(query, arr) // get with filter from db
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET Terminal": err})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *HandlerRepo) TerminalGetByID(c *gin.Context) {
	id := c.Param("id")
	terminal, err := h.Terminal.GetById(id) // get by id from db
	if err != nil {
		fmt.Println("Error get by ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET BY ID Terminal": err})
		return
	}
	c.JSON(200, terminal)
}

// HandlerRepo struct for Update
func (h *HandlerRepo) TerminalUpdate(c *gin.Context) {
	terminal := models.Terminal{}
	err := c.ShouldBindJSON(&terminal)
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN READ FROM BODY Terminal": err})
		return
	}
	id := c.Param("id")
	err = h.Terminal.Update(terminal, id) //update terminal from db
	if err != nil {
		fmt.Println("ERROR In Update Terminal: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN UPDATE Terminal": err})
		return
	}
	c.JSON(200, gin.H{"Terminal": "Successfully Updated"})
}

// HandlerRepo struct for soft delete
func (h *HandlerRepo) TerminalDelete(c *gin.Context) {
	id := c.Param("id")
	err := h.Terminal.Delete(id) //soft delete from db
	if err != nil {
		fmt.Println("ERROR In Delete Terminal: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN SOFT DELETE Terminal": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfully": "Terminal deleted"})
}

package handler

import (
	"fmt"
	"homeworks/homework_43/metro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlerRepo) StationCreate(c *gin.Context) {
	station := models.Station{}
	err := c.ShouldBindJSON(&station)
	if err != nil {
		fmt.Println("Scan station ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in scan"})
		return
	}

	err = h.Station.Create(station)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error create"})
		fmt.Println("Create station ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"succesfull": "station created"})
}

func (h *HandlerRepo) StationdGet(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from station where deleted_at=0`
	// filter

	if len(c.Query("name")) > 0 {
		params["name"] = c.Query("name")
		query += " and name = :name "
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

	users, err := h.Station.Get(query, arr) // get with filter from db
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET Station": err})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *HandlerRepo) StationGetByID(c *gin.Context) {
	id := c.Param("id")
	station, err := h.Station.GetById(id) // get by id from db
	if err != nil {
		fmt.Println("Error get by ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET BY ID Station": err})
		return
	}
	c.JSON(200, station)
}

// HandlerRepo struct for Update
func (h *HandlerRepo) StationUpdate(c *gin.Context) {
	station := models.Station{}
	err := c.ShouldBindJSON(&station)
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN READ FROM BODY Station": err})
		return
	}
	id := c.Param("id")
	err = h.Station.Update(station, id) //update station from db
	if err != nil {
		fmt.Println("ERROR In Update Station: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN UPDATE Station": err})
		return
	}
	c.JSON(200, gin.H{"Station": "Successfully Updated"})
}

// HandlerRepo struct for soft delete
func (h *HandlerRepo) StationDelete(c *gin.Context) {
	id := c.Param("id")
	err := h.Station.Delete(id) //soft delete from db
	if err != nil {
		fmt.Println("ERROR In Delete Station: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN SOFT DELETE Station": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfully": "Station deleted"})
}

package handler

import (
	"fmt"
	"model/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerRepo struct for create user
func (h *HandlerRepo) CourseCreate(c *gin.Context) {
	course := course.Course{}
	err := c.ShouldBindJSON(&course) //for read from body
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(200, gin.H{"ERROR IN READ COURSE": err})
		return
	}
	err = h.Course.Create(course) //insert to db
	if err != nil {
		c.JSON(200, gin.H{"ERROR IN CREATE COURSE": err})
		fmt.Println(err)
		return
	}
}

// HandlerRepo struct for Get users with filter
func (h *HandlerRepo) CourseGet(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from course where deleted_at=0`
	// filter
	if len(c.Query("title")) > 0 {
		params["title"] = c.Query("title")
		query += " and title = :title"
	}

	if len(c.Query("description")) > 0 {
		params["description"] = c.Query("description")
		query += " and description = :description "
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

	course, err := h.Course.Get(query, arr) // get with filter from db
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"ERROR IN GET": err})
		return
	}
	c.JSON(200, course)
}

// HandlerRepo struct for Get By ID
func (h *HandlerRepo) CourseGetByID(c *gin.Context) {
	id := c.Param("id")
	course, err := h.Course.GetById(id) // get by id from db
	if err != nil {
		fmt.Println("Error get by ID: ", err)
		c.JSON(200, gin.H{"ERROR IN Get By ID Course": err})
		return
	}
	c.JSON(200, course)
}

// HandlerRepo struct for Update
func (h *HandlerRepo) CourseUpdate(c *gin.Context) {
	course := course.Course{}
	err := c.ShouldBindJSON(&course)
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN READ Course": err})
		return
	}
	id := c.Param("id")
	err = h.Course.Update(id, course) //update user from db
	if err != nil {
		fmt.Println("ERROR In Update Course: ", err)
		c.JSON(200, gin.H{"ERROR IN UPDATE COURSE": err})
		return
	}
	c.JSON(200, gin.H{"User": "Successfully Updated"})
}

// HandlerRepo struct for soft delete
func (h *HandlerRepo) CourseDelete(c *gin.Context) {
	id := c.Param("id")
	err := h.Course.Delete(id) //soft delete from db
	if err != nil {
		fmt.Println("ERROR In Delete Users: ", err)
		c.JSON(200, gin.H{"ERROR in delete": err})
		return
	}
	c.JSON(200, gin.H{"User": "Successfully deleted"})
}

package handler

import (
	"fmt"
	"go/go11/Fazliddin/Projects/Imtihon_project/lesson"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerRepo struct for create user
func (h *HandlerRepo) LessonCreate(c *gin.Context) {
	lesson := lesson.Lesson{}
	err := c.ShouldBindJSON(&lesson) //for read from body
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(200, gin.H{"ERROR IN READ COURSE": err})
		return
	}
	err = h.Lesson.Create(lesson) //insert to db
	if err != nil {
		c.JSON(200, gin.H{"ERROR IN CREATE COURSE": err})
		fmt.Println(err)
		return
	}
}

// HandlerRepo struct for Get users with filter
func (h *HandlerRepo) LEssonGet(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from lessons where deleted_at=0`
	// filter
	if len(c.Query("course_id")) > 0 {
		params["course_id"] = c.Query("course_id")
		query += " and course_id = :course_id"
	}

	if len(c.Query("title")) > 0 {
		params["title"] = c.Query("title")
		query += " and title = :title "
	}

	if len(c.Query("content")) > 0 {
		params["content"] = c.Query("content")
		query += " and content = :content "
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

	course, err := h.Lesson.Get(query, arr) // get with filter from db
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET": err})
		return
	}
	c.JSON(200, course)
}

// HandlerRepo struct for Get By ID
func (h *HandlerRepo) LessonGetByID(c *gin.Context) {
	id := c.Param("id")
	course, err := h.Lesson.GetById(id) // get by id from db
	if err != nil {
		fmt.Println("Error get by ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN Get By ID Lesson": err})
		return
	}
	c.JSON(200, course)
}

// HandlerRepo struct for Update
func (h *HandlerRepo) LessonUpdate(c *gin.Context) {
	lesson := lesson.Lesson{}
	err := c.ShouldBindJSON(&lesson)
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN READ Lesson": err})
		return
	}
	id := c.Param("id")
	err = h.Lesson.Update(id, lesson) //update user from db
	if err != nil {
		fmt.Println("ERROR In Update Course: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN UPDATE Lesson": err})
		return
	}
	c.JSON(200, gin.H{"User": "Successfully Updated"})
}

// HandlerRepo struct for soft delete
func (h *HandlerRepo) LessonDelete(c *gin.Context) {
	id := c.Param("id")
	err := h.Lesson.Delete(id) //soft delete from db
	if err != nil {
		fmt.Println("ERROR In Delete Lesson: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR in delete": err})
		return
	}
	c.JSON(200, gin.H{"User": "Successfully deleted"})
}

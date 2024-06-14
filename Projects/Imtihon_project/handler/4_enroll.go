package handler

import (
	"fmt"
	"model/enrollment"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerRepo struct for create user
func (h *HandlerRepo) EnrollmentCreate(c *gin.Context) {
	enroll := enrollment.Enrollment{}
	err := c.ShouldBindJSON(&enroll) //for read from body
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(200, gin.H{"ERROR IN READ USERS": err})
		return
	}
	err = h.Enrollment.Create(enroll) //insert to db
	if err != nil {
		c.JSON(200, gin.H{"ERROR IN CREATE ENROLLMENT": err})
		fmt.Println(err)
		return
	}
}

// HandlerRepo struct for Get users with filter
func (h *HandlerRepo) EnrollmentGet(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from enrollment where deleted_at=0`
	// filter
	if len(c.Query("user_id")) > 0 {
		params["user_id"] = c.Query("user_id")
		query += " and user_id = :user_id"
	}

	if len(c.Query("course_id")) > 0 {
		params["course_id"] = c.Query("course_id")
		query += " and course_id = :course_id "
	}

	if len(c.Query("e_data")) > 0 {
		params["e_data"] = c.Query("e_data")
		query += " and enrollment_data = enrollment_data "
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

	users, err := h.Enrollment.Get(query, arr) // get with filter from db
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"ERROR IN GET": err})
		return
	}
	c.JSON(200, users)
}

// HandlerRepo struct for Get By ID
func (h *HandlerRepo) EnrollmentGetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Enrollment.GetById(id) // get by id from db
	if err != nil {
		fmt.Println("Error get by ID: ", err)
		c.JSON(200, gin.H{"ERROR IN Get By ID USERS": err})
		return
	}
	c.JSON(200, user)
}

// HandlerRepo struct for Update
func (h *HandlerRepo) EnrollmentUpdate(c *gin.Context) {
	enroll := enrollment.Enrollment{}
	err := c.ShouldBindJSON(&enroll)
	if err != nil {
		fmt.Println("Error in read from: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN READ USERS": err})
		return
	}
	id := c.Param("id")
	fmt.Println(enroll)
	fmt.Println(enroll)
	fmt.Println(enroll)
	err = h.Enrollment.Update(id, enroll) //update user from db
	if err != nil {
		fmt.Println("ERROR In Update Users: ", err)
		c.JSON(200, gin.H{"ERROR IN UPDATE Enrollment": err})
		return
	}
	c.JSON(200, gin.H{"Enrollment": "Successfully Updated"})
}

// HandlerRepo struct for soft delete
func (h *HandlerRepo) EnrollmentDelete(c *gin.Context) {
	id := c.Param("id")
	err := h.Enrollment.Delete(id) //soft delete from db
	if err != nil {
		fmt.Println("ERROR In Delete Users: ", err)
		c.JSON(200, gin.H{"ERROR in delete": err})
		return
	}
	c.JSON(200, gin.H{"Enrollment": "Successfully deleted"})
}

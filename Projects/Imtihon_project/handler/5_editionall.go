package handler

import (
	"fmt"
	"go/go11/Fazliddin/Projects/Imtihon_project/course"
	"go/go11/Fazliddin/Projects/Imtihon_project/lesson"
	"go/go11/Fazliddin/Projects/Imtihon_project/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	UserId  string          `json:"user_id"`
	Courses []course.Course `json:"courses`
}

type Lesson struct {
	UserId  string          `json:"user_id"`
	Lessons []lesson.Lesson `json:"courses`
}

type User struct {
	CourseId string      `json:"course_id"`
	User     []user.User `json:"users"`
}

// Get courses with user user ID
func (h *HandlerRepo) GetCoursesWithUserID(c *gin.Context) {
	id := c.Param("id")
	courses, err := h.Course.GetCoursesByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		fmt.Println("error in get : ", err)
		return
	}
	course := Course{id, courses}
	c.JSON(http.StatusOK, course)
}

// get lessons by course ID
func (h *HandlerRepo) GetLessonByCourseID(c *gin.Context) {
	id := c.Param("id")
	lessons, err := h.Lesson.GetLessonByCourse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		fmt.Println("error in get : ", err)
		return
	}
	lesson := Lesson{id, lessons}
	c.JSON(http.StatusOK, lesson)
}

// get users by course id
func (h *HandlerRepo) UserGetByCourseID(c *gin.Context) {
	id := c.Param("id")
	users, err := h.User.GetByCourseID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		fmt.Println("error in get : ", err)
		return
	}
	user := User{id, users}
	c.JSON(http.StatusOK, user)
}

// GET the most popular course
func (h *HandlerRepo) CoursePopular(c *gin.Context) {
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")
  fmt.Println(start_date, end_date)
	courses, err := h.Course.Popular(start_date, end_date)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"time_period": []string{start_date, end_date}, "popular_courses": courses})
}

// User search name and email
func (h *HandlerRepo) UserSearch(c *gin.Context){
  var (
		params = make(map[string]interface{})
		arr    []interface{}
    name = c.Query("name")
    email = c.Query("email")
    query = "select * from users where deleted_at = 0 "

	)

  if len(name) > 0 {
		params["name"] = c.Query("name")
		query += " and name = :name"
	}

	if len(email) > 0 {
		params["email"] = c.Query("email")
		query += " and email = :emailf "
	}
  query, arr = ReplaceQueryParams(query, params)

	users, err := h.User.Get(query, arr) // get with filter from db
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR IN GET USERS": err})
		return
	}
	c.JSON(200, gin.H{"results": users})
}
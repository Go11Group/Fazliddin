package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func(h *HandlerRepo) GetCoursesWithUserID(c *gin.Context){
	id := c.Param("id")
	courses, err := h.Course.GetCoursesByUserID(id)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
	fmt.Println("error in get : ", err)
    return
  }

  c.JSON(http.StatusOK, gin.H{"user_id": id, "courses": courses})
}
package handler

import (
	"model/course"
	"model/enrollment"
	"model/lesson"
	"model/user"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type HandlerRepo struct {
	User       *user.UserRepo
	Course     *course.CourseRepo
	Lesson     *lesson.LessonRepo
	Enrollment *enrollment.EnrollmentRepo
}

func Handler(h HandlerRepo) *http.Server {

	gin := gin.Default()

	user := gin.Group("/user")
	course := gin.Group("/course")
	lesson := gin.Group("/lesson")
	enrollment := gin.Group("/enrollment")

	// for users
	user.POST("", h.UserCreate)
	user.GET("", h.UserGet)
	user.GET("/:id", h.UserGetByID)
	user.PUT("/:id", h.UserUpdate)
	user.DELETE("/:id", h.UserDelete)

	//for courses
	course.POST("", h.CourseCreate)
	course.GET("", h.CourseGet)
	course.GET("/:id", h.CourseGetByID)
	course.PUT("/:id", h.CourseUpdate)
	course.DELETE("/:id", h.CourseDelete)
	course.GET("/getCoursesByU/:id", h.GetCoursesWithUserID)

	//for lessons
	lesson.POST("", h.LessonCreate)
	lesson.GET("", h.LEssonGet)
	lesson.GET("/:id", h.LessonGetByID)
	lesson.PUT("/:id", h.LessonUpdate)
	lesson.DELETE("/:id", h.LessonDelete)

	//for lessons
	enrollment.POST("", h.EnrollmentCreate)
	enrollment.GET("", h.EnrollmentGet)
	enrollment.GET("/:id", h.EnrollmentGetByID)
	enrollment.PUT("/:id", h.EnrollmentUpdate)
	enrollment.DELETE("/:id", h.EnrollmentDelete)

	return &http.Server{Addr: ":8080", Handler: gin}
}

// funksiya dlya peredelki query
func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)
	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}
	return namedQuery, args
}

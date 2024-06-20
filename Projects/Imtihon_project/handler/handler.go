package handler

import (
	"go/go11/Fazliddin/Projects/Imtihon_project/course"
	"go/go11/Fazliddin/Projects/Imtihon_project/enrollment"
	"go/go11/Fazliddin/Projects/Imtihon_project/lesson"
	"go/go11/Fazliddin/Projects/Imtihon_project/user"
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

	user := gin.Group("/user")	// grouped for user
	course := gin.Group("/course")	//grouped for sourse
	lesson := gin.Group("/lesson")	// grouped from lesson
	enrollment := gin.Group("/enrollment")	// grouped for enrollment

	// for users
	user.POST("", h.UserCreate)
	user.GET("", h.UserGet)
	user.GET("/:id", h.UserGetByID)
	user.PUT("/:id", h.UserUpdate)
	user.DELETE("/:id", h.UserDelete)
	user.GET("/:id/courses", h.GetCoursesWithUserID) // editional
	user.GET("/search", h.UserSearch)

	//for courses
	course.POST("", h.CourseCreate)
	course.GET("", h.CourseGet)
	course.GET("/:id", h.CourseGetByID)
	course.PUT("/:id", h.CourseUpdate)
	course.DELETE("/:id", h.CourseDelete)
	course.GET("/:id/lesson", h.GetLessonByCourseID)    // editional
	course.GET("/:id/enrollments", h.UserGetByCourseID) // editional
	course.GET("/popular", h.CoursePopular)

	//for lessons
	lesson.POST("", h.LessonCreate)
	lesson.GET("", h.LEssonGet)
	lesson.GET("/:id", h.LessonGetByID)
	lesson.PUT("/:id", h.LessonUpdate)
	lesson.DELETE("/:id", h.LessonDelete)

	//for enrollment
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

package handler

import 	(
	"github.com/gin-gonic/gin"
	"net/http"
)


func Handler() *http.Server {

	gin := gin.Default()

	user := gin.Group("/user")	// grouped for user
	course := gin.Group("/course")	//grouped for sourse
	lesson := gin.Group("/lesson")	// grouped from lesson
	enrollment := gin.Group("/enrollment")	// grouped for enrollment

	// for users
	user.POST("", Create)
	user.GET("", Get)
	user.GET("/:id", GetByID)
	user.PUT("/:id", Update)
	// user.DELETE("/:id", UserDelete)
	// user.GET("/:id/courses", GetCoursesWithUserID) // editional
	// user.GET("/search", UserSearch)

	//for courses
	course.POST("", Create)
	course.GET("", Get)
	course.GET("/:id", GetByID)
	course.PUT("/:id", Update)
	// course.DELETE("/:id", CourseDelete)
	// course.GET("/:id/lesson", GetLessonByCourseID)    // editional
	// course.GET("/:id/enrollments", UserGetByCourseID) // editional
	// course.GET("/popular", CoursePopular)

	//for lessons
	lesson.POST("", Create)
	lesson.GET("", Get)
	lesson.GET("/:id", GetByID)
	lesson.PUT("/:id", Update)
	// lesson.DELETE("/:id", LessonDelete)

	// //for enrollment
	enrollment.POST("", Create)
	enrollment.GET("", Get)
	enrollment.GET("/:id", GetByID)
	enrollment.PUT("/:id", Update)
	// enrollment.DELETE("/:id", EnrollmentDelete)

	return &http.Server{Addr: ":8081", Handler: gin}
}
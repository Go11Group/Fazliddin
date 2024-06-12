package handler

import (
	"model/course"
	"model/enrollment"
	"model/lesson"
	"model/user"
)

type HandlerRepo struct{
	User *user.UserRepo
	Course *course.CourseRepo
	Lesson *lesson.LessonRepo
	Enrollment *enrollment.EnrollmentRepo
}
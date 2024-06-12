package main

import (
	"model/course"
	"model/enrollment"
	"model/handler"
	"model/lesson"
	"model/postgres"
	"model/user"
)

func main() {
	db, err := postgres.Conn()
	if err != nil{
		panic(err)
	}

	u := user.NewUserRepo(db)
	c := course.NewCourseRepo(db)
	l := lesson.NewLessonRepo(db)
	e := enrollment.NewEnrollmentRepo(db)

	h := handler.HandlerRepo{u, c, l, e}

	server := handler.Handler(h)
	err = server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}
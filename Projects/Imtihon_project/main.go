package main

import (
	"go/go11/Fazliddin/Projects/Imtihon_project/course"
	"go/go11/Fazliddin/Projects/Imtihon_project/enrollment"
	"go/go11/Fazliddin/Projects/Imtihon_project/handler"
	"go/go11/Fazliddin/Projects/Imtihon_project/lesson"
	"go/go11/Fazliddin/Projects/Imtihon_project/postgres"
	"go/go11/Fazliddin/Projects/Imtihon_project/user"
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
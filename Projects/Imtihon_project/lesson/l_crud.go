package lesson

import (
	"database/sql"
	"fmt"
	"time"
)

type LessonRepo struct {
	DB *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{DB: db}
}

// Insert into users
func (l *LessonRepo) Create(lesson Lesson) error {
	_, err := l.DB.Exec("insert into Lessons(course_id, title, content) values($1, $2, $3)",
		lesson.CourseID, lesson.Title, lesson.Content)
	return err
}

// ........
func (l *LessonRepo) Get(q string, arr []interface{}) ([]Lesson, error) {
	rows, err := l.DB.Query(q, arr...)
	if err != nil {
		return nil, err
	}

	lessons := []Lesson{}
	for rows.Next() {
		lesson := Lesson{}
		err = rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content,
			&lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

// Get user by ID
func (u *LessonRepo) GetById(id string) (Lesson, error) {
	lesson := Lesson{}
	err := u.DB.QueryRow("select * from lessons where lesson_id = $1 and deleted_at=0", id).
		Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content,
			&lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)

	return lesson, err
}

// Update lesson-> Course_id, Title, Content
func (u *LessonRepo) Update(id string, lesson Lesson) error {
	fmt.Printf("%+v\n\n", lesson)
	_, err := u.DB.Exec("update lessons set course_id=$1, title=$2, content=$3, update_at=$4 where lesson_id=$5 and deleted_at=0",
		lesson.CourseID, lesson.Title, lesson.Content, time.Now(), id)

	return err
}

// Delete by Id
func (l *LessonRepo) Delete(id string) error {
	fmt.Println(id)
	fmt.Println(id)
	fmt.Println(id)
	_, err := l.DB.Exec(`update lessons set deleted_at = date_part('epoch', current_timestamp)::INT
	where lesson_id = $1 and deleted_at = 0`, id)
	return err
}

// editional func ..................................................
//get all lessons by course ID
func (l *LessonRepo) GetLessonByCourse(id string) ([]Lesson, error){
	rows, err := l.DB.Query("select * from lessons where course_id = $1 and deleted_at = 0", id)
	if err != nil {
		return nil, err
	}

	lessons := []Lesson{}
	for rows.Next() {
		lesson := Lesson{}
		err = rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content,
			&lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}
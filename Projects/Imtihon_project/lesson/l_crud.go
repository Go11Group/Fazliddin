package lesson

import (
	"database/sql"
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
	_, err := l.DB.Exec("insert into Lessons(lesson_id, course_id, title, content) values($1, $2, $3, $4)",
		lesson.Lesson_id, lesson.Course_id, lesson.Title, lesson.Content)

	return err
}

// ........
func (l *LessonRepo) Get() ([]Lesson, error) {
	rows, err := l.DB.Query("select * from lessons")
	if err != nil {
		return nil, err
	}

	lessons := []Lesson{}
	for rows.Next() {
		lesson := Lesson{}
		err = rows.Scan(&lesson.Lesson_id, &lesson.Course_id, &lesson.Title, &lesson.Content,
			&lesson.Created_at, &lesson.Update_at, &lesson.Deleted_at)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

// Get user by ID
func (u *LessonRepo) GetById(id int) (Lesson, error) {
	lesson := Lesson{}

	err := u.DB.QueryRow("select * from lessons where id = $1", id).
		Scan(&lesson.Lesson_id, &lesson.Course_id, &lesson.Title, &lesson.Content,
			&lesson.Created_at, &lesson.Update_at, &lesson.Deleted_at)

	return lesson, err
}

// Update lesson-> Course_id, Title, Content
func (u *LessonRepo) Update(lesson Lesson) error {
	_, err := u.DB.Exec("update lessons set course_id=$1, title=$2, content=$3, update_at=$4 where id=$5",
		lesson.Course_id, lesson.Title, lesson.Content, time.Now(), lesson.Lesson_id)

	return err
}

// Delete by Id
func (l *LessonRepo) Delete(id int) error {
	_, err := l.DB.Exec("delete from lessons where id = $1", id)
	return err
}

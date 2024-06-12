package course

import (
	"database/sql"
	"time"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{DB: db}
}

// Insert into course
func (c *CourseRepo) Create(course Course) error {
	_, err := c.DB.Exec("insert into course(course_id, title, description) values($1, $2, $3)",
		course.Course_id, course.Title, course.Description)
	return err
}

// Get with filter
func (c *CourseRepo) Get() ([]Course, error) {
	rows, err := c.DB.Query("select * from course where deleted_at is null")
	if err != nil {
		return nil, err
	}

	courses := []Course{}
	for rows.Next() {
		course := Course{}
		err = rows.Scan(&course.Course_id, &course.Title, &course.Description, &course.Created_at,
			&course.Update_at, &course.Deleted_at)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

// Get user by ID
func (c *CourseRepo) GetById(id int) (Course, error) {
	course := Course{}
	err := c.DB.QueryRow("select * from course where id = $1", id).
		Scan(&course.Course_id, &course.Title, &course.Description, &course.Created_at,
			&course.Update_at, &course.Deleted_at)
	return course, err
}

// Update course-> Title, Description
func (c *CourseRepo) Update(course Course) error {
	_, err := c.DB.Exec("update course set title=$1, description=$2, update_at=$3 where id=$4",
		course.Title, course.Description, time.Now(), course.Course_id)
	return err
}

// Delete by Id
func (c *CourseRepo) Delete(id int) error {
	_, err := c.DB.Exec("update course set deleted_at=$1 where id = $2", time.Now(), id)
	return err
}

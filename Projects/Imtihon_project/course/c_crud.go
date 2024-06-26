package course

import (
	"database/sql"
	"errors"
	"fmt"
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
	_, err := c.DB.Exec("insert into course(title, description) values($1, $2)",
		course.Title, course.Description)
	return err
}

// Get with filter
func (c *CourseRepo) Get(q string, arr []interface{}) ([]Course, error) {
	rows, err := c.DB.Query(q, arr...)
	if err != nil {
		return nil, err
	}
	courses := []Course{}
	for rows.Next() {
		course := Course{}
		err = rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt,
			&course.UpdatedAt, &course.DeletedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

// Get user by ID
func (c *CourseRepo) GetById(id string) (Course, error) {
	course := Course{}
	err := c.DB.QueryRow("select * from course where course_id = $1 and deleted_at = 0", id).
		Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt,
			&course.UpdatedAt, &course.DeletedAt)
	return course, err
}

// Update course-> Title, Description
func (c *CourseRepo) Update(id string, course Course) error {
	_, err := c.DB.Exec("update course set title=$1, description=$2, update_at=$4 where course_id=$3 and deleted_at = 0",
		course.Title, course.Description, id, time.Now())
	return err
}

// Delete by Id
func (c *CourseRepo) Delete(id string) error {
	_, err := c.DB.Exec(`update course set deleted_at = date_part('epoch', current_timestamp)::INT
	where course_id = $1 and deleted_at = 0`, id)
	return err
}

// edirional funcs......................................................................................
// Get Courses By User ID
func (c *CourseRepo) GetCoursesByUserID(id string) ([]Course, error) {
	query := `
  SELECT c.course_id, c.title, c.description, c.created_at, c.update_at, c.deleted_at
  FROM Course c
  INNER JOIN Enrollment e ON c.course_id = e.course_id
  WHERE e.user_id = $1
  `
	rows, err := c.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []Course
	for rows.Next() {
		course := Course{}
		err := rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt,
			&course.UpdatedAt, &course.DeletedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

type CoursePopular struct {
	CourseID         string `json:"course_id"`
	CourseTitle      string `json:"course_title"`
	EnrollmentsCount int    `json:"enrollments_count"`
}

// Get The Most Popular Course
func (c *CourseRepo) Popular(start_date, end_date string) ([]CoursePopular, error) {
	if start_date == "" || end_date == "" {
		fmt.Println(start_date, end_date)
		return nil, errors.New("invalid dates")
	}
	query := `
	select c.course_id, c.title, count(e.user_id) as enrollment_count
	from course c
	join enrollment e
	on c.course_id = e.course_id
	where c.deleted_at = 0 and e.deleted_at = 0 and e.enrollment_data between $1 and $2
	group by c.course_id
	order by enrollment_count desc;`

	rows, err := c.DB.Query(query, start_date, end_date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []CoursePopular
	for rows.Next() {
		var course CoursePopular
		err := rows.Scan(&course.CourseID, &course.CourseTitle, &course.EnrollmentsCount)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

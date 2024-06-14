package enrollment

import (
	"database/sql"
	"fmt"
	"time"
)

type EnrollmentRepo struct {
	DB *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{DB: db}
}

// Insert into Enrollment
func (e *EnrollmentRepo) Create(en Enrollment) error {
	_, err := e.DB.Exec(`insert into enrollment(user_id, course_id, enrollment_data) 
	values($1, $2, $3)`,
		en.UserID, en.CourseID, en.EnrollmentDate)
	return err
}

// Get with filter
func (e *EnrollmentRepo) Get(q string, arr []interface{}) ([]Enrollment, error) {
	rows, err := e.DB.Query(q, arr...)
	if err != nil {
		return nil, err
	}
	enrollments := []Enrollment{}
	for rows.Next() {
		enroll := Enrollment{}
		err = rows.Scan(&enroll.EnrollmentID, &enroll.UserID, &enroll.CourseID,
			&enroll.EnrollmentDate, &enroll.CreatedAt, &enroll.UpdatedAt, &enroll.DeletedAt)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enroll)
	}
	return enrollments, err
}

// Get user by ID
func (e *EnrollmentRepo) GetById(id string) (Enrollment, error) {
	enroll := Enrollment{}
	err := e.DB.QueryRow("select * from enrollment where enrollment_id = $1 and deleted_at = 0", id).
		Scan(&enroll.EnrollmentID, &enroll.UserID, &enroll.CourseID,
			&enroll.EnrollmentDate, &enroll.CreatedAt, &enroll.UpdatedAt, &enroll.DeletedAt)

	return enroll, err
}

// Update user-> Name, Email, Birthday, Password
func (e *EnrollmentRepo) Update(id string, enroll Enrollment) error {
	_, err := e.DB.Exec(`update enrollment set user_id=$1, course_id=$2, enrollment_data=$3, created_at=$4, update_at=$5
	where enrollment_id=$6 and deleted_at = 0`,
		enroll.UserID, enroll.CourseID, enroll.EnrollmentDate, enroll.CreatedAt, time.Now(), enroll.CourseID)
	fmt.Println(err)
	fmt.Println(err)
	fmt.Println(err)
	return err
}

// Delete by Id
func (e *EnrollmentRepo) Delete(id string) error {
	_, err := e.DB.Exec(`update enrollment set deleted_at = date_part('epoch', current_timestamp)::INT
	where enrollment_id = $1 and deleted_at = 0`, id)
	return err
}
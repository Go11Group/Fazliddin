package enrollment

import (
	"database/sql"
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
	_, err := e.DB.Exec(`insert intoenrollmant(enrollmant_id, user_id, course_id, enrollment_data) 
	values($1, $2, $3, $4)`,
		en.Enrollment_id, en.User_id, en.Course_id, en.Enrollment_data)
	return err
}

// Get with filter
func (e *EnrollmentRepo) Get() ([]Enrollment, error) {
	rows, err := e.DB.Query("select * from users where deleted_at is null")
	if err != nil {
		return nil, err
	}
	enrollments := []Enrollment{}
	for rows.Next() {
		enroll := Enrollment{}
		err = rows.Scan(&enroll.Enrollment_id, &enroll.User_id, &enroll.Course_id,
			&enroll.Enrollment_data, &enroll.Created_at, &enroll.Update_at)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enroll)
	}
	return enrollments, err
}

// Get user by ID
func (e *EnrollmentRepo) GetById(id int) (Enrollment, error) {
	enroll := Enrollment{}
	err := e.DB.QueryRow("select * from enrollment where id = $1 and deleted_at is null", id).
		Scan(&enroll.Enrollment_id, &enroll.User_id, &enroll.Course_id,
			&enroll.Enrollment_data, &enroll.Created_at, &enroll.Update_at)

	return enroll, err
}

// Update user-> Name, Email, Birthday, Password
func (e *EnrollmentRepo) Update(enroll Enrollment) error {
	_, err := e.DB.Exec("update enrollment set user_id=$1, course_id=$2, enrollment_data=$3, created_at=$4, update_at=$5 where id=$6",
		enroll.User_id, enroll.Course_id, enroll.Enrollment_data, enroll.Created_at, time.Now(), enroll.Course_id)
	return err
}

// Delete by Id
func (e *EnrollmentRepo) Delete(id int) error {
	_, err := e.DB.Exec("update enrollment set deleted_at=$1 where id = $2", time.Now(), id)
	return err
}

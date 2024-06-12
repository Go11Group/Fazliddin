package enrollment

import "time"

type Enrollment struct{
	Enrollment_id string
    User_id string
    Course_id string
    Enrollment_data time.Time
    Created_at time.Time
    Update_at time.Time
    Deleted_at time.Time
}
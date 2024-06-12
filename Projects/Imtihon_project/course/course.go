package course

import "time"

type Course struct{
	Course_id string
    Title string
    Description string
    Created_at time.Time
    Update_at time.Time
    Deleted_at time.Time
}
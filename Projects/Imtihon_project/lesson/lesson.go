package lesson

import "time"

type Lesson struct{
	Lesson_id string
    Course_id string
    Title string
    Content string
    Created_at time.Time
    Update_at time.Time
    Deleted_at time.Time
}
package user

import "time"

type User struct{
	User_id string
    Name string
    Email string
    Birthday *time.Time
    Password string
    Created_at *time.Time
    Update_at *time.Time
    Deleted_at *time.Time
}
package user

import (
	"database/sql"
	"time"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

// Insert into users
func (u *UserRepo) Create(user User) error {
	_, err := u.DB.Exec("insert into users(user_id, name, email, birthday, password) values($1, $2, $3, $4, $5)",
		user.User_id, user.Name, user.Email, user.Birthday, user.Password)
	return err
}

// Get with filter
func (u *UserRepo) Get() ([]User, error) {
	rows, err := u.DB.Query("select * from users")
	if err != nil {
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.User_id, &user.Name, &user.Email, &user.Birthday,
			&user.Password, &user.Created_at, &user.Update_at, &user.Deleted_at)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Get user by ID
func (u *UserRepo) GetById(id int) (User, error) {
	user := User{}
	err := u.DB.QueryRow("select * from users where id = $1", id).
		Scan(&user.User_id, &user.Name, &user.Email, &user.Birthday,
			&user.Password, &user.Created_at, &user.Update_at, &user.Deleted_at)
	return user, err
}

// Update user-> Name, Email, Birthday, Password
func (u *UserRepo) Update(user User) error {
	_, err := u.DB.Exec("update users set name=$1, email=$2, birthday=$3, password=$4, update_at=$5 where id=$6",
		user.Name, user.Email, user.Birthday, user.Password, time.Now(), user.User_id)
	return err
}

// Delete by Id
func (u *UserRepo) Delete(id int) error {
	_, err := u.DB.Exec("delete from users where id = $1", id)
	return err
}

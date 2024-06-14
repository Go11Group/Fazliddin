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
	_, err := u.DB.Exec("insert into users(name, email, birthday, password) values($1, $2, $3, $4)",
		user.Name, user.Email, user.Birthday, user.Password)
	return err
}

// Get with filter
func (u *UserRepo) Get(a string, arr []interface{}) ([]User, error) {
	rows, err := u.DB.Query(a, arr...)
	if err != nil {
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday,
			&user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Get user by ID
func (u *UserRepo) GetById(id string) (User, error) {
	user := User{}
	err := u.DB.QueryRow("select * from users where user_id = $1 and deleted_at=0", id).
		Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday,
			&user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	return user, err
}

// Update user-> Name, Email, Birthday, Password
func (u *UserRepo) Update(id string, user User) error {
	_, err := u.DB.Exec(`update users set name=$1, email=$2, birthday=$3, password=$4, update_at=$5
	where user_id=$6 and deleted_at = 0`,
		user.Name, user.Email, user.Birthday, user.Password, time.Now(), id)
	return err
}

// soft delete by Id
func (u *UserRepo) Delete(id string) error {
	_, err := u.DB.Exec(`update users set deleted_at = date_part('epoch', current_timestamp)::INT
	where user_id = $1 and deleted_at = 0`, id)
	return err
}

package postgres

import (
	"database/sql"
	"fmt"
	user "homework_41/server_client/server/modul"
)

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) Create(us user.User) error {
	fmt.Println(us)
	fmt.Println("hello")
	_, err := u.DB.Exec(`insert into users (first_name, last_name, age, email)
	values ($1, $2, $3, $4)`, us.FirstName, us.LastName, us.Age, us.Email)
	return err
}

func (u *UserRepo) Get() ([]user.User, error) {
	rows, err := u.DB.Query("select * from users")
	if err != nil {
		return nil, err
	}

	users := []user.User{}
	for rows.Next() {
		user := user.User{}

		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func(u *UserRepo) GetByID(id string) (user.User, error) {
	user := user.User{}
	err := u.DB.QueryRow("select * from users where id = $1", id).
	Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age, &user.Email)
	return user, err
}

func(u *UserRepo) Update(user user.User) error{
	_, err := u.DB.Exec("update users set first_name = $1, laset_name = $2, age = $3, email = $4",
	user.FirstName, user.LastName, user.Age, user.Email)
	return err
}

func(u *UserRepo) Delete(id string) error{
	_, err := u.DB.Exec("delete from users where id = $1", id)
	return err
}
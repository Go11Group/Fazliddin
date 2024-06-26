package storage

import (
	"database/sql"
	"homeworks/homework_43/user/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) Create(user models.User) error {
	_, err := u.DB.Exec("insert into users(name, age, phone) values($1, $2, $3)",
		user.Name, user.Age, user.Phone)
	return err
}

func (u *UserRepo) Get(query string, arr []interface{}) ([]models.User, error) {
	rows, err := u.DB.Query(query, arr)
	if err != nil {
		return nil, err
	}

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Phone, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepo) GetById(id string) (models.User, error) {
	user := models.User{}
	err := u.DB.QueryRow("select * from users where user_id = $1 and deleted_at=0", id).
		Scan(&user.Id, &user.Name, &user.Age, &user.Phone, &user.DeletedAt)
	return user, err
}

func (u *UserRepo) Update(user models.User, id string) error {
	_, err := u.DB.Exec("update users set name=$1, age=$2, phone=$3 where user_id=$4 and deleted_at=0",
		user.Name, user.Age, user.Phone, id)

	return err
}

func (u *UserRepo) Delete(id string) error {
	_, err := u.DB.Exec(`update users set
deleted_at = date_part('epoch', current_timestamp)::INT
where user_id = $1 and deleted_at = 0
`, id)
	return err
}

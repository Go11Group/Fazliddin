package user

import (
	"database/sql"
	"fmt"
)

type UserRepo struct{
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo{
	return &UserRepo{DB: db}
}

func (u *UserRepo) Create(user User) error {
	_, err := u.DB.Exec("insert into users(name, age, solved_problems, rating, email, password) values($1, $2, $3, $4, $5, $6)",
	user.Name, user.Age, 0, 0.0, user.Email, user.Password)
	if err != nil{
		fmt.Println("ERROR INSERT : ", err)
	}
	return err
}

func (u *UserRepo) Get() ([]User, error){
	rows, err := u.DB.Query("select * from users")
	if err != nil{
		return nil, err
	}

	users := []User{}
	for rows.Next(){
		user := User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.SolvedProblems, &user.Rating, &user.Email, &user.Password)
		if err != nil{
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepo) GetById(id int) (User, error){
	user := User{}
	err := u.DB.QueryRow("select * from users where id = $1", id).
	Scan(&user.Id, &user.Name, &user.Age, &user.SolvedProblems, &user.Rating, &user.Email, &user.Password)
	return user, err
}

func (u *UserRepo) Update(user User) error {
	_, err := u.DB.Exec("update users set name=$1, age=$2, solved_problems=$3, rating=$4, email=$5, password=$6 where id=$7",
	user.Name, user.Age, user.SolvedProblems, user.Rating, user.Email, user.Password, user.Id)
	
	return err
}

func (u *UserRepo) Delete(id int) error{
	_, err := u.DB.Exec("delete from users where id = $1", id)
	return err
}
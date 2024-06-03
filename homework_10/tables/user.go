package tables

import (
	_"github.com/lib/pq"
	"database/sql"
	"fmt"
)

type User struct{
	Id int
	Name string
	Email string
	Password string
}

type UserRepo struct{
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo{
	return &UserRepo{db}
}

func (u *UserRepo) Create(user User) error{
	a, _:= u.DB.Begin()
	defer a.Commit()
	fmt.Println(user)
	_, err :=u.DB.Exec("insert into users(username, email, password) values($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil{
		return err
	}
	return nil
}

func (u *UserRepo) GetAllUsers() ([]User, error){
	a, _:= u.DB.Begin()
	defer a.Commit()
	users := []User{}
	data, err := u.DB.Query("select * from users")
	if err != nil{
		return nil, err
	}
	for data.Next(){
		user := User{}
		err := data.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil{
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (u *UserRepo) GetById(id int) (User, error){
	a, _:= u.DB.Begin()
	defer a.Commit()
	user := User{}
	err :=u.DB.QueryRow("select * from users where id = $1", id).Scan(user.Id, user.Name, user.Email, user.Password)
	return user, err
}

func (u *UserRepo) Update(user User) error{
	a, _:= u.DB.Begin()
	defer a.Commit()
	_, err := u.DB.Exec("update users set username = $1, email = $2, password =$3 where id = $4",
	user.Name, user.Email, user.Password, user.Id)
	if err != nil{
		return err
	}
	return nil
}

func (u *UserRepo) Delete(id int) error{
	a, _:= u.DB.Begin()
	defer a.Commit()
	_, err := u.DB.Exec("delete from users where id = $1", id)
	return err
}

type UserO struct{
	Id int
	Name string
	Email string
	OrderId []int
	OrderName []string
	Password string

}

func (u *UserRepo) ShowAllWithOrders() ([]UserO, error){
	a, _:= u.DB.Begin()
	defer a.Commit()
	data, err := u.DB.Query(`select u.id, u.username, u.email, array_agg(p.id), array_agg(p.name), u.password
	from users u
	left join users_products up on up.user_id = u.id
	left join products p on p.id = up.product_id
	group by u.id`)

	users := []UserO{}

	for data.Next(){
		user := UserO{}
		err := data.Scan(&user.Id, &user.Name, &user.Email, &user.OrderId, &user.OrderName, &user.Password)
		if err != nil{
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}
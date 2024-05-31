package main

import (
	"fmt"
	"my_mod/connect"
	user "my_mod/for_users"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}

	u := user.NewUserRepo(db)

	us := user.User{
		FirsName: "Jhon",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password123",
		Age:       30,
		Field:     "Software Engineering",
		Gender:    "Male",
		IsEmploye: true,
	}
	u.CreateTable(user.User{})

	u.Create(us)

	a := u.GetById("john@example.com")
	fmt.Println(a)
	
	users := u.GetUsers()
	fmt.Println(users)

	u.Update(us)

	u.Delete("john@example.com")
}

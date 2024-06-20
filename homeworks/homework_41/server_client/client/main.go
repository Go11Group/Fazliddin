package main

import (
	"fmt"
	"homework_41/server_client/client/user"
)

func main() {

	// u := user.User{
	// 	FirstName: "BittaBola",
	// 	LastName: "Kimdurov",
	// 	Age: 55,
	// 	Email: "nmadur@gmail.com",
	// }

	// err := user.Post(u)
	// fmt.Println(err)

	// users, err := user.Get()
	// if err != nil{
	// 	panic(err)
	// }
	// for _, i := range users{
	// 	fmt.Println(i)
	// }

	// user, err := user.GetByID("7832cab7-295a-4755-8320-524863cd53cf")
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println(user)

	err := user.Delete("d8eaf3f8-c9a6-4bd4-bcce-2ad29a919fef")
	fmt.Println(err)
}
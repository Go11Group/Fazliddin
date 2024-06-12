package main

import (
	_"fmt"
	con "model/postgres"
	t "model/tables"
)

func main() {
	db, err := con.Connect()
	if err != nil{
		panic(err)
	}

	// u := t.NewUserRepo(db)

	// user := t.User{}
	// user.Name = "Bittabole"
	// user.Email = "bittabilan@.com"
	// user.Password = "123321"
	// fmt.Println(user)
	// err = u.Create(user)
	// if err != nil{
	// 	fmt.Println(err)
	// }

	// us := t.User{}
	// us.Id = 1
	// us.Name = "aaaa"
	// us.Email = "bitta"
	// us.Password = "000"
	// err = u.Update(us)
	// if err != nil{
	// 	fmt.Println(err)
	// }

	a := t.NewPRepo(db)

	p := t.Product{}
	p.Name = "pochka"
	p.Price = 150.00
	p.Stock_quantity = 2

	err = a.Create(p)
	if err != nil{
		panic(err)
	}
}
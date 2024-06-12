package main

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "dars_2"
	password = "123321"
)

type users struct {
	Id   int
	Name string
	Age  int
	OrderId string
	OrderName string
}

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("hello")
	defer db.Close()
	user := users{}
	err = db.QueryRow("select u.id, u.name,  o.name, o.id from users u join orders o on u.id = o.user_id").
		Scan(&user.Id, &user.Name, &user.OrderName, &user.OrderId)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// _, err = db.Exec("insert into users(name, age) values($1, $2)", "Abdulaziz", 14)
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = db.Exec("insert into orders(name, user_id) values($1, $2)", "MacBookAir", 1)
	// if err != nil {
	// 	panic(err)
	// }

	data, err := db.Query("select u.id, u.age, u.name, o.name, o.id from users u join orders o on u.id = o.user_id")
	for data.Next(){
		err = data.Scan(&user.Id, &user.Age, &user.Name, &user.OrderName, &user.OrderId)
		if err != nil {
            log.Fatal(err)
        }
		fmt.Println(user)
	}
}
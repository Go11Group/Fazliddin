package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/go-faker/faker/v3"
	_ "github.com/go-faker/faker/v3"
	_ "github.com/lib/pq"
)

func BoolReturn() bool {
	if rand.Intn(2) == 2 {
		return false
	}
	return true
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123321@localhost:5432/dars_2?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	for i := 0; i < 100000; i++ {
		_, err := db.Exec("insert into persons(first_name, last_name, email, age, is_maried) values($1, $2, $3, $4, $5)",
			faker.FirstName(), faker.LastName(), faker.Email(), rand.Intn(110), BoolReturn())
		if err != nil {
			fmt.Println(err)
		}
		if i%1000 == 0{
			fmt.Println(i)
		}
	}

}

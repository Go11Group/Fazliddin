package connect

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error){
	db, err := gorm.Open(postgres.Open("postgres://postgres:123321@localhost:5432/dars_2?sslmode=disable"))
	if err != nil{
		return nil, err
	}
	return db, nil
}
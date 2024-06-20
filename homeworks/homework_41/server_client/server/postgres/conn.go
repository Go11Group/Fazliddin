package postgres

import "database/sql"

func Conn() (*sql.DB, error){
	db, err := sql.Open("postgres", "postgres://postgres:123321@localhost:5432/dars_4?sslmode=disable&search_path=public")
	return db, err
}
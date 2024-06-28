package storage

import (
	"database/sql"
	pb "homeworks/homewrok_45/genproto/generator"
)

type BookRepo struct{
	DB *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo{
	return &BookRepo{db}
}

func(b *BookRepo) Create(book pb.Book) error{
	_, err := b.DB.Exec("insert into book (id, title, author) values($1, $2, $3)",
	book.BookId, book.Title, book.Author)
	return err
}

func(b *BookRepo) Get () ([]*pb.Book , error){
	rows, err := b.DB.Query("select * from book")
	if err != nil{
		return nil, err
	}

	books := []*pb.Book{}
	for rows.Next() {
		book := pb.Book{}
		err = rows.Scan(&book.BookId, &book.Title, &book.Author)
		if err != nil{
			return nil, err
		}
		books = append(books, &book)	
	}
	return books, err
}
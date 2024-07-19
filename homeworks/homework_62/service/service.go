package service

import (
	"homework_62/models"
	"homework_62/storage/redis_1"
)

type Service struct {
	Dtb *redis_1.BookRepo
}

func (s *Service) PostBook(book models.Book) (string, error) {
	err := s.Dtb.Create(book)
	if err != nil {
		return "", err
	}
	return "user created", nil
}

func (s *Service) GetBook(bookName models.BookName) (*models.Book, error) {
	book, err := s.Dtb.GetByName(bookName)
	return book, err
}

func (s *Service) DeleteBook(bookName models.BookName) (string, error) {
	err := s.Dtb.Delete(bookName)
	if err != nil {
		return "", err
	}
	return "book deleted", nil
}

func (s *Service) GetBooks() (*[]string, error) {
	result, err := s.Dtb.Get()
	return result, err
}

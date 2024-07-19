package service

import (
	"homework_62/models"
	"homework_62/storage/redis_1"
)

type Service struct {
	dtb *redis_1.BookRepo
}

func (s *Service) PostBook(book models.Book) (string, error) {
	err := s.dtb.Create(book)
	if err != nil {
		return "", err
	}
	return "user created", nil
}

func (s *Service) GetBook(bookName models.BookName) (*models.Book, error) {
	book, err := s.dtb.GetByName(bookName)
	return book, err
}

func (s *Service) DeleteBook(bookName models.BookName) (string, error) {
	err := s.dtb.Delete(bookName)
	if err != nil {
		return "", err
	}
	return "book deleted", nil
}

func (s *Service) GetBooks() (*[]string, error) {
	result, err := s.dtb.Get()
	return result, err
}

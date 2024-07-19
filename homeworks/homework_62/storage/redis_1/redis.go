package redis_1

import (
	"context"
	"encoding/json"
	"fmt"
	"homework_62/models"
	"log"

	"github.com/redis/go-redis/v9"
)

type BookRepo struct {
	Redis *redis.Client
}

func NewBookRepo (rd *redis.Client) *BookRepo {
	return &BookRepo{Redis: rd}
}

func (b *BookRepo) Create (book models.Books) error {
	jsonData, err := json.Marshal(book)
	if err != nil{
		log.Println(err)
		return err
	}

	err = b.Redis.Set(context.Background(), book.BookName, jsonData, 0).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (b *BookRepo) Update (book models.Books) error {
	jsonData, err := json.Marshal(book)
	if err != nil{
		log.Println(err)
		return err
	}

	err = b.Redis.Set(context.Background(), book.BookName, jsonData, 0).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func 
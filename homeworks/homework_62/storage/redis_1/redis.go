package redis_1

import (
	"context"
	"encoding/json"
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

func (b *BookRepo) Create (book models.Book) error {
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

func (b *BookRepo) Update (book models.Book) error {
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

func (b *BookRepo) Delete (book models.BookName) error {
	err := b.Redis.Del(context.Background(), book.Name).Err()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (b *BookRepo) Get () (*[]string, error) {
	keys := []string{}
	kursor := 0
	for {
		key, kursor, err := b.Redis.Scan(context.Background(), uint64(kursor), "*", 0).Result()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		keys = append(keys, key...)
		if kursor == 0 {
			break
		}
	}

	values := []string{}

	for _, key := range keys {
        value, err := b.Redis.Get(context.Background(), key).Result()
        if err == redis.Nil {
            log.Printf("Key %s does not exist\n", key)
			return nil, err
		}
		values = append(values, value)
    }
	return &values, nil
}

func (b *BookRepo) GetByName (bk models.BookName) (*models.Book, error) {
	result, err := b.Redis.Get(context.Background(), bk.Name).Result()
	if err != nil {
		return nil, err
	}

	var book models.Book
	err = json.Unmarshal([]byte(result), &book)
	
	return &book, err
}
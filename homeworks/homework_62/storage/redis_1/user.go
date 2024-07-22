package redis_1

import (
	"context"
	"encoding/json"
	"homework_62/models"
)

func (b *BookRepo) CreateUser(user models.UserReq) error {
	byteUser, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = b.Redis.Set(context.Background(), user.Name, byteUser, 0).Err()
	return err
}

package repositories

import (
	"context"
	"fmt"
	"pi/internal/core/domains"
	"pi/internal/core/ports"
	"time"

	"github.com/redis/go-redis/v9"
)

type userCacheRepository struct {
	rc *redis.Client
}

func NewUserCacheRepository(rc *redis.Client) ports.UserCacheRepository {
	return &userCacheRepository{rc}
}

func (u *userCacheRepository) GetByID(ctx context.Context, id uint) (*domains.User, error) {
	result, err := u.rc.HGetAll(ctx, fmt.Sprintf("user-%d", id)).Result()
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	out := domains.User{
		ID:       id,
		Username: result["username"],
		Email:    result["email"],
	}
	return &out, nil
}

func (u *userCacheRepository) Update(ctx context.Context, user *domains.User) error {
	data := map[string]string{
		"username": user.Username,
		"email":    user.Email,
	}
	if err := u.rc.HSet(ctx, fmt.Sprintf("user-%d", user.ID), data).Err(); err != nil {
		return err
	}
	if err := u.rc.Expire(ctx, fmt.Sprintf("user-%d", user.ID), 5*time.Hour).Err(); err != nil {
		return err
	}
	return nil
}

func (u *userCacheRepository) Delete(ctx context.Context, id uint) error {
	return u.rc.Del(ctx, fmt.Sprintf("user-%d", id)).Err()
}

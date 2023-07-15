package ports

import (
	"context"
	"pi/internal/core/domains"
)

type UserRepository interface {
	Create(context.Context, *domains.CreateUserRequest) (*domains.User, error)
	GetByID(context.Context, uint) (*domains.User, error)
	Update(context.Context, uint, *domains.UpdateUserRequest) error
	Delete(context.Context, uint) error
}

type UserCacheRepository interface {
	GetByID(context.Context, uint) (*domains.User, error)
	Update(context.Context, *domains.User) error
	Delete(context.Context, uint) error
}

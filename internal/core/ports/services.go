package ports

import (
	"context"
	"pi/internal/core/domains"
)

type UserService interface {
	CreateUser(context.Context, *domains.CreateUserRequest) (*domains.User, error)
	GetUserByID(context.Context, uint) (*domains.User, error)
	Update(context.Context, uint, *domains.UpdateUserRequest) error
	Delete(context.Context, uint) error
}

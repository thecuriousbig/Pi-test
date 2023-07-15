package usersvc

import (
	"context"
	"pi/internal/core/domains"
	"pi/internal/core/ports"
	"pi/internal/errmsg"
)

type service struct {
	ur  ports.UserRepository
	ucr ports.UserCacheRepository
}

func New(ur ports.UserRepository, ucr ports.UserCacheRepository) ports.UserService {
	return &service{ur, ucr}
}

func (s *service) CreateUser(ctx context.Context, req *domains.CreateUserRequest) (*domains.User, error) {
	user, err := s.ur.Create(ctx, req)
	if err != nil {
		return nil, errmsg.UserCreateFailed
	}
	return user, nil
}

func (s *service) GetUserByID(ctx context.Context, id uint) (*domains.User, error) {
	var user *domains.User
	var err error

	user, err = s.ucr.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		user, err = s.ur.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}

		if user == nil {
			return nil, errmsg.UserNotFound
		}

		err = s.ucr.Update(ctx, user)
		if err != nil {
			return nil, errmsg.UserCacheUpdate
		}
	}
	return user, nil
}

func (s *service) Update(ctx context.Context, id uint, req *domains.UpdateUserRequest) error {
	err := s.ur.Update(ctx, id, req)
	if err != nil {
		return errmsg.UserUpdateFailed
	}
	user, err := s.ur.GetByID(ctx, id)
	if err != nil {
		return err
	}
	err = s.ucr.Update(ctx, user)
	if err != nil {
		return errmsg.UserCacheUpdate
	}
	return nil
}

func (s *service) Delete(ctx context.Context, id uint) error {
	err := s.ur.Delete(ctx, id)
	if err != nil {
		return errmsg.UserDeleteFailed
	}
	err = s.ucr.Delete(ctx, id)
	if err != nil {
		return errmsg.UserCacheDelete
	}
	return nil
}

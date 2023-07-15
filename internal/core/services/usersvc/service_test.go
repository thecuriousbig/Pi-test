package usersvc_test

import (
	"context"
	"errors"
	"pi/internal/core/domains"
	"pi/internal/core/ports"
	"pi/internal/core/ports/mocks"
	"pi/internal/core/services/usersvc"
	"pi/internal/errmsg"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testModule struct {
	ur  *mocks.UserRepository
	ucr *mocks.UserCacheRepository
	svc ports.UserService
}

type test struct {
	name     string
	args     []interface{}
	mockFn   func(*testModule)
	assertFn func(*testModule)
}

var (
	ctx = context.TODO()
)

func new(t *testing.T) *testModule {
	ur := mocks.NewUserRepository(t)
	ucr := mocks.NewUserCacheRepository(t)
	return &testModule{
		ur:  ur,
		ucr: ucr,
		svc: usersvc.New(ur, ucr),
	}
}

func TestCreate(t *testing.T) {
	var result *domains.User
	var err error
	mockReq := &domains.CreateUserRequest{
		Username: "username",
		Email:    "email",
	}

	mockCreatedUser := &domains.User{
		ID:       1,
		Username: "username",
		Email:    "email",
	}

	tests := []*test{
		{
			name: "return error when create failed",
			args: []interface{}{
				ctx,
				mockReq,
			},
			mockFn: func(m *testModule) {
				m.ur.On("Create", ctx, mockReq).Return(nil, errors.New("error")).Once()
			},
			assertFn: func(m *testModule) {
				assert.Error(t, err)
				assert.EqualError(t, err, errmsg.UserCreateFailed.Error())
			},
		},
		{
			name: "success",
			args: []interface{}{
				ctx,
				mockReq,
			},
			mockFn: func(m *testModule) {
				m.ur.On("Create", ctx, mockReq).Return(mockCreatedUser, nil).Once()
			},
			assertFn: func(m *testModule) {
				assert.NoError(t, err)
				assert.Equal(t, result, mockCreatedUser)
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := new(t)
			tc.mockFn(m)
			result, err = m.svc.CreateUser(tc.args[0].(context.Context), tc.args[1].(*domains.CreateUserRequest))
			tc.assertFn(m)
		})
	}
}

func TestGetUserByID(t *testing.T) {
	var result *domains.User
	var err error
	mockId := uint(1)

	mockUser := &domains.User{
		ID:       1,
		Username: "username",
		Email:    "email",
	}

	tests := []*test{
		{
			name: "return error when get cache failed",
			args: []interface{}{
				ctx,
				mockId,
			},
			mockFn: func(m *testModule) {
				m.ucr.On("GetByID", ctx, mockId).Return(nil, errors.New("error")).Once()
			},
			assertFn: func(m *testModule) {
				assert.Error(t, err)
			},
		},
		{
			name: "return error when get user from db failed",
			args: []interface{}{
				ctx,
				mockId,
			},
			mockFn: func(m *testModule) {
				m.ucr.On("GetByID", ctx, mockId).Return(nil, nil).Once()
				m.ur.On("GetByID", ctx, mockId).Return(nil, errors.New("error")).Once()
			},
			assertFn: func(m *testModule) {
				assert.Error(t, err)
			},
		},
		{
			name: "return error when get user from db not found",
			args: []interface{}{
				ctx,
				mockId,
			},
			mockFn: func(m *testModule) {
				m.ucr.On("GetByID", ctx, mockId).Return(nil, nil).Once()
				m.ur.On("GetByID", ctx, mockId).Return(nil, nil).Once()
			},
			assertFn: func(m *testModule) {
				assert.Error(t, err)
				assert.EqualError(t, err, errmsg.UserNotFound.Error())
			},
		},
		{
			name: "return error when update user failed",
			args: []interface{}{
				ctx,
				mockId,
			},
			mockFn: func(m *testModule) {
				m.ucr.On("GetByID", ctx, mockId).Return(nil, nil).Once()
				m.ur.On("GetByID", ctx, mockId).Return(mockUser, nil).Once()
				m.ucr.On("Update", ctx, mockUser).Return(errors.New("error")).Once()
			},
			assertFn: func(m *testModule) {
				assert.Error(t, err)
				assert.EqualError(t, err, errmsg.UserCacheUpdate.Error())
			},
		},
		{
			name: "success",
			args: []interface{}{
				ctx,
				mockId,
			},
			mockFn: func(m *testModule) {
				m.ucr.On("GetByID", ctx, mockId).Return(nil, nil).Once()
				m.ur.On("GetByID", ctx, mockId).Return(mockUser, nil).Once()
				m.ucr.On("Update", ctx, mockUser).Return(nil).Once()
			},
			assertFn: func(m *testModule) {
				assert.NoError(t, err)
				assert.Equal(t, result, mockUser)
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := new(t)
			tc.mockFn(m)
			result, err = m.svc.GetUserByID(tc.args[0].(context.Context), tc.args[1].(uint))
			tc.assertFn(m)
		})
	}
}

func TestUpdate(t *testing.T) {
	var err error
	mockId := uint(1)
	mockReq := &domains.UpdateUserRequest{
		Username: "username",
		Email:    "email",
	}
	mockUser := &domains.User{
		ID:       1,
		Username: "username",
		Email:    "email",
	}

	tests := []*test{
		{
			name: "return error when update failed",
			args: []interface{}{
				ctx,
				mockId,
				mockReq,
			},
			mockFn: func(m *testModule) {
				m.ur.On("Update", ctx, mockId, mockReq).Return(errors.New("error")).Once()
			},
			assertFn: func(m *testModule) {
				assert.Error(t, err)
				assert.EqualError(t, err, errmsg.UserUpdateFailed.Message)
			},
		},
		{
			name: "return error when get user failed",
			args: []interface{}{
				ctx,
				mockId,
				mockReq,
			},
			mockFn: func(m *testModule) {
				m.ur.On("Update", ctx, mockId, mockReq).Return(nil).Once()
				m.ur.On("GetByID", ctx, mockId).Return(nil, errors.New("error")).Once()
			},
			assertFn: func(m *testModule) {
				assert.Error(t, err)
				assert.EqualError(t, err, "error")
			},
		},
		{
			name: "return error when update cache failed",
			args: []interface{}{
				ctx,
				mockId,
				mockReq,
			},
			mockFn: func(m *testModule) {
				m.ur.On("Update", ctx, mockId, mockReq).Return(nil).Once()
				m.ur.On("GetByID", ctx, mockId).Return(mockUser, nil).Once()
				m.ucr.On("Update", ctx, mockUser).Return(errors.New("error")).Once()
			},
			assertFn: func(m *testModule) {
				assert.Error(t, err)
				assert.EqualError(t, err, errmsg.UserCacheUpdate.Error())
			},
		},
		{
			name: "success",
			args: []interface{}{
				ctx,
				mockId,
				mockReq,
			},
			mockFn: func(m *testModule) {
				m.ur.On("Update", ctx, mockId, mockReq).Return(nil).Once()
				m.ur.On("GetByID", ctx, mockId).Return(mockUser, nil).Once()
				m.ucr.On("Update", ctx, mockUser).Return(nil).Once()
			},
			assertFn: func(m *testModule) {
				assert.NoError(t, err)
				assert.Equal(t, err, nil)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := new(t)
			tc.mockFn(m)
			err = m.svc.Update(tc.args[0].(context.Context), tc.args[1].(uint), tc.args[2].(*domains.UpdateUserRequest))
			tc.assertFn(m)
		})
	}
}

func TestDelete(t *testing.T) {
	var err error
	mockId := uint(1)

	tests := []*test{
		{
			name: "return error when delete failed",
			args: []interface{}{
				ctx,
				mockId,
			},
			mockFn: func(m *testModule) {
				m.ur.On("Delete", ctx, mockId).Return(errors.New("error")).Once()
			},
			assertFn: func(m *testModule) {
				assert.Error(t, err)
				assert.EqualError(t, err, errmsg.UserDeleteFailed.Message)
			},
		},
		{
			name: "return error when delete user cache failed",
			args: []interface{}{
				ctx,
				mockId,
			},
			mockFn: func(m *testModule) {
				m.ur.On("Delete", ctx, mockId).Return(nil).Once()
				m.ucr.On("Delete", ctx, mockId).Return(errors.New("error")).Once()
			},
			assertFn: func(m *testModule) {
				assert.Error(t, err)
				assert.EqualError(t, err, errmsg.UserCacheDelete.Error())
			},
		},
		{
			name: "success",
			args: []interface{}{
				ctx,
				mockId,
			},
			mockFn: func(m *testModule) {
				m.ur.On("Delete", ctx, mockId).Return(nil).Once()
				m.ucr.On("Delete", ctx, mockId).Return(nil).Once()
			},
			assertFn: func(m *testModule) {
				assert.NoError(t, err)
				assert.Equal(t, err, nil)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := new(t)
			tc.mockFn(m)
			err = m.svc.Delete(tc.args[0].(context.Context), tc.args[1].(uint))
			tc.assertFn(m)
		})
	}
}

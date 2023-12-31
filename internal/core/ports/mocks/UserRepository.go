// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	context "context"
	domains "pi/internal/core/domains"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

type UserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepository) EXPECT() *UserRepository_Expecter {
	return &UserRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *UserRepository) Create(_a0 context.Context, _a1 *domains.CreateUserRequest) (*domains.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domains.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domains.CreateUserRequest) (*domains.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domains.CreateUserRequest) *domains.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domains.CreateUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type UserRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *domains.CreateUserRequest
func (_e *UserRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *UserRepository_Create_Call {
	return &UserRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *UserRepository_Create_Call) Run(run func(_a0 context.Context, _a1 *domains.CreateUserRequest)) *UserRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domains.CreateUserRequest))
	})
	return _c
}

func (_c *UserRepository_Create_Call) Return(_a0 *domains.User, _a1 error) *UserRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepository_Create_Call) RunAndReturn(run func(context.Context, *domains.CreateUserRequest) (*domains.User, error)) *UserRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *UserRepository) Delete(_a0 context.Context, _a1 uint) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type UserRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint
func (_e *UserRepository_Expecter) Delete(_a0 interface{}, _a1 interface{}) *UserRepository_Delete_Call {
	return &UserRepository_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *UserRepository_Delete_Call) Run(run func(_a0 context.Context, _a1 uint)) *UserRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *UserRepository_Delete_Call) Return(_a0 error) *UserRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_Delete_Call) RunAndReturn(run func(context.Context, uint) error) *UserRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: _a0, _a1
func (_m *UserRepository) GetByID(_a0 context.Context, _a1 uint) (*domains.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domains.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*domains.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *domains.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type UserRepository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint
func (_e *UserRepository_Expecter) GetByID(_a0 interface{}, _a1 interface{}) *UserRepository_GetByID_Call {
	return &UserRepository_GetByID_Call{Call: _e.mock.On("GetByID", _a0, _a1)}
}

func (_c *UserRepository_GetByID_Call) Run(run func(_a0 context.Context, _a1 uint)) *UserRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *UserRepository_GetByID_Call) Return(_a0 *domains.User, _a1 error) *UserRepository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepository_GetByID_Call) RunAndReturn(run func(context.Context, uint) (*domains.User, error)) *UserRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1, _a2
func (_m *UserRepository) Update(_a0 context.Context, _a1 uint, _a2 *domains.UpdateUserRequest) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, *domains.UpdateUserRequest) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type UserRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint
//   - _a2 *domains.UpdateUserRequest
func (_e *UserRepository_Expecter) Update(_a0 interface{}, _a1 interface{}, _a2 interface{}) *UserRepository_Update_Call {
	return &UserRepository_Update_Call{Call: _e.mock.On("Update", _a0, _a1, _a2)}
}

func (_c *UserRepository_Update_Call) Run(run func(_a0 context.Context, _a1 uint, _a2 *domains.UpdateUserRequest)) *UserRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint), args[2].(*domains.UpdateUserRequest))
	})
	return _c
}

func (_c *UserRepository_Update_Call) Return(_a0 error) *UserRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_Update_Call) RunAndReturn(run func(context.Context, uint, *domains.UpdateUserRequest) error) *UserRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

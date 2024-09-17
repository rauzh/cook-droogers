// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	models "cookdroogers/models"

	mock "github.com/stretchr/testify/mock"
)

// UserRepo is an autogenerated mock type for the UserRepo type
type UserRepo struct {
	mock.Mock
}

type UserRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepo) EXPECT() *UserRepo_Expecter {
	return &UserRepo_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *UserRepo) Create(_a0 context.Context, _a1 *models.User) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type UserRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *models.User
func (_e *UserRepo_Expecter) Create(_a0 interface{}, _a1 interface{}) *UserRepo_Create_Call {
	return &UserRepo_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *UserRepo_Create_Call) Run(run func(_a0 context.Context, _a1 *models.User)) *UserRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.User))
	})
	return _c
}

func (_c *UserRepo_Create_Call) Return(_a0 error) *UserRepo_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepo_Create_Call) RunAndReturn(run func(context.Context, *models.User) error) *UserRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *UserRepo) Get(_a0 context.Context, _a1 uint64) (*models.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*models.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type UserRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
func (_e *UserRepo_Expecter) Get(_a0 interface{}, _a1 interface{}) *UserRepo_Get_Call {
	return &UserRepo_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *UserRepo_Get_Call) Run(run func(_a0 context.Context, _a1 uint64)) *UserRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *UserRepo_Get_Call) Return(_a0 *models.User, _a1 error) *UserRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepo_Get_Call) RunAndReturn(run func(context.Context, uint64) (*models.User, error)) *UserRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByEmail provides a mock function with given fields: _a0, _a1
func (_m *UserRepo) GetByEmail(_a0 context.Context, _a1 string) (*models.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_GetByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmail'
type UserRepo_GetByEmail_Call struct {
	*mock.Call
}

// GetByEmail is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *UserRepo_Expecter) GetByEmail(_a0 interface{}, _a1 interface{}) *UserRepo_GetByEmail_Call {
	return &UserRepo_GetByEmail_Call{Call: _e.mock.On("GetByEmail", _a0, _a1)}
}

func (_c *UserRepo_GetByEmail_Call) Run(run func(_a0 context.Context, _a1 string)) *UserRepo_GetByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserRepo_GetByEmail_Call) Return(_a0 *models.User, _a1 error) *UserRepo_GetByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepo_GetByEmail_Call) RunAndReturn(run func(context.Context, string) (*models.User, error)) *UserRepo_GetByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetForAdmin provides a mock function with given fields: ctx
func (_m *UserRepo) GetForAdmin(ctx context.Context) ([]models.User, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetForAdmin")
	}

	var r0 []models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_GetForAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetForAdmin'
type UserRepo_GetForAdmin_Call struct {
	*mock.Call
}

// GetForAdmin is a helper method to define mock.On call
//   - ctx context.Context
func (_e *UserRepo_Expecter) GetForAdmin(ctx interface{}) *UserRepo_GetForAdmin_Call {
	return &UserRepo_GetForAdmin_Call{Call: _e.mock.On("GetForAdmin", ctx)}
}

func (_c *UserRepo_GetForAdmin_Call) Run(run func(ctx context.Context)) *UserRepo_GetForAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *UserRepo_GetForAdmin_Call) Return(_a0 []models.User, _a1 error) *UserRepo_GetForAdmin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepo_GetForAdmin_Call) RunAndReturn(run func(context.Context) ([]models.User, error)) *UserRepo_GetForAdmin_Call {
	_c.Call.Return(run)
	return _c
}

// SetRole provides a mock function with given fields: ctx, role
func (_m *UserRepo) SetRole(ctx context.Context, role models.UserType) error {
	ret := _m.Called(ctx, role)

	if len(ret) == 0 {
		panic("no return value specified for SetRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UserType) error); ok {
		r0 = rf(ctx, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepo_SetRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetRole'
type UserRepo_SetRole_Call struct {
	*mock.Call
}

// SetRole is a helper method to define mock.On call
//   - ctx context.Context
//   - role models.UserType
func (_e *UserRepo_Expecter) SetRole(ctx interface{}, role interface{}) *UserRepo_SetRole_Call {
	return &UserRepo_SetRole_Call{Call: _e.mock.On("SetRole", ctx, role)}
}

func (_c *UserRepo_SetRole_Call) Run(run func(ctx context.Context, role models.UserType)) *UserRepo_SetRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.UserType))
	})
	return _c
}

func (_c *UserRepo_SetRole_Call) Return(_a0 error) *UserRepo_SetRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepo_SetRole_Call) RunAndReturn(run func(context.Context, models.UserType) error) *UserRepo_SetRole_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *UserRepo) Update(_a0 context.Context, _a1 *models.User) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepo_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type UserRepo_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *models.User
func (_e *UserRepo_Expecter) Update(_a0 interface{}, _a1 interface{}) *UserRepo_Update_Call {
	return &UserRepo_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *UserRepo_Update_Call) Run(run func(_a0 context.Context, _a1 *models.User)) *UserRepo_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.User))
	})
	return _c
}

func (_c *UserRepo_Update_Call) Return(_a0 error) *UserRepo_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepo_Update_Call) RunAndReturn(run func(context.Context, *models.User) error) *UserRepo_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateType provides a mock function with given fields: ctx, userID, typ
func (_m *UserRepo) UpdateType(ctx context.Context, userID uint64, typ models.UserType) error {
	ret := _m.Called(ctx, userID, typ)

	if len(ret) == 0 {
		panic("no return value specified for UpdateType")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, models.UserType) error); ok {
		r0 = rf(ctx, userID, typ)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepo_UpdateType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateType'
type UserRepo_UpdateType_Call struct {
	*mock.Call
}

// UpdateType is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint64
//   - typ models.UserType
func (_e *UserRepo_Expecter) UpdateType(ctx interface{}, userID interface{}, typ interface{}) *UserRepo_UpdateType_Call {
	return &UserRepo_UpdateType_Call{Call: _e.mock.On("UpdateType", ctx, userID, typ)}
}

func (_c *UserRepo_UpdateType_Call) Run(run func(ctx context.Context, userID uint64, typ models.UserType)) *UserRepo_UpdateType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(models.UserType))
	})
	return _c
}

func (_c *UserRepo_UpdateType_Call) Return(_a0 error) *UserRepo_UpdateType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepo_UpdateType_Call) RunAndReturn(run func(context.Context, uint64, models.UserType) error) *UserRepo_UpdateType_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserRepo creates a new instance of UserRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepo {
	mock := &UserRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

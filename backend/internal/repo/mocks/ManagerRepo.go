// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "cookdroogers/models"

	mock "github.com/stretchr/testify/mock"
)

// ManagerRepo is an autogenerated mock type for the ManagerRepo type
type ManagerRepo struct {
	mock.Mock
}

type ManagerRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *ManagerRepo) EXPECT() *ManagerRepo_Expecter {
	return &ManagerRepo_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *ManagerRepo) Create(_a0 context.Context, _a1 *models.Manager) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Manager) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ManagerRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ManagerRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *models.Manager
func (_e *ManagerRepo_Expecter) Create(_a0 interface{}, _a1 interface{}) *ManagerRepo_Create_Call {
	return &ManagerRepo_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *ManagerRepo_Create_Call) Run(run func(_a0 context.Context, _a1 *models.Manager)) *ManagerRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Manager))
	})
	return _c
}

func (_c *ManagerRepo_Create_Call) Return(_a0 error) *ManagerRepo_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ManagerRepo_Create_Call) RunAndReturn(run func(context.Context, *models.Manager) error) *ManagerRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *ManagerRepo) Get(_a0 context.Context, _a1 uint64) (*models.Manager, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.Manager
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*models.Manager, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Manager); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Manager)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManagerRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ManagerRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
func (_e *ManagerRepo_Expecter) Get(_a0 interface{}, _a1 interface{}) *ManagerRepo_Get_Call {
	return &ManagerRepo_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *ManagerRepo_Get_Call) Run(run func(_a0 context.Context, _a1 uint64)) *ManagerRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *ManagerRepo_Get_Call) Return(_a0 *models.Manager, _a1 error) *ManagerRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ManagerRepo_Get_Call) RunAndReturn(run func(context.Context, uint64) (*models.Manager, error)) *ManagerRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUserID provides a mock function with given fields: ctx, userID
func (_m *ManagerRepo) GetByUserID(ctx context.Context, userID uint64) (*models.Manager, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetByUserID")
	}

	var r0 *models.Manager
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*models.Manager, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Manager); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Manager)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManagerRepo_GetByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUserID'
type ManagerRepo_GetByUserID_Call struct {
	*mock.Call
}

// GetByUserID is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint64
func (_e *ManagerRepo_Expecter) GetByUserID(ctx interface{}, userID interface{}) *ManagerRepo_GetByUserID_Call {
	return &ManagerRepo_GetByUserID_Call{Call: _e.mock.On("GetByUserID", ctx, userID)}
}

func (_c *ManagerRepo_GetByUserID_Call) Run(run func(ctx context.Context, userID uint64)) *ManagerRepo_GetByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *ManagerRepo_GetByUserID_Call) Return(_a0 *models.Manager, _a1 error) *ManagerRepo_GetByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ManagerRepo_GetByUserID_Call) RunAndReturn(run func(context.Context, uint64) (*models.Manager, error)) *ManagerRepo_GetByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// GetForAdmin provides a mock function with given fields: ctx
func (_m *ManagerRepo) GetForAdmin(ctx context.Context) ([]models.Manager, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetForAdmin")
	}

	var r0 []models.Manager
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Manager, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Manager); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Manager)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManagerRepo_GetForAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetForAdmin'
type ManagerRepo_GetForAdmin_Call struct {
	*mock.Call
}

// GetForAdmin is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ManagerRepo_Expecter) GetForAdmin(ctx interface{}) *ManagerRepo_GetForAdmin_Call {
	return &ManagerRepo_GetForAdmin_Call{Call: _e.mock.On("GetForAdmin", ctx)}
}

func (_c *ManagerRepo_GetForAdmin_Call) Run(run func(ctx context.Context)) *ManagerRepo_GetForAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ManagerRepo_GetForAdmin_Call) Return(_a0 []models.Manager, _a1 error) *ManagerRepo_GetForAdmin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ManagerRepo_GetForAdmin_Call) RunAndReturn(run func(context.Context) ([]models.Manager, error)) *ManagerRepo_GetForAdmin_Call {
	_c.Call.Return(run)
	return _c
}

// GetRandManagerID provides a mock function with given fields: _a0
func (_m *ManagerRepo) GetRandManagerID(_a0 context.Context) (uint64, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetRandManagerID")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManagerRepo_GetRandManagerID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRandManagerID'
type ManagerRepo_GetRandManagerID_Call struct {
	*mock.Call
}

// GetRandManagerID is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *ManagerRepo_Expecter) GetRandManagerID(_a0 interface{}) *ManagerRepo_GetRandManagerID_Call {
	return &ManagerRepo_GetRandManagerID_Call{Call: _e.mock.On("GetRandManagerID", _a0)}
}

func (_c *ManagerRepo_GetRandManagerID_Call) Run(run func(_a0 context.Context)) *ManagerRepo_GetRandManagerID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ManagerRepo_GetRandManagerID_Call) Return(_a0 uint64, _a1 error) *ManagerRepo_GetRandManagerID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ManagerRepo_GetRandManagerID_Call) RunAndReturn(run func(context.Context) (uint64, error)) *ManagerRepo_GetRandManagerID_Call {
	_c.Call.Return(run)
	return _c
}

// NewManagerRepo creates a new instance of ManagerRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManagerRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *ManagerRepo {
	mock := &ManagerRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

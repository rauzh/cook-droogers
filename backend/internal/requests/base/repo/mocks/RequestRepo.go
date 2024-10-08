// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	base "cookdroogers/internal/requests/base"

	mock "github.com/stretchr/testify/mock"
)

// RequestRepo is an autogenerated mock type for the RequestRepo type
type RequestRepo struct {
	mock.Mock
}

type RequestRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *RequestRepo) EXPECT() *RequestRepo_Expecter {
	return &RequestRepo_Expecter{mock: &_m.Mock}
}

// GetAllByManagerID provides a mock function with given fields: _a0, _a1
func (_m *RequestRepo) GetAllByManagerID(_a0 context.Context, _a1 uint64) ([]base.Request, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByManagerID")
	}

	var r0 []base.Request
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]base.Request, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []base.Request); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]base.Request)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestRepo_GetAllByManagerID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByManagerID'
type RequestRepo_GetAllByManagerID_Call struct {
	*mock.Call
}

// GetAllByManagerID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
func (_e *RequestRepo_Expecter) GetAllByManagerID(_a0 interface{}, _a1 interface{}) *RequestRepo_GetAllByManagerID_Call {
	return &RequestRepo_GetAllByManagerID_Call{Call: _e.mock.On("GetAllByManagerID", _a0, _a1)}
}

func (_c *RequestRepo_GetAllByManagerID_Call) Run(run func(_a0 context.Context, _a1 uint64)) *RequestRepo_GetAllByManagerID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *RequestRepo_GetAllByManagerID_Call) Return(_a0 []base.Request, _a1 error) *RequestRepo_GetAllByManagerID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RequestRepo_GetAllByManagerID_Call) RunAndReturn(run func(context.Context, uint64) ([]base.Request, error)) *RequestRepo_GetAllByManagerID_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllByUserID provides a mock function with given fields: _a0, _a1
func (_m *RequestRepo) GetAllByUserID(_a0 context.Context, _a1 uint64) ([]base.Request, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByUserID")
	}

	var r0 []base.Request
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]base.Request, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []base.Request); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]base.Request)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestRepo_GetAllByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByUserID'
type RequestRepo_GetAllByUserID_Call struct {
	*mock.Call
}

// GetAllByUserID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
func (_e *RequestRepo_Expecter) GetAllByUserID(_a0 interface{}, _a1 interface{}) *RequestRepo_GetAllByUserID_Call {
	return &RequestRepo_GetAllByUserID_Call{Call: _e.mock.On("GetAllByUserID", _a0, _a1)}
}

func (_c *RequestRepo_GetAllByUserID_Call) Run(run func(_a0 context.Context, _a1 uint64)) *RequestRepo_GetAllByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *RequestRepo_GetAllByUserID_Call) Return(_a0 []base.Request, _a1 error) *RequestRepo_GetAllByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RequestRepo_GetAllByUserID_Call) RunAndReturn(run func(context.Context, uint64) ([]base.Request, error)) *RequestRepo_GetAllByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, uint642
func (_m *RequestRepo) GetByID(ctx context.Context, uint642 uint64) (*base.Request, error) {
	ret := _m.Called(ctx, uint642)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *base.Request
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*base.Request, error)); ok {
		return rf(ctx, uint642)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *base.Request); ok {
		r0 = rf(ctx, uint642)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*base.Request)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, uint642)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestRepo_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type RequestRepo_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - uint642 uint64
func (_e *RequestRepo_Expecter) GetByID(ctx interface{}, uint642 interface{}) *RequestRepo_GetByID_Call {
	return &RequestRepo_GetByID_Call{Call: _e.mock.On("GetByID", ctx, uint642)}
}

func (_c *RequestRepo_GetByID_Call) Run(run func(ctx context.Context, uint642 uint64)) *RequestRepo_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *RequestRepo_GetByID_Call) Return(_a0 *base.Request, _a1 error) *RequestRepo_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RequestRepo_GetByID_Call) RunAndReturn(run func(context.Context, uint64) (*base.Request, error)) *RequestRepo_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewRequestRepo creates a new instance of RequestRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequestRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequestRepo {
	mock := &RequestRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

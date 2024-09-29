// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	criteria "cookdroogers/internal/requests/criteria_controller"

	mock "github.com/stretchr/testify/mock"
)

// AbstractCriteriaFabric is an autogenerated mock type for the AbstractCriteriaFabric type
type AbstractCriteriaFabric struct {
	mock.Mock
}

type AbstractCriteriaFabric_Expecter struct {
	mock *mock.Mock
}

func (_m *AbstractCriteriaFabric) EXPECT() *AbstractCriteriaFabric_Expecter {
	return &AbstractCriteriaFabric_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields:
func (_m *AbstractCriteriaFabric) Create() (criteria.Criteria, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 criteria.Criteria
	var r1 error
	if rf, ok := ret.Get(0).(func() (criteria.Criteria, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() criteria.Criteria); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(criteria.Criteria)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractCriteriaFabric_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type AbstractCriteriaFabric_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
func (_e *AbstractCriteriaFabric_Expecter) Create() *AbstractCriteriaFabric_Create_Call {
	return &AbstractCriteriaFabric_Create_Call{Call: _e.mock.On("Create")}
}

func (_c *AbstractCriteriaFabric_Create_Call) Run(run func()) *AbstractCriteriaFabric_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AbstractCriteriaFabric_Create_Call) Return(_a0 criteria.Criteria, _a1 error) *AbstractCriteriaFabric_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractCriteriaFabric_Create_Call) RunAndReturn(run func() (criteria.Criteria, error)) *AbstractCriteriaFabric_Create_Call {
	_c.Call.Return(run)
	return _c
}

// NewAbstractCriteriaFabric creates a new instance of AbstractCriteriaFabric. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAbstractCriteriaFabric(t interface {
	mock.TestingT
	Cleanup(func())
}) *AbstractCriteriaFabric {
	mock := &AbstractCriteriaFabric{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
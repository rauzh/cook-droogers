// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "cookdroogers/models"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// PublicationRepo is an autogenerated mock type for the PublicationRepo type
type PublicationRepo struct {
	mock.Mock
}

type PublicationRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *PublicationRepo) EXPECT() *PublicationRepo_Expecter {
	return &PublicationRepo_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *PublicationRepo) Create(_a0 context.Context, _a1 *models.Publication) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Publication) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublicationRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type PublicationRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *models.Publication
func (_e *PublicationRepo_Expecter) Create(_a0 interface{}, _a1 interface{}) *PublicationRepo_Create_Call {
	return &PublicationRepo_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *PublicationRepo_Create_Call) Run(run func(_a0 context.Context, _a1 *models.Publication)) *PublicationRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Publication))
	})
	return _c
}

func (_c *PublicationRepo_Create_Call) Return(_a0 error) *PublicationRepo_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PublicationRepo_Create_Call) RunAndReturn(run func(context.Context, *models.Publication) error) *PublicationRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *PublicationRepo) Get(_a0 context.Context, _a1 uint64) (*models.Publication, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.Publication
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*models.Publication, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Publication); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Publication)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicationRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type PublicationRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
func (_e *PublicationRepo_Expecter) Get(_a0 interface{}, _a1 interface{}) *PublicationRepo_Get_Call {
	return &PublicationRepo_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *PublicationRepo_Get_Call) Run(run func(_a0 context.Context, _a1 uint64)) *PublicationRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *PublicationRepo_Get_Call) Return(_a0 *models.Publication, _a1 error) *PublicationRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PublicationRepo_Get_Call) RunAndReturn(run func(context.Context, uint64) (*models.Publication, error)) *PublicationRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllByArtistSinceDate provides a mock function with given fields: ctx, date, artistID
func (_m *PublicationRepo) GetAllByArtistSinceDate(ctx context.Context, date time.Time, artistID uint64) ([]models.Publication, error) {
	ret := _m.Called(ctx, date, artistID)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByArtistSinceDate")
	}

	var r0 []models.Publication
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, uint64) ([]models.Publication, error)); ok {
		return rf(ctx, date, artistID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, uint64) []models.Publication); ok {
		r0 = rf(ctx, date, artistID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Publication)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, uint64) error); ok {
		r1 = rf(ctx, date, artistID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicationRepo_GetAllByArtistSinceDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByArtistSinceDate'
type PublicationRepo_GetAllByArtistSinceDate_Call struct {
	*mock.Call
}

// GetAllByArtistSinceDate is a helper method to define mock.On call
//   - ctx context.Context
//   - date time.Time
//   - artistID uint64
func (_e *PublicationRepo_Expecter) GetAllByArtistSinceDate(ctx interface{}, date interface{}, artistID interface{}) *PublicationRepo_GetAllByArtistSinceDate_Call {
	return &PublicationRepo_GetAllByArtistSinceDate_Call{Call: _e.mock.On("GetAllByArtistSinceDate", ctx, date, artistID)}
}

func (_c *PublicationRepo_GetAllByArtistSinceDate_Call) Run(run func(ctx context.Context, date time.Time, artistID uint64)) *PublicationRepo_GetAllByArtistSinceDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(time.Time), args[2].(uint64))
	})
	return _c
}

func (_c *PublicationRepo_GetAllByArtistSinceDate_Call) Return(_a0 []models.Publication, _a1 error) *PublicationRepo_GetAllByArtistSinceDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PublicationRepo_GetAllByArtistSinceDate_Call) RunAndReturn(run func(context.Context, time.Time, uint64) ([]models.Publication, error)) *PublicationRepo_GetAllByArtistSinceDate_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllByDate provides a mock function with given fields: _a0, _a1
func (_m *PublicationRepo) GetAllByDate(_a0 context.Context, _a1 time.Time) ([]models.Publication, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByDate")
	}

	var r0 []models.Publication
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) ([]models.Publication, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) []models.Publication); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Publication)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicationRepo_GetAllByDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByDate'
type PublicationRepo_GetAllByDate_Call struct {
	*mock.Call
}

// GetAllByDate is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 time.Time
func (_e *PublicationRepo_Expecter) GetAllByDate(_a0 interface{}, _a1 interface{}) *PublicationRepo_GetAllByDate_Call {
	return &PublicationRepo_GetAllByDate_Call{Call: _e.mock.On("GetAllByDate", _a0, _a1)}
}

func (_c *PublicationRepo_GetAllByDate_Call) Run(run func(_a0 context.Context, _a1 time.Time)) *PublicationRepo_GetAllByDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(time.Time))
	})
	return _c
}

func (_c *PublicationRepo_GetAllByDate_Call) Return(_a0 []models.Publication, _a1 error) *PublicationRepo_GetAllByDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PublicationRepo_GetAllByDate_Call) RunAndReturn(run func(context.Context, time.Time) ([]models.Publication, error)) *PublicationRepo_GetAllByDate_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllByManager provides a mock function with given fields: ctx, mng
func (_m *PublicationRepo) GetAllByManager(ctx context.Context, mng uint64) ([]models.Publication, error) {
	ret := _m.Called(ctx, mng)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByManager")
	}

	var r0 []models.Publication
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]models.Publication, error)); ok {
		return rf(ctx, mng)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []models.Publication); ok {
		r0 = rf(ctx, mng)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Publication)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, mng)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicationRepo_GetAllByManager_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllByManager'
type PublicationRepo_GetAllByManager_Call struct {
	*mock.Call
}

// GetAllByManager is a helper method to define mock.On call
//   - ctx context.Context
//   - mng uint64
func (_e *PublicationRepo_Expecter) GetAllByManager(ctx interface{}, mng interface{}) *PublicationRepo_GetAllByManager_Call {
	return &PublicationRepo_GetAllByManager_Call{Call: _e.mock.On("GetAllByManager", ctx, mng)}
}

func (_c *PublicationRepo_GetAllByManager_Call) Run(run func(ctx context.Context, mng uint64)) *PublicationRepo_GetAllByManager_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *PublicationRepo_GetAllByManager_Call) Return(_a0 []models.Publication, _a1 error) *PublicationRepo_GetAllByManager_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PublicationRepo_GetAllByManager_Call) RunAndReturn(run func(context.Context, uint64) ([]models.Publication, error)) *PublicationRepo_GetAllByManager_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *PublicationRepo) Update(_a0 context.Context, _a1 *models.Publication) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Publication) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublicationRepo_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type PublicationRepo_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *models.Publication
func (_e *PublicationRepo_Expecter) Update(_a0 interface{}, _a1 interface{}) *PublicationRepo_Update_Call {
	return &PublicationRepo_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *PublicationRepo_Update_Call) Run(run func(_a0 context.Context, _a1 *models.Publication)) *PublicationRepo_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Publication))
	})
	return _c
}

func (_c *PublicationRepo_Update_Call) Return(_a0 error) *PublicationRepo_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PublicationRepo_Update_Call) RunAndReturn(run func(context.Context, *models.Publication) error) *PublicationRepo_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewPublicationRepo creates a new instance of PublicationRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPublicationRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *PublicationRepo {
	mock := &PublicationRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

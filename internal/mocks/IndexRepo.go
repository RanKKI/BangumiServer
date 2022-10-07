// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bangumi/server/internal/domain"
	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"
)

// IndexRepo is an autogenerated mock type for the IndexRepo type
type IndexRepo struct {
	mock.Mock
}

type IndexRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *IndexRepo) EXPECT() *IndexRepo_Expecter {
	return &IndexRepo_Expecter{mock: &_m.Mock}
}

// AddIndexSubject provides a mock function with given fields: ctx, id, subject_id, sort, comment
func (_m *IndexRepo) AddIndexSubject(ctx context.Context, id uint32, subject_id model.SubjectID, sort uint32, comment string) (*domain.IndexSubject, error) {
	ret := _m.Called(ctx, id, subject_id, sort, comment)

	var r0 *domain.IndexSubject
	if rf, ok := ret.Get(0).(func(context.Context, uint32, model.SubjectID, uint32, string) *domain.IndexSubject); ok {
		r0 = rf(ctx, id, subject_id, sort, comment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.IndexSubject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, model.SubjectID, uint32, string) error); ok {
		r1 = rf(ctx, id, subject_id, sort, comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexRepo_AddIndexSubject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddIndexSubject'
type IndexRepo_AddIndexSubject_Call struct {
	*mock.Call
}

// AddIndexSubject is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
//   - subject_id model.SubjectID
//   - sort uint32
//   - comment string
func (_e *IndexRepo_Expecter) AddIndexSubject(ctx interface{}, id interface{}, subject_id interface{}, sort interface{}, comment interface{}) *IndexRepo_AddIndexSubject_Call {
	return &IndexRepo_AddIndexSubject_Call{Call: _e.mock.On("AddIndexSubject", ctx, id, subject_id, sort, comment)}
}

func (_c *IndexRepo_AddIndexSubject_Call) Run(run func(ctx context.Context, id uint32, subject_id model.SubjectID, sort uint32, comment string)) *IndexRepo_AddIndexSubject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(model.SubjectID), args[3].(uint32), args[4].(string))
	})
	return _c
}

func (_c *IndexRepo_AddIndexSubject_Call) Return(_a0 *domain.IndexSubject, _a1 error) *IndexRepo_AddIndexSubject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CountSubjects provides a mock function with given fields: ctx, id, subjectType
func (_m *IndexRepo) CountSubjects(ctx context.Context, id uint32, subjectType uint8) (int64, error) {
	ret := _m.Called(ctx, id, subjectType)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8) int64); ok {
		r0 = rf(ctx, id, subjectType)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint8) error); ok {
		r1 = rf(ctx, id, subjectType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexRepo_CountSubjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountSubjects'
type IndexRepo_CountSubjects_Call struct {
	*mock.Call
}

// CountSubjects is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
//   - subjectType uint8
func (_e *IndexRepo_Expecter) CountSubjects(ctx interface{}, id interface{}, subjectType interface{}) *IndexRepo_CountSubjects_Call {
	return &IndexRepo_CountSubjects_Call{Call: _e.mock.On("CountSubjects", ctx, id, subjectType)}
}

func (_c *IndexRepo_CountSubjects_Call) Run(run func(ctx context.Context, id uint32, subjectType uint8)) *IndexRepo_CountSubjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint8))
	})
	return _c
}

func (_c *IndexRepo_CountSubjects_Call) Return(_a0 int64, _a1 error) *IndexRepo_CountSubjects_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *IndexRepo) Delete(ctx context.Context, id uint32) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexRepo_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type IndexRepo_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
func (_e *IndexRepo_Expecter) Delete(ctx interface{}, id interface{}) *IndexRepo_Delete_Call {
	return &IndexRepo_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *IndexRepo_Delete_Call) Run(run func(ctx context.Context, id uint32)) *IndexRepo_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *IndexRepo_Delete_Call) Return(_a0 error) *IndexRepo_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// DeleteIndexSubject provides a mock function with given fields: ctx, id, subject_id
func (_m *IndexRepo) DeleteIndexSubject(ctx context.Context, id uint32, subject_id model.SubjectID) error {
	ret := _m.Called(ctx, id, subject_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, model.SubjectID) error); ok {
		r0 = rf(ctx, id, subject_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexRepo_DeleteIndexSubject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteIndexSubject'
type IndexRepo_DeleteIndexSubject_Call struct {
	*mock.Call
}

// DeleteIndexSubject is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
//   - subject_id model.SubjectID
func (_e *IndexRepo_Expecter) DeleteIndexSubject(ctx interface{}, id interface{}, subject_id interface{}) *IndexRepo_DeleteIndexSubject_Call {
	return &IndexRepo_DeleteIndexSubject_Call{Call: _e.mock.On("DeleteIndexSubject", ctx, id, subject_id)}
}

func (_c *IndexRepo_DeleteIndexSubject_Call) Run(run func(ctx context.Context, id uint32, subject_id model.SubjectID)) *IndexRepo_DeleteIndexSubject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(model.SubjectID))
	})
	return _c
}

func (_c *IndexRepo_DeleteIndexSubject_Call) Return(_a0 error) *IndexRepo_DeleteIndexSubject_Call {
	_c.Call.Return(_a0)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *IndexRepo) Get(ctx context.Context, id uint32) (model.Index, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Index
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.Index); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Index)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type IndexRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
func (_e *IndexRepo_Expecter) Get(ctx interface{}, id interface{}) *IndexRepo_Get_Call {
	return &IndexRepo_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *IndexRepo_Get_Call) Run(run func(ctx context.Context, id uint32)) *IndexRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *IndexRepo_Get_Call) Return(_a0 model.Index, _a1 error) *IndexRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListSubjects provides a mock function with given fields: ctx, id, subjectType, limit, offset
func (_m *IndexRepo) ListSubjects(ctx context.Context, id uint32, subjectType uint8, limit int, offset int) ([]domain.IndexSubject, error) {
	ret := _m.Called(ctx, id, subjectType, limit, offset)

	var r0 []domain.IndexSubject
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8, int, int) []domain.IndexSubject); ok {
		r0 = rf(ctx, id, subjectType, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.IndexSubject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint8, int, int) error); ok {
		r1 = rf(ctx, id, subjectType, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexRepo_ListSubjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSubjects'
type IndexRepo_ListSubjects_Call struct {
	*mock.Call
}

// ListSubjects is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
//   - subjectType uint8
//   - limit int
//   - offset int
func (_e *IndexRepo_Expecter) ListSubjects(ctx interface{}, id interface{}, subjectType interface{}, limit interface{}, offset interface{}) *IndexRepo_ListSubjects_Call {
	return &IndexRepo_ListSubjects_Call{Call: _e.mock.On("ListSubjects", ctx, id, subjectType, limit, offset)}
}

func (_c *IndexRepo_ListSubjects_Call) Run(run func(ctx context.Context, id uint32, subjectType uint8, limit int, offset int)) *IndexRepo_ListSubjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint8), args[3].(int), args[4].(int))
	})
	return _c
}

func (_c *IndexRepo_ListSubjects_Call) Return(_a0 []domain.IndexSubject, _a1 error) *IndexRepo_ListSubjects_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// New provides a mock function with given fields: ctx, i
func (_m *IndexRepo) New(ctx context.Context, i *model.Index) error {
	ret := _m.Called(ctx, i)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Index) error); ok {
		r0 = rf(ctx, i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexRepo_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type IndexRepo_New_Call struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - ctx context.Context
//   - i *model.Index
func (_e *IndexRepo_Expecter) New(ctx interface{}, i interface{}) *IndexRepo_New_Call {
	return &IndexRepo_New_Call{Call: _e.mock.On("New", ctx, i)}
}

func (_c *IndexRepo_New_Call) Run(run func(ctx context.Context, i *model.Index)) *IndexRepo_New_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Index))
	})
	return _c
}

func (_c *IndexRepo_New_Call) Return(_a0 error) *IndexRepo_New_Call {
	_c.Call.Return(_a0)
	return _c
}

// Update provides a mock function with given fields: ctx, id, title, desc
func (_m *IndexRepo) Update(ctx context.Context, id uint32, title string, desc string) error {
	ret := _m.Called(ctx, id, title, desc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, string, string) error); ok {
		r0 = rf(ctx, id, title, desc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexRepo_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type IndexRepo_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
//   - title string
//   - desc string
func (_e *IndexRepo_Expecter) Update(ctx interface{}, id interface{}, title interface{}, desc interface{}) *IndexRepo_Update_Call {
	return &IndexRepo_Update_Call{Call: _e.mock.On("Update", ctx, id, title, desc)}
}

func (_c *IndexRepo_Update_Call) Run(run func(ctx context.Context, id uint32, title string, desc string)) *IndexRepo_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *IndexRepo_Update_Call) Return(_a0 error) *IndexRepo_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

// UpdateIndexSubject provides a mock function with given fields: ctx, id, subject_id, sort, comment
func (_m *IndexRepo) UpdateIndexSubject(ctx context.Context, id uint32, subject_id model.SubjectID, sort uint32, comment string) error {
	ret := _m.Called(ctx, id, subject_id, sort, comment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, model.SubjectID, uint32, string) error); ok {
		r0 = rf(ctx, id, subject_id, sort, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexRepo_UpdateIndexSubject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateIndexSubject'
type IndexRepo_UpdateIndexSubject_Call struct {
	*mock.Call
}

// UpdateIndexSubject is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
//   - subject_id model.SubjectID
//   - sort uint32
//   - comment string
func (_e *IndexRepo_Expecter) UpdateIndexSubject(ctx interface{}, id interface{}, subject_id interface{}, sort interface{}, comment interface{}) *IndexRepo_UpdateIndexSubject_Call {
	return &IndexRepo_UpdateIndexSubject_Call{Call: _e.mock.On("UpdateIndexSubject", ctx, id, subject_id, sort, comment)}
}

func (_c *IndexRepo_UpdateIndexSubject_Call) Run(run func(ctx context.Context, id uint32, subject_id model.SubjectID, sort uint32, comment string)) *IndexRepo_UpdateIndexSubject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(model.SubjectID), args[3].(uint32), args[4].(string))
	})
	return _c
}

func (_c *IndexRepo_UpdateIndexSubject_Call) Return(_a0 error) *IndexRepo_UpdateIndexSubject_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewIndexRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewIndexRepo creates a new instance of IndexRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIndexRepo(t mockConstructorTestingTNewIndexRepo) *IndexRepo {
	mock := &IndexRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

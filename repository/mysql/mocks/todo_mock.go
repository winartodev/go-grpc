// Code generated by mockery v2.32.4. DO NOT EDIT.

package todorepositorymock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/winartodev/go-grpc/types"
)

// TodoRepositoryInterface is an autogenerated mock type for the TodoRepositoryInterface type
type TodoRepositoryInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data
func (_m *TodoRepositoryInterface) Create(ctx context.Context, data types.Task) (int64, error) {
	ret := _m.Called(ctx, data)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Task) (int64, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.Task) int64); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.Task) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByIDDB provides a mock function with given fields: ctx, id
func (_m *TodoRepositoryInterface) DeleteByIDDB(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTaskDB provides a mock function with given fields: ctx
func (_m *TodoRepositoryInterface) GetAllTaskDB(ctx context.Context) ([]types.Task, error) {
	ret := _m.Called(ctx)

	var r0 []types.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]types.Task, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []types.Task); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TodoRepositoryInterface) GetByID(ctx context.Context, id int64) (*types.Task, error) {
	ret := _m.Called(ctx, id)

	var r0 *types.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*types.Task, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *types.Task); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByIDDB provides a mock function with given fields: ctx, id, data
func (_m *TodoRepositoryInterface) UpdateByIDDB(ctx context.Context, id int64, data types.Task) error {
	ret := _m.Called(ctx, id, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, types.Task) error); ok {
		r0 = rf(ctx, id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTodoRepositoryInterface creates a new instance of TodoRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTodoRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *TodoRepositoryInterface {
	mock := &TodoRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
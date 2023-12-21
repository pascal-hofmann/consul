// Code generated by mockery v2.37.1. DO NOT EDIT.

package indexmock

import (
	pbresource "github.com/hashicorp/consul/proto-public/pbresource"
	mock "github.com/stretchr/testify/mock"
)

// MultiIndexer is an autogenerated mock type for the MultiIndexer type
type MultiIndexer struct {
	mock.Mock
}

type MultiIndexer_Expecter struct {
	mock *mock.Mock
}

func (_m *MultiIndexer) EXPECT() *MultiIndexer_Expecter {
	return &MultiIndexer_Expecter{mock: &_m.Mock}
}

// FromArgs provides a mock function with given fields: args
func (_m *MultiIndexer) FromArgs(args ...interface{}) ([]byte, error) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(...interface{}) ([]byte, error)); ok {
		return rf(args...)
	}
	if rf, ok := ret.Get(0).(func(...interface{}) []byte); ok {
		r0 = rf(args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MultiIndexer_FromArgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FromArgs'
type MultiIndexer_FromArgs_Call struct {
	*mock.Call
}

// FromArgs is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MultiIndexer_Expecter) FromArgs(args ...interface{}) *MultiIndexer_FromArgs_Call {
	return &MultiIndexer_FromArgs_Call{Call: _e.mock.On("FromArgs",
		append([]interface{}{}, args...)...)}
}

func (_c *MultiIndexer_FromArgs_Call) Run(run func(args ...interface{})) *MultiIndexer_FromArgs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MultiIndexer_FromArgs_Call) Return(_a0 []byte, _a1 error) *MultiIndexer_FromArgs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MultiIndexer_FromArgs_Call) RunAndReturn(run func(...interface{}) ([]byte, error)) *MultiIndexer_FromArgs_Call {
	_c.Call.Return(run)
	return _c
}

// FromResource provides a mock function with given fields: r
func (_m *MultiIndexer) FromResource(r *pbresource.Resource) (bool, [][]byte, error) {
	ret := _m.Called(r)

	var r0 bool
	var r1 [][]byte
	var r2 error
	if rf, ok := ret.Get(0).(func(*pbresource.Resource) (bool, [][]byte, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(*pbresource.Resource) bool); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*pbresource.Resource) [][]byte); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]byte)
		}
	}

	if rf, ok := ret.Get(2).(func(*pbresource.Resource) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MultiIndexer_FromResource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FromResource'
type MultiIndexer_FromResource_Call struct {
	*mock.Call
}

// FromResource is a helper method to define mock.On call
//   - r *pbresource.Resource
func (_e *MultiIndexer_Expecter) FromResource(r interface{}) *MultiIndexer_FromResource_Call {
	return &MultiIndexer_FromResource_Call{Call: _e.mock.On("FromResource", r)}
}

func (_c *MultiIndexer_FromResource_Call) Run(run func(r *pbresource.Resource)) *MultiIndexer_FromResource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*pbresource.Resource))
	})
	return _c
}

func (_c *MultiIndexer_FromResource_Call) Return(_a0 bool, _a1 [][]byte, _a2 error) *MultiIndexer_FromResource_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MultiIndexer_FromResource_Call) RunAndReturn(run func(*pbresource.Resource) (bool, [][]byte, error)) *MultiIndexer_FromResource_Call {
	_c.Call.Return(run)
	return _c
}

// NewMultiIndexer creates a new instance of MultiIndexer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMultiIndexer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MultiIndexer {
	mock := &MultiIndexer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

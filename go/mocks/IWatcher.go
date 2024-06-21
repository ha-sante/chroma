// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	memberlist_manager "github.com/chroma-core/chroma/go/pkg/memberlist_manager"
	mock "github.com/stretchr/testify/mock"
)

// IWatcher is an autogenerated mock type for the IWatcher type
type IWatcher struct {
	mock.Mock
}

// ListReadyMembers provides a mock function with given fields:
func (_m *IWatcher) ListReadyMembers() (memberlist_manager.Memberlist, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListReadyMembers")
	}

	var r0 memberlist_manager.Memberlist
	var r1 error
	if rf, ok := ret.Get(0).(func() (memberlist_manager.Memberlist, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() memberlist_manager.Memberlist); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(memberlist_manager.Memberlist)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterCallback provides a mock function with given fields: callback
func (_m *IWatcher) RegisterCallback(callback memberlist_manager.NodeWatcherCallback) {
	_m.Called(callback)
}

// Start provides a mock function with given fields:
func (_m *IWatcher) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *IWatcher) Stop() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIWatcher creates a new instance of IWatcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIWatcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *IWatcher {
	mock := &IWatcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

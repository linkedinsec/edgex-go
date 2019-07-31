// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SchedulerQueueDeleter is an autogenerated mock type for the SchedulerQueueDeleter type
type SchedulerQueueDeleter struct {
	mock.Mock
}

// RemoveIntervalInQueue provides a mock function with given fields: intervalId
func (_m *SchedulerQueueDeleter) RemoveIntervalInQueue(intervalId string) error {
	ret := _m.Called(intervalId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(intervalId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
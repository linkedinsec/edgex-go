// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ServiceStopper is an autogenerated mock type for the ServiceStopper type
type ServiceStopper struct {
	mock.Mock
}

// Stop provides a mock function with given fields: service
func (_m *ServiceStopper) Stop(service string) error {
	ret := _m.Called(service)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(service)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
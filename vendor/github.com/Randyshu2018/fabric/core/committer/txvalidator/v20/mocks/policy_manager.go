// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	policies "github.com/Randyshu2018/fabric/common/policies"
	mock "github.com/stretchr/testify/mock"
)

// PolicyManager is an autogenerated mock type for the PolicyManager type
type PolicyManager struct {
	mock.Mock
}

// GetPolicy provides a mock function with given fields: id
func (_m *PolicyManager) GetPolicy(id string) (policies.Policy, bool) {
	ret := _m.Called(id)

	var r0 policies.Policy
	if rf, ok := ret.Get(0).(func(string) policies.Policy); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(policies.Policy)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Manager provides a mock function with given fields: path
func (_m *PolicyManager) Manager(path []string) (policies.Manager, bool) {
	ret := _m.Called(path)

	var r0 policies.Manager
	if rf, ok := ret.Get(0).(func([]string) policies.Manager); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(policies.Manager)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func([]string) bool); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

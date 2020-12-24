// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	cluster "github.com/Randyshu2018/fabric/orderer/common/cluster"
	mock "github.com/stretchr/testify/mock"
)

// LedgerFactory is an autogenerated mock type for the LedgerFactory type
type LedgerFactory struct {
	mock.Mock
}

// GetOrCreate provides a mock function with given fields: chainID
func (_m *LedgerFactory) GetOrCreate(chainID string) (cluster.LedgerWriter, error) {
	ret := _m.Called(chainID)

	var r0 cluster.LedgerWriter
	if rf, ok := ret.Get(0).(func(string) cluster.LedgerWriter); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cluster.LedgerWriter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

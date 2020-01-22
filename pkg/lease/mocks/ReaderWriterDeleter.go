// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import lease "github.com/Optum/dce/pkg/lease"
import mock "github.com/stretchr/testify/mock"

// ReaderWriterDeleter is an autogenerated mock type for the ReaderWriterDeleter type
type ReaderWriterDeleter struct {
	mock.Mock
}

// Delete provides a mock function with given fields: input
func (_m *ReaderWriterDeleter) Delete(input *lease.Lease) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(*lease.Lease) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: leaseID
func (_m *ReaderWriterDeleter) Get(leaseID string) (*lease.Lease, error) {
	ret := _m.Called(leaseID)

	var r0 *lease.Lease
	if rf, ok := ret.Get(0).(func(string) *lease.Lease); ok {
		r0 = rf(leaseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lease.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(leaseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0
func (_m *ReaderWriterDeleter) List(_a0 *lease.Lease) (*lease.Leases, error) {
	ret := _m.Called(_a0)

	var r0 *lease.Leases
	if rf, ok := ret.Get(0).(func(*lease.Lease) *lease.Leases); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lease.Leases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*lease.Lease) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Write provides a mock function with given fields: input, lastModifiedOn
func (_m *ReaderWriterDeleter) Write(input *lease.Lease, lastModifiedOn *int64) error {
	ret := _m.Called(input, lastModifiedOn)

	var r0 error
	if rf, ok := ret.Get(0).(func(*lease.Lease, *int64) error); ok {
		r0 = rf(input, lastModifiedOn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
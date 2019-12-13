// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import account "github.com/Optum/dce/pkg/account"
import mock "github.com/stretchr/testify/mock"
import model "github.com/Optum/dce/pkg/model"

// Accounter is an autogenerated mock type for the Accounter type
type Accounter struct {
	mock.Mock
}

// Update provides a mock function with given fields: d, u, am
func (_m *Accounter) Update(d model.Account, u account.Updater, am account.Manager) error {
	ret := _m.Called(d, u, am)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Account, account.Updater, account.Manager) error); ok {
		r0 = rf(d, u, am)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
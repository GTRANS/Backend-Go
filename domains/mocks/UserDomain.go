// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "WallE/models"

	mock "github.com/stretchr/testify/mock"
)

// UserDomain is an autogenerated mock type for the UserDomain type
type UserDomain struct {
	mock.Mock
}

// CreateResetPassword provides a mock function with given fields: reset
func (_m *UserDomain) CreateResetPassword(reset models.ResetPassword) error {
	ret := _m.Called(reset)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.ResetPassword) error); ok {
		r0 = rf(reset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserDomain) GetByEmail(email string) (models.User, error) {
	ret := _m.Called(email)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResetPassword provides a mock function with given fields: email
func (_m *UserDomain) GetResetPassword(email string) (models.ResetPassword, error) {
	ret := _m.Called(email)

	var r0 models.ResetPassword
	if rf, ok := ret.Get(0).(func(string) models.ResetPassword); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(models.ResetPassword)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *UserDomain) GetUserByEmail(email string) (models.User, error) {
	ret := _m.Called(email)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserDataById provides a mock function with given fields: id
func (_m *UserDomain) GetUserDataById(id uint) (models.User, error) {
	ret := _m.Called(id)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(uint) models.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: user
func (_m *UserDomain) Register(user models.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePassword provides a mock function with given fields: email, password
func (_m *UserDomain) UpdatePassword(email string, password string) error {
	ret := _m.Called(email, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateResetTable provides a mock function with given fields: email
func (_m *UserDomain) UpdateResetTable(email string) error {
	ret := _m.Called(email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserData provides a mock function with given fields: id, user
func (_m *UserDomain) UpdateUserData(id uint, user models.User) error {
	ret := _m.Called(id, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, models.User) error); ok {
		r0 = rf(id, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Verifikasi provides a mock function with given fields: id
func (_m *UserDomain) Verifikasi(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserDomain interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserDomain creates a new instance of UserDomain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserDomain(t mockConstructorTestingTNewUserDomain) *UserDomain {
	mock := &UserDomain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
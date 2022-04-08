// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	users "Daily-Calorie-App-API/business/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) (string, error) {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUser provides a mock function with given fields:
func (_m *Repository) GetAllUser() (*[]users.Domain, error) {
	ret := _m.Called()

	var r0 *[]users.Domain
	if rf, ok := ret.Get(0).(func() *[]users.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]users.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *Repository) GetUserByEmail(email string) (*users.Domain, error) {
	ret := _m.Called(email)

	var r0 *users.Domain
	if rf, ok := ret.Get(0).(func(string) *users.Domain); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: id
func (_m *Repository) GetUserByID(id int) (*users.Domain, error) {
	ret := _m.Called(id)

	var r0 *users.Domain
	if rf, ok := ret.Get(0).(func(int) *users.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: user
func (_m *Repository) Insert(user *users.Domain) (*users.Domain, error) {
	ret := _m.Called(user)

	var r0 *users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) *users.Domain); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*users.Domain) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, user
func (_m *Repository) Update(id int, user *users.Domain) (*users.Domain, error) {
	ret := _m.Called(id, user)

	var r0 *users.Domain
	if rf, ok := ret.Get(0).(func(int, *users.Domain) *users.Domain); ok {
		r0 = rf(id, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *users.Domain) error); ok {
		r1 = rf(id, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

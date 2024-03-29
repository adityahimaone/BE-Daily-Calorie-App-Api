// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	admins "Daily-Calorie-App-API/business/admins"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetAdminByUsername provides a mock function with given fields: username
func (_m *Repository) GetAdminByUsername(username string) (*admins.Domain, error) {
	ret := _m.Called(username)

	var r0 *admins.Domain
	if rf, ok := ret.Get(0).(func(string) *admins.Domain); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admins.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: domain
func (_m *Repository) Insert(domain *admins.Domain) (*admins.Domain, error) {
	ret := _m.Called(domain)

	var r0 *admins.Domain
	if rf, ok := ret.Get(0).(func(*admins.Domain) *admins.Domain); ok {
		r0 = rf(domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admins.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*admins.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

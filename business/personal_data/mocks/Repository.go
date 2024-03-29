// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	personal_data "Daily-Calorie-App-API/business/personal_data"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: personaldata
func (_m *Repository) Insert(personaldata *personal_data.Domain) (*personal_data.Domain, error) {
	ret := _m.Called(personaldata)

	var r0 *personal_data.Domain
	if rf, ok := ret.Get(0).(func(*personal_data.Domain) *personal_data.Domain); ok {
		r0 = rf(personaldata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*personal_data.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*personal_data.Domain) error); ok {
		r1 = rf(personaldata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, personaldata
func (_m *Repository) Update(id int, personaldata *personal_data.Domain) (*personal_data.Domain, error) {
	ret := _m.Called(id, personaldata)

	var r0 *personal_data.Domain
	if rf, ok := ret.Get(0).(func(int, *personal_data.Domain) *personal_data.Domain); ok {
		r0 = rf(id, personaldata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*personal_data.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *personal_data.Domain) error); ok {
		r1 = rf(id, personaldata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

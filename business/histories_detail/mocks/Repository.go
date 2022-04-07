// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	histories_detail "Daily-Calorie-App-API/business/histories_detail"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) (*histories_detail.Domain, error) {
	ret := _m.Called(id)

	var r0 *histories_detail.Domain
	if rf, ok := ret.Get(0).(func(int) *histories_detail.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*histories_detail.Domain)
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

// GetAllHistoriesDetailByHistoriesID provides a mock function with given fields: historiesID
func (_m *Repository) GetAllHistoriesDetailByHistoriesID(historiesID int) (*[]histories_detail.Domain, error) {
	ret := _m.Called(historiesID)

	var r0 *[]histories_detail.Domain
	if rf, ok := ret.Get(0).(func(int) *[]histories_detail.Domain); ok {
		r0 = rf(historiesID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]histories_detail.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(historiesID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: historiesDetail
func (_m *Repository) Insert(historiesDetail *histories_detail.Domain) (*histories_detail.Domain, error) {
	ret := _m.Called(historiesDetail)

	var r0 *histories_detail.Domain
	if rf, ok := ret.Get(0).(func(*histories_detail.Domain) *histories_detail.Domain); ok {
		r0 = rf(historiesDetail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*histories_detail.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*histories_detail.Domain) error); ok {
		r1 = rf(historiesDetail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SumCalories provides a mock function with given fields: historiesID
func (_m *Repository) SumCalories(historiesID int) (float64, error) {
	ret := _m.Called(historiesID)

	var r0 float64
	if rf, ok := ret.Get(0).(func(int) float64); ok {
		r0 = rf(historiesID)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(historiesID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
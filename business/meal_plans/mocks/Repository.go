// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	meal_plans "Daily-Calorie-App-API/business/meal_plans"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetLastMealPlans provides a mock function with given fields: userID
func (_m *Repository) GetLastMealPlans(userID int) (*meal_plans.Domain, error) {
	ret := _m.Called(userID)

	var r0 *meal_plans.Domain
	if rf, ok := ret.Get(0).(func(int) *meal_plans.Domain); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*meal_plans.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: domain
func (_m *Repository) Insert(domain *meal_plans.Domain) (*meal_plans.Domain, error) {
	ret := _m.Called(domain)

	var r0 *meal_plans.Domain
	if rf, ok := ret.Get(0).(func(*meal_plans.Domain) *meal_plans.Domain); ok {
		r0 = rf(domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*meal_plans.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*meal_plans.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

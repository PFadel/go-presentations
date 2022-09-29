// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Calculator is an autogenerated mock type for the Calculator type
type Calculator struct {
	mock.Mock
}

type Calculator_Expecter struct {
	mock *mock.Mock
}

func (_m *Calculator) EXPECT() *Calculator_Expecter {
	return &Calculator_Expecter{mock: &_m.Mock}
}

// Distance provides a mock function with given fields: lat1, lng1, lat2, lng2, unit
func (_m *Calculator) Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	_va := make([]interface{}, len(unit))
	for _i := range unit {
		_va[_i] = unit[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, lat1, lng1, lat2, lng2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 float64
	if rf, ok := ret.Get(0).(func(float64, float64, float64, float64, ...string) float64); ok {
		r0 = rf(lat1, lng1, lat2, lng2, unit...)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Calculator_Distance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Distance'
type Calculator_Distance_Call struct {
	*mock.Call
}

// Distance is a helper method to define mock.On call
//   - lat1 float64
//   - lng1 float64
//   - lat2 float64
//   - lng2 float64
//   - unit ...string
func (_e *Calculator_Expecter) Distance(lat1 interface{}, lng1 interface{}, lat2 interface{}, lng2 interface{}, unit ...interface{}) *Calculator_Distance_Call {
	return &Calculator_Distance_Call{Call: _e.mock.On("Distance",
		append([]interface{}{lat1, lng1, lat2, lng2}, unit...)...)}
}

func (_c *Calculator_Distance_Call) Run(run func(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string)) *Calculator_Distance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-4)
		for i, a := range args[4:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(float64), args[1].(float64), args[2].(float64), args[3].(float64), variadicArgs...)
	})
	return _c
}

func (_c *Calculator_Distance_Call) Return(_a0 float64) *Calculator_Distance_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewCalculator interface {
	mock.TestingT
	Cleanup(func())
}

// NewCalculator creates a new instance of Calculator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCalculator(t mockConstructorTestingTNewCalculator) *Calculator {
	mock := &Calculator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

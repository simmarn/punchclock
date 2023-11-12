// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PreferencesWrapper is an autogenerated mock type for the PreferencesWrapper type
type PreferencesWrapper struct {
	mock.Mock
}

type PreferencesWrapper_Expecter struct {
	mock *mock.Mock
}

func (_m *PreferencesWrapper) EXPECT() *PreferencesWrapper_Expecter {
	return &PreferencesWrapper_Expecter{mock: &_m.Mock}
}

// GetBool provides a mock function with given fields: _a0
func (_m *PreferencesWrapper) GetBool(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PreferencesWrapper_GetBool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBool'
type PreferencesWrapper_GetBool_Call struct {
	*mock.Call
}

// GetBool is a helper method to define mock.On call
//   - _a0 string
func (_e *PreferencesWrapper_Expecter) GetBool(_a0 interface{}) *PreferencesWrapper_GetBool_Call {
	return &PreferencesWrapper_GetBool_Call{Call: _e.mock.On("GetBool", _a0)}
}

func (_c *PreferencesWrapper_GetBool_Call) Run(run func(_a0 string)) *PreferencesWrapper_GetBool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *PreferencesWrapper_GetBool_Call) Return(_a0 bool) *PreferencesWrapper_GetBool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PreferencesWrapper_GetBool_Call) RunAndReturn(run func(string) bool) *PreferencesWrapper_GetBool_Call {
	_c.Call.Return(run)
	return _c
}

// GetFloatWithFallback provides a mock function with given fields: _a0, _a1
func (_m *PreferencesWrapper) GetFloatWithFallback(_a0 string, _a1 float64) float64 {
	ret := _m.Called(_a0, _a1)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string, float64) float64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// PreferencesWrapper_GetFloatWithFallback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFloatWithFallback'
type PreferencesWrapper_GetFloatWithFallback_Call struct {
	*mock.Call
}

// GetFloatWithFallback is a helper method to define mock.On call
//   - _a0 string
//   - _a1 float64
func (_e *PreferencesWrapper_Expecter) GetFloatWithFallback(_a0 interface{}, _a1 interface{}) *PreferencesWrapper_GetFloatWithFallback_Call {
	return &PreferencesWrapper_GetFloatWithFallback_Call{Call: _e.mock.On("GetFloatWithFallback", _a0, _a1)}
}

func (_c *PreferencesWrapper_GetFloatWithFallback_Call) Run(run func(_a0 string, _a1 float64)) *PreferencesWrapper_GetFloatWithFallback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(float64))
	})
	return _c
}

func (_c *PreferencesWrapper_GetFloatWithFallback_Call) Return(_a0 float64) *PreferencesWrapper_GetFloatWithFallback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PreferencesWrapper_GetFloatWithFallback_Call) RunAndReturn(run func(string, float64) float64) *PreferencesWrapper_GetFloatWithFallback_Call {
	_c.Call.Return(run)
	return _c
}

// GetString provides a mock function with given fields: _a0
func (_m *PreferencesWrapper) GetString(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PreferencesWrapper_GetString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetString'
type PreferencesWrapper_GetString_Call struct {
	*mock.Call
}

// GetString is a helper method to define mock.On call
//   - _a0 string
func (_e *PreferencesWrapper_Expecter) GetString(_a0 interface{}) *PreferencesWrapper_GetString_Call {
	return &PreferencesWrapper_GetString_Call{Call: _e.mock.On("GetString", _a0)}
}

func (_c *PreferencesWrapper_GetString_Call) Run(run func(_a0 string)) *PreferencesWrapper_GetString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *PreferencesWrapper_GetString_Call) Return(_a0 string) *PreferencesWrapper_GetString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PreferencesWrapper_GetString_Call) RunAndReturn(run func(string) string) *PreferencesWrapper_GetString_Call {
	_c.Call.Return(run)
	return _c
}

// SetBool provides a mock function with given fields: _a0, _a1
func (_m *PreferencesWrapper) SetBool(_a0 string, _a1 bool) {
	_m.Called(_a0, _a1)
}

// PreferencesWrapper_SetBool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetBool'
type PreferencesWrapper_SetBool_Call struct {
	*mock.Call
}

// SetBool is a helper method to define mock.On call
//   - _a0 string
//   - _a1 bool
func (_e *PreferencesWrapper_Expecter) SetBool(_a0 interface{}, _a1 interface{}) *PreferencesWrapper_SetBool_Call {
	return &PreferencesWrapper_SetBool_Call{Call: _e.mock.On("SetBool", _a0, _a1)}
}

func (_c *PreferencesWrapper_SetBool_Call) Run(run func(_a0 string, _a1 bool)) *PreferencesWrapper_SetBool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *PreferencesWrapper_SetBool_Call) Return() *PreferencesWrapper_SetBool_Call {
	_c.Call.Return()
	return _c
}

func (_c *PreferencesWrapper_SetBool_Call) RunAndReturn(run func(string, bool)) *PreferencesWrapper_SetBool_Call {
	_c.Call.Return(run)
	return _c
}

// SetFloat provides a mock function with given fields: _a0, _a1
func (_m *PreferencesWrapper) SetFloat(_a0 string, _a1 float64) {
	_m.Called(_a0, _a1)
}

// PreferencesWrapper_SetFloat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFloat'
type PreferencesWrapper_SetFloat_Call struct {
	*mock.Call
}

// SetFloat is a helper method to define mock.On call
//   - _a0 string
//   - _a1 float64
func (_e *PreferencesWrapper_Expecter) SetFloat(_a0 interface{}, _a1 interface{}) *PreferencesWrapper_SetFloat_Call {
	return &PreferencesWrapper_SetFloat_Call{Call: _e.mock.On("SetFloat", _a0, _a1)}
}

func (_c *PreferencesWrapper_SetFloat_Call) Run(run func(_a0 string, _a1 float64)) *PreferencesWrapper_SetFloat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(float64))
	})
	return _c
}

func (_c *PreferencesWrapper_SetFloat_Call) Return() *PreferencesWrapper_SetFloat_Call {
	_c.Call.Return()
	return _c
}

func (_c *PreferencesWrapper_SetFloat_Call) RunAndReturn(run func(string, float64)) *PreferencesWrapper_SetFloat_Call {
	_c.Call.Return(run)
	return _c
}

// SetString provides a mock function with given fields: _a0, _a1
func (_m *PreferencesWrapper) SetString(_a0 string, _a1 string) {
	_m.Called(_a0, _a1)
}

// PreferencesWrapper_SetString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetString'
type PreferencesWrapper_SetString_Call struct {
	*mock.Call
}

// SetString is a helper method to define mock.On call
//   - _a0 string
//   - _a1 string
func (_e *PreferencesWrapper_Expecter) SetString(_a0 interface{}, _a1 interface{}) *PreferencesWrapper_SetString_Call {
	return &PreferencesWrapper_SetString_Call{Call: _e.mock.On("SetString", _a0, _a1)}
}

func (_c *PreferencesWrapper_SetString_Call) Run(run func(_a0 string, _a1 string)) *PreferencesWrapper_SetString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *PreferencesWrapper_SetString_Call) Return() *PreferencesWrapper_SetString_Call {
	_c.Call.Return()
	return _c
}

func (_c *PreferencesWrapper_SetString_Call) RunAndReturn(run func(string, string)) *PreferencesWrapper_SetString_Call {
	_c.Call.Return(run)
	return _c
}

// NewPreferencesWrapper creates a new instance of PreferencesWrapper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPreferencesWrapper(t interface {
	mock.TestingT
	Cleanup(func())
}) *PreferencesWrapper {
	mock := &PreferencesWrapper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package sentryhook

import (
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	mock "github.com/stretchr/testify/mock"
)

// NewMockISentryHook creates a new instance of MockISentryHook. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockISentryHook(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockISentryHook {
	mock := &MockISentryHook{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockISentryHook is an autogenerated mock type for the ISentryHook type
type MockISentryHook struct {
	mock.Mock
}

type MockISentryHook_Expecter struct {
	mock *mock.Mock
}

func (_m *MockISentryHook) EXPECT() *MockISentryHook_Expecter {
	return &MockISentryHook_Expecter{mock: &_m.Mock}
}

// Run provides a mock function for the type MockISentryHook
func (_mock *MockISentryHook) Run(event *zerolog.Event, level sentry.Level, msg string) {
	_mock.Called(event, level, msg)
	return
}

// MockISentryHook_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockISentryHook_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - event *zerolog.Event
//   - level sentry.Level
//   - msg string
func (_e *MockISentryHook_Expecter) Run(event interface{}, level interface{}, msg interface{}) *MockISentryHook_Run_Call {
	return &MockISentryHook_Run_Call{Call: _e.mock.On("Run", event, level, msg)}
}

func (_c *MockISentryHook_Run_Call) Run(run func(event *zerolog.Event, level sentry.Level, msg string)) *MockISentryHook_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 *zerolog.Event
		if args[0] != nil {
			arg0 = args[0].(*zerolog.Event)
		}
		var arg1 sentry.Level
		if args[1] != nil {
			arg1 = args[1].(sentry.Level)
		}
		var arg2 string
		if args[2] != nil {
			arg2 = args[2].(string)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *MockISentryHook_Run_Call) Return() *MockISentryHook_Run_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockISentryHook_Run_Call) RunAndReturn(run func(event *zerolog.Event, level sentry.Level, msg string)) *MockISentryHook_Run_Call {
	_c.Run(run)
	return _c
}

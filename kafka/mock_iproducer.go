// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package kafka

import (
	"context"

	mock "github.com/stretchr/testify/mock"
)

// NewMockIProducer creates a new instance of MockIProducer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIProducer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIProducer {
	mock := &MockIProducer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockIProducer is an autogenerated mock type for the IProducer type
type MockIProducer struct {
	mock.Mock
}

type MockIProducer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIProducer) EXPECT() *MockIProducer_Expecter {
	return &MockIProducer_Expecter{mock: &_m.Mock}
}

// Close provides a mock function for the type MockIProducer
func (_mock *MockIProducer) Close() error {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func() error); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockIProducer_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockIProducer_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockIProducer_Expecter) Close() *MockIProducer_Close_Call {
	return &MockIProducer_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockIProducer_Close_Call) Run(run func()) *MockIProducer_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIProducer_Close_Call) Return(err error) *MockIProducer_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockIProducer_Close_Call) RunAndReturn(run func() error) *MockIProducer_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function for the type MockIProducer
func (_mock *MockIProducer) Send(ctx context.Context, key []byte, value []byte, headers map[string]string) (int32, int64, error) {
	ret := _mock.Called(ctx, key, value, headers)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 int32
	var r1 int64
	var r2 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, []byte, []byte, map[string]string) (int32, int64, error)); ok {
		return returnFunc(ctx, key, value, headers)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, []byte, []byte, map[string]string) int32); ok {
		r0 = returnFunc(ctx, key, value, headers)
	} else {
		r0 = ret.Get(0).(int32)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, []byte, []byte, map[string]string) int64); ok {
		r1 = returnFunc(ctx, key, value, headers)
	} else {
		r1 = ret.Get(1).(int64)
	}
	if returnFunc, ok := ret.Get(2).(func(context.Context, []byte, []byte, map[string]string) error); ok {
		r2 = returnFunc(ctx, key, value, headers)
	} else {
		r2 = ret.Error(2)
	}
	return r0, r1, r2
}

// MockIProducer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockIProducer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - ctx context.Context
//   - key []byte
//   - value []byte
//   - headers map[string]string
func (_e *MockIProducer_Expecter) Send(ctx interface{}, key interface{}, value interface{}, headers interface{}) *MockIProducer_Send_Call {
	return &MockIProducer_Send_Call{Call: _e.mock.On("Send", ctx, key, value, headers)}
}

func (_c *MockIProducer_Send_Call) Run(run func(ctx context.Context, key []byte, value []byte, headers map[string]string)) *MockIProducer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 []byte
		if args[1] != nil {
			arg1 = args[1].([]byte)
		}
		var arg2 []byte
		if args[2] != nil {
			arg2 = args[2].([]byte)
		}
		var arg3 map[string]string
		if args[3] != nil {
			arg3 = args[3].(map[string]string)
		}
		run(
			arg0,
			arg1,
			arg2,
			arg3,
		)
	})
	return _c
}

func (_c *MockIProducer_Send_Call) Return(partition int32, offset int64, err error) *MockIProducer_Send_Call {
	_c.Call.Return(partition, offset, err)
	return _c
}

func (_c *MockIProducer_Send_Call) RunAndReturn(run func(ctx context.Context, key []byte, value []byte, headers map[string]string) (int32, int64, error)) *MockIProducer_Send_Call {
	_c.Call.Return(run)
	return _c
}

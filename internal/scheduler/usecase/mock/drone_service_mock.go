// Code generated by mockery v2.14.0. DO NOT EDIT.

package mock

import (
	context "context"

	food "github.com/dietdoctor/be-test/pkg/food/v1"
	mock "github.com/stretchr/testify/mock"
)

// DroneService is an autogenerated mock type for the DroneService type
type DroneService struct {
	mock.Mock
}

type DroneService_Expecter struct {
	mock *mock.Mock
}

func (_m *DroneService) EXPECT() *DroneService_Expecter {
	return &DroneService_Expecter{mock: &_m.Mock}
}

// CreateShipment provides a mock function with given fields: ctx, orderName
func (_m *DroneService) CreateShipment(ctx context.Context, orderName string) (*food.Shipment, error) {
	ret := _m.Called(ctx, orderName)

	var r0 *food.Shipment
	if rf, ok := ret.Get(0).(func(context.Context, string) *food.Shipment); ok {
		r0 = rf(ctx, orderName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*food.Shipment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orderName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DroneService_CreateShipment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateShipment'
type DroneService_CreateShipment_Call struct {
	*mock.Call
}

// CreateShipment is a helper method to define mock.On call
//  - ctx context.Context
//  - orderName string
func (_e *DroneService_Expecter) CreateShipment(ctx interface{}, orderName interface{}) *DroneService_CreateShipment_Call {
	return &DroneService_CreateShipment_Call{Call: _e.mock.On("CreateShipment", ctx, orderName)}
}

func (_c *DroneService_CreateShipment_Call) Run(run func(ctx context.Context, orderName string)) *DroneService_CreateShipment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DroneService_CreateShipment_Call) Return(_a0 *food.Shipment, _a1 error) *DroneService_CreateShipment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetShipment provides a mock function with given fields: ctx, orderName
func (_m *DroneService) GetShipment(ctx context.Context, orderName string) (*food.Shipment, error) {
	ret := _m.Called(ctx, orderName)

	var r0 *food.Shipment
	if rf, ok := ret.Get(0).(func(context.Context, string) *food.Shipment); ok {
		r0 = rf(ctx, orderName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*food.Shipment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orderName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DroneService_GetShipment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShipment'
type DroneService_GetShipment_Call struct {
	*mock.Call
}

// GetShipment is a helper method to define mock.On call
//  - ctx context.Context
//  - orderName string
func (_e *DroneService_Expecter) GetShipment(ctx interface{}, orderName interface{}) *DroneService_GetShipment_Call {
	return &DroneService_GetShipment_Call{Call: _e.mock.On("GetShipment", ctx, orderName)}
}

func (_c *DroneService_GetShipment_Call) Run(run func(ctx context.Context, orderName string)) *DroneService_GetShipment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DroneService_GetShipment_Call) Return(_a0 *food.Shipment, _a1 error) *DroneService_GetShipment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewDroneService interface {
	mock.TestingT
	Cleanup(func())
}

// NewDroneService creates a new instance of DroneService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDroneService(t mockConstructorTestingTNewDroneService) *DroneService {
	mock := &DroneService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
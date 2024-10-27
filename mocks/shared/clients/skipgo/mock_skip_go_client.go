// Code generated by mockery v2.46.2. DO NOT EDIT.

package skipgo

import (
	context "context"
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	skipgo "github.com/skip-mev/go-fast-solver/shared/clients/skipgo"
)

// MockSkipGoClient is an autogenerated mock type for the SkipGoClient type
type MockSkipGoClient struct {
	mock.Mock
}

type MockSkipGoClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSkipGoClient) EXPECT() *MockSkipGoClient_Expecter {
	return &MockSkipGoClient_Expecter{mock: &_m.Mock}
}

// Balance provides a mock function with given fields: ctx, chainID, address, denom
func (_m *MockSkipGoClient) Balance(ctx context.Context, chainID string, address string, denom string) (string, error) {
	ret := _m.Called(ctx, chainID, address, denom)

	if len(ret) == 0 {
		panic("no return value specified for Balance")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (string, error)); ok {
		return rf(ctx, chainID, address, denom)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) string); ok {
		r0 = rf(ctx, chainID, address, denom)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, chainID, address, denom)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSkipGoClient_Balance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Balance'
type MockSkipGoClient_Balance_Call struct {
	*mock.Call
}

// Balance is a helper method to define mock.On call
//   - ctx context.Context
//   - chainID string
//   - address string
//   - denom string
func (_e *MockSkipGoClient_Expecter) Balance(ctx interface{}, chainID interface{}, address interface{}, denom interface{}) *MockSkipGoClient_Balance_Call {
	return &MockSkipGoClient_Balance_Call{Call: _e.mock.On("Balance", ctx, chainID, address, denom)}
}

func (_c *MockSkipGoClient_Balance_Call) Run(run func(ctx context.Context, chainID string, address string, denom string)) *MockSkipGoClient_Balance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockSkipGoClient_Balance_Call) Return(_a0 string, _a1 error) *MockSkipGoClient_Balance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSkipGoClient_Balance_Call) RunAndReturn(run func(context.Context, string, string, string) (string, error)) *MockSkipGoClient_Balance_Call {
	_c.Call.Return(run)
	return _c
}

// Msgs provides a mock function with given fields: ctx, sourceAssetDenom, sourceAssetChainID, sourceChainAddress, destAssetDenom, destAssetChainID, destChainAddress, amountIn, amountOut, addressList, operations
func (_m *MockSkipGoClient) Msgs(ctx context.Context, sourceAssetDenom string, sourceAssetChainID string, sourceChainAddress string, destAssetDenom string, destAssetChainID string, destChainAddress string, amountIn *big.Int, amountOut *big.Int, addressList []string, operations []any) ([]skipgo.Tx, error) {
	ret := _m.Called(ctx, sourceAssetDenom, sourceAssetChainID, sourceChainAddress, destAssetDenom, destAssetChainID, destChainAddress, amountIn, amountOut, addressList, operations)

	if len(ret) == 0 {
		panic("no return value specified for Msgs")
	}

	var r0 []skipgo.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string, string, *big.Int, *big.Int, []string, []any) ([]skipgo.Tx, error)); ok {
		return rf(ctx, sourceAssetDenom, sourceAssetChainID, sourceChainAddress, destAssetDenom, destAssetChainID, destChainAddress, amountIn, amountOut, addressList, operations)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string, string, *big.Int, *big.Int, []string, []any) []skipgo.Tx); ok {
		r0 = rf(ctx, sourceAssetDenom, sourceAssetChainID, sourceChainAddress, destAssetDenom, destAssetChainID, destChainAddress, amountIn, amountOut, addressList, operations)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]skipgo.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, string, string, *big.Int, *big.Int, []string, []any) error); ok {
		r1 = rf(ctx, sourceAssetDenom, sourceAssetChainID, sourceChainAddress, destAssetDenom, destAssetChainID, destChainAddress, amountIn, amountOut, addressList, operations)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSkipGoClient_Msgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Msgs'
type MockSkipGoClient_Msgs_Call struct {
	*mock.Call
}

// Msgs is a helper method to define mock.On call
//   - ctx context.Context
//   - sourceAssetDenom string
//   - sourceAssetChainID string
//   - sourceChainAddress string
//   - destAssetDenom string
//   - destAssetChainID string
//   - destChainAddress string
//   - amountIn *big.Int
//   - amountOut *big.Int
//   - addressList []string
//   - operations []any
func (_e *MockSkipGoClient_Expecter) Msgs(ctx interface{}, sourceAssetDenom interface{}, sourceAssetChainID interface{}, sourceChainAddress interface{}, destAssetDenom interface{}, destAssetChainID interface{}, destChainAddress interface{}, amountIn interface{}, amountOut interface{}, addressList interface{}, operations interface{}) *MockSkipGoClient_Msgs_Call {
	return &MockSkipGoClient_Msgs_Call{Call: _e.mock.On("Msgs", ctx, sourceAssetDenom, sourceAssetChainID, sourceChainAddress, destAssetDenom, destAssetChainID, destChainAddress, amountIn, amountOut, addressList, operations)}
}

func (_c *MockSkipGoClient_Msgs_Call) Run(run func(ctx context.Context, sourceAssetDenom string, sourceAssetChainID string, sourceChainAddress string, destAssetDenom string, destAssetChainID string, destChainAddress string, amountIn *big.Int, amountOut *big.Int, addressList []string, operations []any)) *MockSkipGoClient_Msgs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string), args[5].(string), args[6].(string), args[7].(*big.Int), args[8].(*big.Int), args[9].([]string), args[10].([]any))
	})
	return _c
}

func (_c *MockSkipGoClient_Msgs_Call) Return(_a0 []skipgo.Tx, _a1 error) *MockSkipGoClient_Msgs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSkipGoClient_Msgs_Call) RunAndReturn(run func(context.Context, string, string, string, string, string, string, *big.Int, *big.Int, []string, []any) ([]skipgo.Tx, error)) *MockSkipGoClient_Msgs_Call {
	_c.Call.Return(run)
	return _c
}

// Route provides a mock function with given fields: ctx, sourceAssetDenom, sourceAssetChainID, destAssetDenom, destAssetChainID, amountIn
func (_m *MockSkipGoClient) Route(ctx context.Context, sourceAssetDenom string, sourceAssetChainID string, destAssetDenom string, destAssetChainID string, amountIn *big.Int) (*skipgo.RouteResponse, error) {
	ret := _m.Called(ctx, sourceAssetDenom, sourceAssetChainID, destAssetDenom, destAssetChainID, amountIn)

	if len(ret) == 0 {
		panic("no return value specified for Route")
	}

	var r0 *skipgo.RouteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, *big.Int) (*skipgo.RouteResponse, error)); ok {
		return rf(ctx, sourceAssetDenom, sourceAssetChainID, destAssetDenom, destAssetChainID, amountIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, *big.Int) *skipgo.RouteResponse); ok {
		r0 = rf(ctx, sourceAssetDenom, sourceAssetChainID, destAssetDenom, destAssetChainID, amountIn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*skipgo.RouteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, *big.Int) error); ok {
		r1 = rf(ctx, sourceAssetDenom, sourceAssetChainID, destAssetDenom, destAssetChainID, amountIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSkipGoClient_Route_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Route'
type MockSkipGoClient_Route_Call struct {
	*mock.Call
}

// Route is a helper method to define mock.On call
//   - ctx context.Context
//   - sourceAssetDenom string
//   - sourceAssetChainID string
//   - destAssetDenom string
//   - destAssetChainID string
//   - amountIn *big.Int
func (_e *MockSkipGoClient_Expecter) Route(ctx interface{}, sourceAssetDenom interface{}, sourceAssetChainID interface{}, destAssetDenom interface{}, destAssetChainID interface{}, amountIn interface{}) *MockSkipGoClient_Route_Call {
	return &MockSkipGoClient_Route_Call{Call: _e.mock.On("Route", ctx, sourceAssetDenom, sourceAssetChainID, destAssetDenom, destAssetChainID, amountIn)}
}

func (_c *MockSkipGoClient_Route_Call) Run(run func(ctx context.Context, sourceAssetDenom string, sourceAssetChainID string, destAssetDenom string, destAssetChainID string, amountIn *big.Int)) *MockSkipGoClient_Route_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string), args[5].(*big.Int))
	})
	return _c
}

func (_c *MockSkipGoClient_Route_Call) Return(_a0 *skipgo.RouteResponse, _a1 error) *MockSkipGoClient_Route_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSkipGoClient_Route_Call) RunAndReturn(run func(context.Context, string, string, string, string, *big.Int) (*skipgo.RouteResponse, error)) *MockSkipGoClient_Route_Call {
	_c.Call.Return(run)
	return _c
}

// Status provides a mock function with given fields: ctx, tx, chainID
func (_m *MockSkipGoClient) Status(ctx context.Context, tx skipgo.TxHash, chainID string) (*skipgo.StatusResponse, error) {
	ret := _m.Called(ctx, tx, chainID)

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 *skipgo.StatusResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, skipgo.TxHash, string) (*skipgo.StatusResponse, error)); ok {
		return rf(ctx, tx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, skipgo.TxHash, string) *skipgo.StatusResponse); ok {
		r0 = rf(ctx, tx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*skipgo.StatusResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, skipgo.TxHash, string) error); ok {
		r1 = rf(ctx, tx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSkipGoClient_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type MockSkipGoClient_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
//   - ctx context.Context
//   - tx skipgo.TxHash
//   - chainID string
func (_e *MockSkipGoClient_Expecter) Status(ctx interface{}, tx interface{}, chainID interface{}) *MockSkipGoClient_Status_Call {
	return &MockSkipGoClient_Status_Call{Call: _e.mock.On("Status", ctx, tx, chainID)}
}

func (_c *MockSkipGoClient_Status_Call) Run(run func(ctx context.Context, tx skipgo.TxHash, chainID string)) *MockSkipGoClient_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(skipgo.TxHash), args[2].(string))
	})
	return _c
}

func (_c *MockSkipGoClient_Status_Call) Return(_a0 *skipgo.StatusResponse, _a1 error) *MockSkipGoClient_Status_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSkipGoClient_Status_Call) RunAndReturn(run func(context.Context, skipgo.TxHash, string) (*skipgo.StatusResponse, error)) *MockSkipGoClient_Status_Call {
	_c.Call.Return(run)
	return _c
}

// SubmitTx provides a mock function with given fields: ctx, tx, chainID
func (_m *MockSkipGoClient) SubmitTx(ctx context.Context, tx []byte, chainID string) (skipgo.TxHash, error) {
	ret := _m.Called(ctx, tx, chainID)

	if len(ret) == 0 {
		panic("no return value specified for SubmitTx")
	}

	var r0 skipgo.TxHash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) (skipgo.TxHash, error)); ok {
		return rf(ctx, tx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) skipgo.TxHash); ok {
		r0 = rf(ctx, tx, chainID)
	} else {
		r0 = ret.Get(0).(skipgo.TxHash)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, string) error); ok {
		r1 = rf(ctx, tx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSkipGoClient_SubmitTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubmitTx'
type MockSkipGoClient_SubmitTx_Call struct {
	*mock.Call
}

// SubmitTx is a helper method to define mock.On call
//   - ctx context.Context
//   - tx []byte
//   - chainID string
func (_e *MockSkipGoClient_Expecter) SubmitTx(ctx interface{}, tx interface{}, chainID interface{}) *MockSkipGoClient_SubmitTx_Call {
	return &MockSkipGoClient_SubmitTx_Call{Call: _e.mock.On("SubmitTx", ctx, tx, chainID)}
}

func (_c *MockSkipGoClient_SubmitTx_Call) Run(run func(ctx context.Context, tx []byte, chainID string)) *MockSkipGoClient_SubmitTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte), args[2].(string))
	})
	return _c
}

func (_c *MockSkipGoClient_SubmitTx_Call) Return(_a0 skipgo.TxHash, _a1 error) *MockSkipGoClient_SubmitTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSkipGoClient_SubmitTx_Call) RunAndReturn(run func(context.Context, []byte, string) (skipgo.TxHash, error)) *MockSkipGoClient_SubmitTx_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSkipGoClient creates a new instance of MockSkipGoClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSkipGoClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSkipGoClient {
	mock := &MockSkipGoClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

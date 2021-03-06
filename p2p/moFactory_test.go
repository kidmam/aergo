// Code generated by mockery v1.0.0. DO NOT EDIT.
package p2p

import mock "github.com/stretchr/testify/mock"
import types "github.com/aergoio/aergo/types"

// MockMoFactory is an autogenerated mock type for the moFactory type
type MockMoFactory struct {
	mock.Mock
}

// newMsgBlkBroadcastOrder provides a mock function with given fields: noticeMsg
func (_m *MockMoFactory) newMsgBlkBroadcastOrder(noticeMsg *types.NewBlockNotice) msgOrder {
	ret := _m.Called(noticeMsg)

	var r0 msgOrder
	if rf, ok := ret.Get(0).(func(*types.NewBlockNotice) msgOrder); ok {
		r0 = rf(noticeMsg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(msgOrder)
		}
	}

	return r0
}

// newMsgRequestOrder provides a mock function with given fields: expectedResponse, protocolID, message
func (_m *MockMoFactory) newMsgRequestOrder(expectResponse bool, protocolID SubProtocol, message pbMessage) msgOrder {
	ret := _m.Called(expectResponse, protocolID, message)

	var r0 msgOrder
	if rf, ok := ret.Get(0).(func(bool, SubProtocol, pbMessage) msgOrder); ok {
		r0 = rf(expectResponse, protocolID, message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(msgOrder)
		}
	}

	return r0
}

// newMsgBlockRequestOrder provides a mock function with given fields: respReceiver, protocolID, message
func (_m *MockMoFactory) newMsgBlockRequestOrder(respReceiver ResponseReceiver, protocolID SubProtocol, message pbMessage) msgOrder {
	ret := _m.Called(respReceiver, protocolID, message)

	var r0 msgOrder
	if rf, ok := ret.Get(0).(func(ResponseReceiver, SubProtocol, pbMessage) msgOrder); ok {
		r0 = rf(respReceiver, protocolID, message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(msgOrder)
		}
	}

	return r0
}

// newMsgResponseOrder provides a mock function with given fields: reqID, protocolID, message
func (_m *MockMoFactory) newMsgResponseOrder(reqID MsgID, protocolID SubProtocol, message pbMessage) msgOrder {
	ret := _m.Called(reqID, protocolID, message)

	var r0 msgOrder
	if rf, ok := ret.Get(0).(func(MsgID, SubProtocol, pbMessage) msgOrder); ok {
		r0 = rf(reqID, protocolID, message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(msgOrder)
		}
	}

	return r0
}

// newMsgTxBroadcastOrder provides a mock function with given fields: message
func (_m *MockMoFactory) newMsgTxBroadcastOrder(noticeMsg *types.NewTransactionsNotice) msgOrder {
	ret := _m.Called(noticeMsg)

	var r0 msgOrder
	if rf, ok := ret.Get(0).(func(*types.NewTransactionsNotice) msgOrder); ok {
		r0 = rf(noticeMsg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(msgOrder)
		}
	}

	return r0
}


// newMsgBPBroadcastOrder provides a mock function with given fields: message
func (_m *MockMoFactory) newMsgBPBroadcastOrder(noticeMsg *types.BlockProducedNotice) msgOrder {
	ret := _m.Called(noticeMsg)

	var r0 msgOrder
	if rf, ok := ret.Get(0).(func(*types.BlockProducedNotice) msgOrder); ok {
		r0 = rf(noticeMsg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(msgOrder)
		}
	}

	return r0
}

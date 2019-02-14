// Code generated by mockery v1.0.0. DO NOT EDIT.
package p2p

import (
	"github.com/aergoio/aergo/p2p/audit"
	"github.com/stretchr/testify/mock"
	"time"
)
import "github.com/libp2p/go-libp2p-peer"
import "github.com/aergoio/aergo/types"

// MockRemotePeer is an autogenerated mock type for the RemotePeer type
type MockRemotePeer struct {
	mock.Mock
}

// consumeRequest provides a mock function with given fields: msgID
func (_m *MockRemotePeer) consumeRequest(msgID MsgID) {
	_m.Called(msgID)
}

// pushTxsNotice provides a mock function with given fields: txHashes
func (_m *MockRemotePeer) pushTxsNotice(txHashes []types.TxID) {
	_m.Called(txHashes)
}

// runPeer provides a mock function with given fields:
func (_m *MockRemotePeer) runPeer() {
	_m.Called()
}

// sendMessage provides a mock function with given fields: msg
func (_m *MockRemotePeer) sendMessage(msg msgOrder) {
	_m.Called(msg)
}

func (_m *MockRemotePeer) GetReceiver(id MsgID) ResponseReceiver {
	ret := _m.Called(id)

	var r0 ResponseReceiver
	if rf, ok := ret.Get(0).(func(id MsgID) ResponseReceiver); ok {
		r0 = rf(id)
	} else {
		rr0 := ret.Get(0)
		if rr0 == nil {
			r0 = nil
		} else {
			r0 = rr0.(ResponseReceiver)
		}
	}

	return r0
}

func (_m *MockRemotePeer) sendAndWaitMessage(msg msgOrder, ttl time.Duration) error {
	ret := _m.Called(msg, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(msg msgOrder, ttl time.Duration) error); ok {
		r0 = rf(msg, ttl)
	} else {
		rr0 := ret.Get(0)
		if rr0 == nil {
			r0 = nil
		} else {
			r0 = rr0.(error)
		}
	}

	return r0
}

// stop provides a mock function with given fields:
func (_m *MockRemotePeer) stop() {
	_m.Called()
}

// updateBlkCache provides a mock function with given fields: hash
func (_m *MockRemotePeer) updateBlkCache(blkHash []byte, blkNumber uint64) bool {
	ret := _m.Called(blkHash, blkNumber)

	var r0 bool
	if rf, ok := ret.Get(0).(func(blkHash []byte, blkNumber uint64) bool); ok {
		r0 = rf(blkHash, blkNumber)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// updateTxCache provides a mock function with given fields: hashes
func (_m *MockRemotePeer) updateTxCache(hashes []types.TxID) []types.TxID {
	ret := _m.Called(hashes)

	var r0 []types.TxID
	if rf, ok := ret.Get(0).(func([]types.TxID) []types.TxID); ok {
		r0 = rf(hashes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.TxID)
		}
	}

	return r0
}

func (_m *MockRemotePeer) AddPenalty(penalty audit.Penalty) {
	_m.Called(penalty)
}

// updateBlkCache provides a mock function with given fields: hash
func (_m *MockRemotePeer) updateLastNotice(blkHash []byte, blkNumber uint64) {
	_m.Called(blkHash, blkNumber)
}

// ID provides a mock function with given fields:
func (_m *MockRemotePeer) ID() peer.ID {
	ret := _m.Called()

	var r0 peer.ID
	if rf, ok := ret.Get(0).(func() peer.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(peer.ID)
	}

	return r0
}

// MF provides a mock function with given fields:
func (_m *MockRemotePeer) MF() moFactory {
	ret := _m.Called()

	var r0 moFactory
	if rf, ok := ret.Get(0).(func() moFactory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(moFactory)
		}
	}

	return r0
}

// Meta provides a mock function with given fields:
func (_m *MockRemotePeer) Meta() PeerMeta {
	ret := _m.Called()

	var r0 PeerMeta
	if rf, ok := ret.Get(0).(func() PeerMeta); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(PeerMeta)
	}

	return r0
}

func (_m *MockRemotePeer) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

func (_m *MockRemotePeer) ManageNumber() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// State provides a mock function with given fields:
func (_m *MockRemotePeer) State() types.PeerState {
	ret := _m.Called()

	var r0 types.PeerState
	if rf, ok := ret.Get(0).(func() types.PeerState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.PeerState)
	}

	return r0
}


// LastNotice provides a mock function with given fields:
func (_m *MockRemotePeer) LastNotice() *LastBlockStatus {
	ret := _m.Called()

	var r0 *LastBlockStatus
	if rf, ok := ret.Get(0).(func() *LastBlockStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(*LastBlockStatus)
	}

	return r0
}

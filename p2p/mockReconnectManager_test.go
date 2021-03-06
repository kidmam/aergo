// Code generated by mockery v1.0.0. DO NOT EDIT.
package p2p

import mock "github.com/stretchr/testify/mock"
import peer "github.com/libp2p/go-libp2p-peer"

// MockReconnectManager is an autogenerated mock type for the ReconnectManager type
type MockReconnectManager struct {
	mock.Mock
}

// jobFinished provides a mock function with given fields: pid
func (_m *MockReconnectManager) jobFinished(pid peer.ID) {
	_m.Called(pid)
}

// AddJob provides a mock function with given fields: meta
func (_m *MockReconnectManager) AddJob(meta PeerMeta) {
	_m.Called(meta)
}

// CancelJob provides a mock function with given fields: pid
func (_m *MockReconnectManager) CancelJob(pid peer.ID) {
	_m.Called(pid)
}

// Stop provides a mock function with given fields:
func (_m *MockReconnectManager) Stop() {
	_m.Called()
}

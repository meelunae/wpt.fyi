// Code generated by MockGen. DO NOT EDIT.
// Source: api/query/cache/monitor/monitor.go

// Package monitor is a generated GoMock package.
package monitor

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockRuntime is a mock of Runtime interface
type MockRuntime struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeMockRecorder
}

// MockRuntimeMockRecorder is the mock recorder for MockRuntime
type MockRuntimeMockRecorder struct {
	mock *MockRuntime
}

// NewMockRuntime creates a new mock instance
func NewMockRuntime(ctrl *gomock.Controller) *MockRuntime {
	mock := &MockRuntime{ctrl: ctrl}
	mock.recorder = &MockRuntimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRuntime) EXPECT() *MockRuntimeMockRecorder {
	return m.recorder
}

// GetHeapBytes mocks base method
func (m *MockRuntime) GetHeapBytes() uint64 {
	ret := m.ctrl.Call(m, "GetHeapBytes")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetHeapBytes indicates an expected call of GetHeapBytes
func (mr *MockRuntimeMockRecorder) GetHeapBytes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeapBytes", reflect.TypeOf((*MockRuntime)(nil).GetHeapBytes))
}

// MockMonitor is a mock of Monitor interface
type MockMonitor struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorMockRecorder
}

// MockMonitorMockRecorder is the mock recorder for MockMonitor
type MockMonitorMockRecorder struct {
	mock *MockMonitor
}

// NewMockMonitor creates a new mock instance
func NewMockMonitor(ctrl *gomock.Controller) *MockMonitor {
	mock := &MockMonitor{ctrl: ctrl}
	mock.recorder = &MockMonitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMonitor) EXPECT() *MockMonitorMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockMonitor) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockMonitorMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockMonitor)(nil).Start))
}

// Stop mocks base method
func (m *MockMonitor) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockMonitorMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockMonitor)(nil).Stop))
}

// SetInterval mocks base method
func (m *MockMonitor) SetInterval(arg0 time.Duration) error {
	ret := m.ctrl.Call(m, "SetInterval", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInterval indicates an expected call of SetInterval
func (mr *MockMonitorMockRecorder) SetInterval(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInterval", reflect.TypeOf((*MockMonitor)(nil).SetInterval), arg0)
}

// SetMaxHeapBytes mocks base method
func (m *MockMonitor) SetMaxHeapBytes(arg0 uint64) error {
	ret := m.ctrl.Call(m, "SetMaxHeapBytes", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMaxHeapBytes indicates an expected call of SetMaxHeapBytes
func (mr *MockMonitorMockRecorder) SetMaxHeapBytes(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxHeapBytes", reflect.TypeOf((*MockMonitor)(nil).SetMaxHeapBytes), arg0)
}
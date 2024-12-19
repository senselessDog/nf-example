// Code generated by MockGen. DO NOT EDIT.
// Source: internal/sbi/server.go
//
// Generated by this command:
//
//	mockgen -source=internal/sbi/server.go -package sbi
//
// Package sbi is a generated GoMock package.
package sbi

import (
	context "github.com/andy89923/nf-example/internal/context"
	processor "github.com/andy89923/nf-example/internal/sbi/processor"
	factory "github.com/andy89923/nf-example/pkg/factory"
	gomock "go.uber.org/mock/gomock"
	reflect "reflect"
)

// MocknfApp is a mock of nfApp interface.
type MocknfApp struct {
	ctrl     *gomock.Controller
	recorder *MocknfAppMockRecorder
}

// MocknfAppMockRecorder is the mock recorder for MocknfApp.
type MocknfAppMockRecorder struct {
	mock *MocknfApp
}

// NewMocknfApp creates a new mock instance.
func NewMocknfApp(ctrl *gomock.Controller) *MocknfApp {
	mock := &MocknfApp{ctrl: ctrl}
	mock.recorder = &MocknfAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocknfApp) EXPECT() *MocknfAppMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MocknfApp) Config() *factory.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*factory.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MocknfAppMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MocknfApp)(nil).Config))
}

// Context mocks base method.
func (m *MocknfApp) Context() *context.NFContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(*context.NFContext)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MocknfAppMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MocknfApp)(nil).Context))
}

// Processor mocks base method.
func (m *MocknfApp) Processor() *processor.Processor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Processor")
	ret0, _ := ret[0].(*processor.Processor)
	return ret0
}

// Processor indicates an expected call of Processor.
func (mr *MocknfAppMockRecorder) Processor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Processor", reflect.TypeOf((*MocknfApp)(nil).Processor))
}

// SetLogEnable mocks base method.
func (m *MocknfApp) SetLogEnable(enable bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogEnable", enable)
}

// SetLogEnable indicates an expected call of SetLogEnable.
func (mr *MocknfAppMockRecorder) SetLogEnable(enable any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogEnable", reflect.TypeOf((*MocknfApp)(nil).SetLogEnable), enable)
}

// SetLogLevel mocks base method.
func (m *MocknfApp) SetLogLevel(level string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogLevel", level)
}

// SetLogLevel indicates an expected call of SetLogLevel.
func (mr *MocknfAppMockRecorder) SetLogLevel(level any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogLevel", reflect.TypeOf((*MocknfApp)(nil).SetLogLevel), level)
}

// SetReportCaller mocks base method.
func (m *MocknfApp) SetReportCaller(reportCaller bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReportCaller", reportCaller)
}

// SetReportCaller indicates an expected call of SetReportCaller.
func (mr *MocknfAppMockRecorder) SetReportCaller(reportCaller any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReportCaller", reflect.TypeOf((*MocknfApp)(nil).SetReportCaller), reportCaller)
}

// Start mocks base method.
func (m *MocknfApp) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MocknfAppMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MocknfApp)(nil).Start))
}

// Terminate mocks base method.
func (m *MocknfApp) Terminate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Terminate")
}

// Terminate indicates an expected call of Terminate.
func (mr *MocknfAppMockRecorder) Terminate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terminate", reflect.TypeOf((*MocknfApp)(nil).Terminate))
}

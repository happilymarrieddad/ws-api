// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/happilymarrieddad/ws-api/internal/wsclient (interfaces: WSClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	wsclient "github.com/happilymarrieddad/ws-api/internal/wsclient"
	types "github.com/happilymarrieddad/ws-api/types"
)

// MockWSClient is a mock of WSClient interface.
type MockWSClient struct {
	ctrl     *gomock.Controller
	recorder *MockWSClientMockRecorder
}

// MockWSClientMockRecorder is the mock recorder for MockWSClient.
type MockWSClientMockRecorder struct {
	mock *MockWSClient
}

// NewMockWSClient creates a new mock instance.
func NewMockWSClient(ctrl *gomock.Controller) *MockWSClient {
	mock := &MockWSClient{ctrl: ctrl}
	mock.recorder = &MockWSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWSClient) EXPECT() *MockWSClientMockRecorder {
	return m.recorder
}

// GetWeatherAtLongLat mocks base method.
func (m *MockWSClient) GetWeatherAtLongLat(arg0 context.Context, arg1, arg2 float64, arg3 *wsclient.TempType) (*types.GetWeatherResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeatherAtLongLat", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.GetWeatherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWeatherAtLongLat indicates an expected call of GetWeatherAtLongLat.
func (mr *MockWSClientMockRecorder) GetWeatherAtLongLat(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeatherAtLongLat", reflect.TypeOf((*MockWSClient)(nil).GetWeatherAtLongLat), arg0, arg1, arg2, arg3)
}

// GetWeatherConditonFromTempType mocks base method.
func (m *MockWSClient) GetWeatherConditonFromTempType(arg0 wsclient.TempType, arg1 float64) (wsclient.WeatherCondition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeatherConditonFromTempType", arg0, arg1)
	ret0, _ := ret[0].(wsclient.WeatherCondition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWeatherConditonFromTempType indicates an expected call of GetWeatherConditonFromTempType.
func (mr *MockWSClientMockRecorder) GetWeatherConditonFromTempType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeatherConditonFromTempType", reflect.TypeOf((*MockWSClient)(nil).GetWeatherConditonFromTempType), arg0, arg1)
}

// GetWeatherDataAtLongLat mocks base method.
func (m *MockWSClient) GetWeatherDataAtLongLat(arg0 context.Context, arg1, arg2 float64, arg3 *wsclient.TempType) (*wsclient.WeatherResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeatherDataAtLongLat", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*wsclient.WeatherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWeatherDataAtLongLat indicates an expected call of GetWeatherDataAtLongLat.
func (mr *MockWSClientMockRecorder) GetWeatherDataAtLongLat(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeatherDataAtLongLat", reflect.TypeOf((*MockWSClient)(nil).GetWeatherDataAtLongLat), arg0, arg1, arg2, arg3)
}

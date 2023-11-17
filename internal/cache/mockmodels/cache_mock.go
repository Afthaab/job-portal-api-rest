// Code generated by MockGen. DO NOT EDIT.
// Source: cache.go
//
// Generated by this command:
//
//	mockgen -source=cache.go -destination=mockmodels/cache_mock.go -package=mockmodels
//
// Package mockmodels is a generated GoMock package.
package mockmodels

import (
	context "context"
	reflect "reflect"

	models "github.com/afthaab/job-portal/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockCaching is a mock of Caching interface.
type MockCaching struct {
	ctrl     *gomock.Controller
	recorder *MockCachingMockRecorder
}

// MockCachingMockRecorder is the mock recorder for MockCaching.
type MockCachingMockRecorder struct {
	mock *MockCaching
}

// NewMockCaching creates a new mock instance.
func NewMockCaching(ctrl *gomock.Controller) *MockCaching {
	mock := &MockCaching{ctrl: ctrl}
	mock.recorder = &MockCachingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCaching) EXPECT() *MockCachingMockRecorder {
	return m.recorder
}

// AddToTheCache mocks base method.
func (m *MockCaching) AddToTheCache(ctx context.Context, jid uint, jobData models.Jobs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToTheCache", ctx, jid, jobData)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToTheCache indicates an expected call of AddToTheCache.
func (mr *MockCachingMockRecorder) AddToTheCache(ctx, jid, jobData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToTheCache", reflect.TypeOf((*MockCaching)(nil).AddToTheCache), ctx, jid, jobData)
}

// GetTheCacheData mocks base method.
func (m *MockCaching) GetTheCacheData(ctx context.Context, jid uint) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTheCacheData", ctx, jid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTheCacheData indicates an expected call of GetTheCacheData.
func (mr *MockCachingMockRecorder) GetTheCacheData(ctx, jid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTheCacheData", reflect.TypeOf((*MockCaching)(nil).GetTheCacheData), ctx, jid)
}

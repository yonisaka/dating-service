// Code generated by MockGen. DO NOT EDIT.
// Source: health.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// GoMockHealthRepo is a mock of HealthRepo interface.
type GoMockHealthRepo struct {
	ctrl     *gomock.Controller
	recorder *GoMockHealthRepoMockRecorder
}

// GoMockHealthRepoMockRecorder is the mock recorder for GoMockHealthRepo.
type GoMockHealthRepoMockRecorder struct {
	mock *GoMockHealthRepo
}

// NewGoMockHealthRepo creates a new mock instance.
func NewGoMockHealthRepo(ctrl *gomock.Controller) *GoMockHealthRepo {
	mock := &GoMockHealthRepo{ctrl: ctrl}
	mock.recorder = &GoMockHealthRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *GoMockHealthRepo) EXPECT() *GoMockHealthRepoMockRecorder {
	return m.recorder
}

// GetLiveness mocks base method.
func (m *GoMockHealthRepo) GetLiveness(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLiveness", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetLiveness indicates an expected call of GetLiveness.
func (mr *GoMockHealthRepoMockRecorder) GetLiveness(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLiveness", reflect.TypeOf((*GoMockHealthRepo)(nil).GetLiveness), ctx)
}

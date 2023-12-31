// Code generated by MockGen. DO NOT EDIT.
// Source: user_preference.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// GoMockUserPreferenceRepo is a mock of UserPreferenceRepo interface.
type GoMockUserPreferenceRepo struct {
	ctrl     *gomock.Controller
	recorder *GoMockUserPreferenceRepoMockRecorder
}

// GoMockUserPreferenceRepoMockRecorder is the mock recorder for GoMockUserPreferenceRepo.
type GoMockUserPreferenceRepoMockRecorder struct {
	mock *GoMockUserPreferenceRepo
}

// NewGoMockUserPreferenceRepo creates a new mock instance.
func NewGoMockUserPreferenceRepo(ctrl *gomock.Controller) *GoMockUserPreferenceRepo {
	mock := &GoMockUserPreferenceRepo{ctrl: ctrl}
	mock.recorder = &GoMockUserPreferenceRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *GoMockUserPreferenceRepo) EXPECT() *GoMockUserPreferenceRepoMockRecorder {
	return m.recorder
}

// FindByUserID mocks base method.
func (m *GoMockUserPreferenceRepo) FindByUserID(ctx context.Context, userID int64) (*UserPreference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", ctx, userID)
	ret0, _ := ret[0].(*UserPreference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID.
func (mr *GoMockUserPreferenceRepoMockRecorder) FindByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*GoMockUserPreferenceRepo)(nil).FindByUserID), ctx, userID)
}

// Update mocks base method.
func (m *GoMockUserPreferenceRepo) Update(ctx context.Context, userPreference UserPreference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, userPreference)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *GoMockUserPreferenceRepoMockRecorder) Update(ctx, userPreference interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*GoMockUserPreferenceRepo)(nil).Update), ctx, userPreference)
}

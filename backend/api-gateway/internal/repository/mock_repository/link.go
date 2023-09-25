// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/thinhle/Projects/tinyurl/backend/api-gateway/internal/repository/link.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
        model "github.com/thinhlvv/tinyurl/backend/api-gateway/internal/model"
)

// MockLinker is a mock of Linker interface.
type MockLinker struct {
        ctrl     *gomock.Controller
        recorder *MockLinkerMockRecorder
}

// MockLinkerMockRecorder is the mock recorder for MockLinker.
type MockLinkerMockRecorder struct {
        mock *MockLinker
}

// NewMockLinker creates a new mock instance.
func NewMockLinker(ctrl *gomock.Controller) *MockLinker {
        mock := &MockLinker{ctrl: ctrl}
        mock.recorder = &MockLinkerMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLinker) EXPECT() *MockLinkerMockRecorder {
        return m.recorder
}

// Create mocks base method.
func (m *MockLinker) Create(link model.Link) (*model.Link, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Create", link)
        ret0, _ := ret[0].(*model.Link)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLinkerMockRecorder) Create(link interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLinker)(nil).Create), link)
}

// GetByLongLink mocks base method.
func (m *MockLinker) GetByLongLink(longLink string) (*model.Link, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetByLongLink", longLink)
        ret0, _ := ret[0].(*model.Link)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetByLongLink indicates an expected call of GetByLongLink.
func (mr *MockLinkerMockRecorder) GetByLongLink(longLink interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByLongLink", reflect.TypeOf((*MockLinker)(nil).GetByLongLink), longLink)
}

// GetByShortLink mocks base method.
func (m *MockLinker) GetByShortLink(shortLink string) (*model.Link, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetByShortLink", shortLink)
        ret0, _ := ret[0].(*model.Link)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetByShortLink indicates an expected call of GetByShortLink.
func (mr *MockLinkerMockRecorder) GetByShortLink(shortLink interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByShortLink", reflect.TypeOf((*MockLinker)(nil).GetByShortLink), shortLink)
}

// Update mocks base method.
func (m *MockLinker) Update(link *model.Link) (*model.Link, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Update", link)
        ret0, _ := ret[0].(*model.Link)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockLinkerMockRecorder) Update(link interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLinker)(nil).Update), link)
}
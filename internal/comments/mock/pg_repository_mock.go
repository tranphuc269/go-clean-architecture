// Code generated by MockGen. DO NOT EDIT.
// Source: pg_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "go-clean-architecture/internal/models"
	utils "go-clean-architecture/pkg/utils"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockRepository) Create(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, comment)
	ret0, _ := ret[0].(*models.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(ctx, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, comment)
}

// Update mocks base method
func (m *MockRepository) Update(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, comment)
	ret0, _ := ret[0].(*models.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockRepositoryMockRecorder) Update(ctx, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, comment)
}

// Delete mocks base method
func (m *MockRepository) Delete(ctx context.Context, commentID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, commentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(ctx, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, commentID)
}

// GetByID mocks base method
func (m *MockRepository) GetByID(ctx context.Context, commentID uuid.UUID) (*models.CommentBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, commentID)
	ret0, _ := ret[0].(*models.CommentBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockRepositoryMockRecorder) GetByID(ctx, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID), ctx, commentID)
}

// GetAllByNewsID mocks base method
func (m *MockRepository) GetAllByNewsID(ctx context.Context, newsID uuid.UUID, query *utils.PaginationQuery) (*models.CommentsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByNewsID", ctx, newsID, query)
	ret0, _ := ret[0].(*models.CommentsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByNewsID indicates an expected call of GetAllByNewsID
func (mr *MockRepositoryMockRecorder) GetAllByNewsID(ctx, newsID, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByNewsID", reflect.TypeOf((*MockRepository)(nil).GetAllByNewsID), ctx, newsID, query)
}

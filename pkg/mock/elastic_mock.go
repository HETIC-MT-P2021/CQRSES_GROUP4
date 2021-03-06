// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/database/elasticsearch/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	database "github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	gomock "github.com/golang/mock/gomock"
)

// MockElasticRepository is a mock of ElasticRepository interface.
type MockElasticRepository struct {
	ctrl     *gomock.Controller
	recorder *MockElasticRepositoryMockRecorder
}

// MockElasticRepositoryMockRecorder is the mock recorder for MockElasticRepository.
type MockElasticRepositoryMockRecorder struct {
	mock *MockElasticRepository
}

// NewMockElasticRepository creates a new mock instance.
func NewMockElasticRepository(ctrl *gomock.Controller) *MockElasticRepository {
	mock := &MockElasticRepository{ctrl: ctrl}
	mock.recorder = &MockElasticRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockElasticRepository) EXPECT() *MockElasticRepositoryMockRecorder {
	return m.recorder
}

// CreateIndexIfNotExists mocks base method.
func (m *MockElasticRepository) CreateIndexIfNotExists(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIndexIfNotExists", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIndexIfNotExists indicates an expected call of CreateIndexIfNotExists.
func (mr *MockElasticRepositoryMockRecorder) CreateIndexIfNotExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIndexIfNotExists", reflect.TypeOf((*MockElasticRepository)(nil).CreateIndexIfNotExists), arg0)
}

// GetReadmodel mocks base method.
func (m *MockElasticRepository) GetReadmodel(arg0 string) (database.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadmodel", arg0)
	ret0, _ := ret[0].(database.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReadmodel indicates an expected call of GetReadmodel.
func (mr *MockElasticRepositoryMockRecorder) GetReadmodel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadmodel", reflect.TypeOf((*MockElasticRepository)(nil).GetReadmodel), arg0)
}

// IsClientReady mocks base method.
func (m *MockElasticRepository) IsClientReady(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClientReady", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsClientReady indicates an expected call of IsClientReady.
func (mr *MockElasticRepositoryMockRecorder) IsClientReady(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClientReady", reflect.TypeOf((*MockElasticRepository)(nil).IsClientReady), arg0)
}

// LoadEvents mocks base method.
func (m *MockElasticRepository) LoadEvents(arg0 string) ([]database.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEvents", arg0)
	ret0, _ := ret[0].([]database.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadEvents indicates an expected call of LoadEvents.
func (mr *MockElasticRepositoryMockRecorder) LoadEvents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEvents", reflect.TypeOf((*MockElasticRepository)(nil).LoadEvents), arg0)
}

// SetUpIndexes mocks base method.
func (m *MockElasticRepository) SetUpIndexes() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUpIndexes")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUpIndexes indicates an expected call of SetUpIndexes.
func (mr *MockElasticRepositoryMockRecorder) SetUpIndexes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpIndexes", reflect.TypeOf((*MockElasticRepository)(nil).SetUpIndexes))
}

// StoreEvent mocks base method.
func (m *MockElasticRepository) StoreEvent(event database.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreEvent", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreEvent indicates an expected call of StoreEvent.
func (mr *MockElasticRepositoryMockRecorder) StoreEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreEvent", reflect.TypeOf((*MockElasticRepository)(nil).StoreEvent), event)
}

// StoreReadmodel mocks base method.
func (m *MockElasticRepository) StoreReadmodel(arg0 database.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreReadmodel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreReadmodel indicates an expected call of StoreReadmodel.
func (mr *MockElasticRepositoryMockRecorder) StoreReadmodel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreReadmodel", reflect.TypeOf((*MockElasticRepository)(nil).StoreReadmodel), arg0)
}

// UpdateReadmodel mocks base method.
func (m *MockElasticRepository) UpdateReadmodel(arg0 string, arg1 database.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReadmodel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReadmodel indicates an expected call of UpdateReadmodel.
func (mr *MockElasticRepositoryMockRecorder) UpdateReadmodel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReadmodel", reflect.TypeOf((*MockElasticRepository)(nil).UpdateReadmodel), arg0, arg1)
}

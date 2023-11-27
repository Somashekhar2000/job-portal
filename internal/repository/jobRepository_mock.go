// Code generated by MockGen. DO NOT EDIT.
// Source: jobRepository.go
//
// Generated by this command:
//
//	mockgen -source=jobRepository.go -destination=jobRepository_mock.go -package=repository
//
// Package repository is a generated GoMock package.
package repository

import (
	model "job-portal-api/internal/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockJobRepository is a mock of JobRepository interface.
type MockJobRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJobRepositoryMockRecorder
}

// MockJobRepositoryMockRecorder is the mock recorder for MockJobRepository.
type MockJobRepositoryMockRecorder struct {
	mock *MockJobRepository
}

// NewMockJobRepository creates a new mock instance.
func NewMockJobRepository(ctrl *gomock.Controller) *MockJobRepository {
	mock := &MockJobRepository{ctrl: ctrl}
	mock.recorder = &MockJobRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobRepository) EXPECT() *MockJobRepositoryMockRecorder {
	return m.recorder
}

// CreateJob mocks base method.
func (m *MockJobRepository) CreateJob(jodData model.Job) (model.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", jodData)
	ret0, _ := ret[0].(model.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob.
func (mr *MockJobRepositoryMockRecorder) CreateJob(jodData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockJobRepository)(nil).CreateJob), jodData)
}

// GetAllJobs mocks base method.
func (m *MockJobRepository) GetAllJobs() ([]model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllJobs")
	ret0, _ := ret[0].([]model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllJobs indicates an expected call of GetAllJobs.
func (mr *MockJobRepositoryMockRecorder) GetAllJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllJobs", reflect.TypeOf((*MockJobRepository)(nil).GetAllJobs))
}

// GetJobByCompanyID mocks base method.
func (m *MockJobRepository) GetJobByCompanyID(cID uint) ([]model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobByCompanyID", cID)
	ret0, _ := ret[0].([]model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobByCompanyID indicates an expected call of GetJobByCompanyID.
func (mr *MockJobRepositoryMockRecorder) GetJobByCompanyID(cID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobByCompanyID", reflect.TypeOf((*MockJobRepository)(nil).GetJobByCompanyID), cID)
}

// GetJobByJobID mocks base method.
func (m *MockJobRepository) GetJobByJobID(cID uint) (model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobByJobID", cID)
	ret0, _ := ret[0].(model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobByJobID indicates an expected call of GetJobByJobID.
func (mr *MockJobRepositoryMockRecorder) GetJobByJobID(cID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobByJobID", reflect.TypeOf((*MockJobRepository)(nil).GetJobByJobID), cID)
}
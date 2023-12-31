// Code generated by MockGen. DO NOT EDIT.
// Source: repository/cache_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	repository "CurrencyExchangeService/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExchangeRateCacheRepository is a mock of ExchangeRateCacheRepository interface.
type MockExchangeRateCacheRepository struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeRateCacheRepositoryMockRecorder
}

// MockExchangeRateCacheRepositoryMockRecorder is the mock recorder for MockExchangeRateCacheRepository.
type MockExchangeRateCacheRepositoryMockRecorder struct {
	mock *MockExchangeRateCacheRepository
}

// NewMockExchangeRateCacheRepository creates a new mock instance.
func NewMockExchangeRateCacheRepository(ctrl *gomock.Controller) *MockExchangeRateCacheRepository {
	mock := &MockExchangeRateCacheRepository{ctrl: ctrl}
	mock.recorder = &MockExchangeRateCacheRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangeRateCacheRepository) EXPECT() *MockExchangeRateCacheRepositoryMockRecorder {
	return m.recorder
}

// AddRates mocks base method.
func (m *MockExchangeRateCacheRepository) AddRates(arg0 repository.ExchangeRate) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRates", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// AddRates indicates an expected call of AddRates.
func (mr *MockExchangeRateCacheRepositoryMockRecorder) AddRates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRates", reflect.TypeOf((*MockExchangeRateCacheRepository)(nil).AddRates), arg0)
}

// GetLatestRates mocks base method.
func (m *MockExchangeRateCacheRepository) GetLatestRates() (repository.ExchangeRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestRates")
	ret0, _ := ret[0].(repository.ExchangeRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestRates indicates an expected call of GetLatestRates.
func (mr *MockExchangeRateCacheRepositoryMockRecorder) GetLatestRates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestRates", reflect.TypeOf((*MockExchangeRateCacheRepository)(nil).GetLatestRates))
}

// StopCache mocks base method.
func (m *MockExchangeRateCacheRepository) StopCache() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopCache")
}

// StopCache indicates an expected call of StopCache.
func (mr *MockExchangeRateCacheRepositoryMockRecorder) StopCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopCache", reflect.TypeOf((*MockExchangeRateCacheRepository)(nil).StopCache))
}

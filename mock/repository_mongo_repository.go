// Code generated by MockGen. DO NOT EDIT.
// Source: repository/mongo_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	repository "CurrencyExchangeService/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExchangeRateRepository is a mock of ExchangeRateRepository interface.
type MockExchangeRateRepository struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeRateRepositoryMockRecorder
}

// MockExchangeRateRepositoryMockRecorder is the mock recorder for MockExchangeRateRepository.
type MockExchangeRateRepositoryMockRecorder struct {
	mock *MockExchangeRateRepository
}

// NewMockExchangeRateRepository creates a new mock instance.
func NewMockExchangeRateRepository(ctrl *gomock.Controller) *MockExchangeRateRepository {
	mock := &MockExchangeRateRepository{ctrl: ctrl}
	mock.recorder = &MockExchangeRateRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangeRateRepository) EXPECT() *MockExchangeRateRepositoryMockRecorder {
	return m.recorder
}

// AddRates mocks base method.
func (m *MockExchangeRateRepository) AddRates(arg0 repository.ExchangeRate) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRates", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// AddRates indicates an expected call of AddRates.
func (mr *MockExchangeRateRepositoryMockRecorder) AddRates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRates", reflect.TypeOf((*MockExchangeRateRepository)(nil).AddRates), arg0)
}

// GetLatestRates mocks base method.
func (m *MockExchangeRateRepository) GetLatestRates() repository.ExchangeRate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestRates")
	ret0, _ := ret[0].(repository.ExchangeRate)
	return ret0
}

// GetLatestRates indicates an expected call of GetLatestRates.
func (mr *MockExchangeRateRepositoryMockRecorder) GetLatestRates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestRates", reflect.TypeOf((*MockExchangeRateRepository)(nil).GetLatestRates))
}

// GetRangeRates mocks base method.
func (m *MockExchangeRateRepository) GetRangeRates(arg0 repository.RangeRateRequest) []repository.ExchangeRate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRangeRates", arg0)
	ret0, _ := ret[0].([]repository.ExchangeRate)
	return ret0
}

// GetRangeRates indicates an expected call of GetRangeRates.
func (mr *MockExchangeRateRepositoryMockRecorder) GetRangeRates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRangeRates", reflect.TypeOf((*MockExchangeRateRepository)(nil).GetRangeRates), arg0)
}

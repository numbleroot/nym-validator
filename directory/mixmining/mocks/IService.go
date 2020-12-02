// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/nymtech/nym/validator/nym/directory/models"
	mock "github.com/stretchr/testify/mock"
)

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// BatchCreateMixStatus provides a mock function with given fields: batchMixStatus
func (_m *IService) BatchCreateMixStatus(batchMixStatus models.BatchMixStatus) []models.PersistedMixStatus {
	ret := _m.Called(batchMixStatus)

	var r0 []models.PersistedMixStatus
	if rf, ok := ret.Get(0).(func(models.BatchMixStatus) []models.PersistedMixStatus); ok {
		r0 = rf(batchMixStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PersistedMixStatus)
		}
	}

	return r0
}

// BatchGetMixStatusReport provides a mock function with given fields:
func (_m *IService) BatchGetMixStatusReport() models.BatchMixStatusReport {
	ret := _m.Called()

	var r0 models.BatchMixStatusReport
	if rf, ok := ret.Get(0).(func() models.BatchMixStatusReport); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.BatchMixStatusReport)
	}

	return r0
}

// CheckForDuplicateIP provides a mock function with given fields: host
func (_m *IService) CheckForDuplicateIP(host string) bool {
	ret := _m.Called(host)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(host)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CreateMixStatus provides a mock function with given fields: mixStatus
func (_m *IService) CreateMixStatus(mixStatus models.MixStatus) models.PersistedMixStatus {
	ret := _m.Called(mixStatus)

	var r0 models.PersistedMixStatus
	if rf, ok := ret.Get(0).(func(models.MixStatus) models.PersistedMixStatus); ok {
		r0 = rf(mixStatus)
	} else {
		r0 = ret.Get(0).(models.PersistedMixStatus)
	}

	return r0
}

// GatewayCount provides a mock function with given fields:
func (_m *IService) GatewayCount() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetActiveTopology provides a mock function with given fields:
func (_m *IService) GetActiveTopology() models.Topology {
	ret := _m.Called()

	var r0 models.Topology
	if rf, ok := ret.Get(0).(func() models.Topology); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Topology)
	}

	return r0
}

// GetRemovedTopology provides a mock function with given fields:
func (_m *IService) GetRemovedTopology() models.Topology {
	ret := _m.Called()

	var r0 models.Topology
	if rf, ok := ret.Get(0).(func() models.Topology); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Topology)
	}

	return r0
}

// GetStatusReport provides a mock function with given fields: pubkey
func (_m *IService) GetStatusReport(pubkey string) models.MixStatusReport {
	ret := _m.Called(pubkey)

	var r0 models.MixStatusReport
	if rf, ok := ret.Get(0).(func(string) models.MixStatusReport); ok {
		r0 = rf(pubkey)
	} else {
		r0 = ret.Get(0).(models.MixStatusReport)
	}

	return r0
}

// GetTopology provides a mock function with given fields:
func (_m *IService) GetTopology() models.Topology {
	ret := _m.Called()

	var r0 models.Topology
	if rf, ok := ret.Get(0).(func() models.Topology); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Topology)
	}

	return r0
}

// ListMixStatus provides a mock function with given fields: pubkey
func (_m *IService) ListMixStatus(pubkey string) []models.PersistedMixStatus {
	ret := _m.Called(pubkey)

	var r0 []models.PersistedMixStatus
	if rf, ok := ret.Get(0).(func(string) []models.PersistedMixStatus); ok {
		r0 = rf(pubkey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PersistedMixStatus)
		}
	}

	return r0
}

// MixCount provides a mock function with given fields:
func (_m *IService) MixCount() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// RegisterGateway provides a mock function with given fields: info
func (_m *IService) RegisterGateway(info models.GatewayRegistrationInfo) {
	_m.Called(info)
}

// RegisterMix provides a mock function with given fields: info
func (_m *IService) RegisterMix(info models.MixRegistrationInfo) {
	_m.Called(info)
}

// SaveBatchStatusReport provides a mock function with given fields: status
func (_m *IService) SaveBatchStatusReport(status []models.PersistedMixStatus) models.BatchMixStatusReport {
	ret := _m.Called(status)

	var r0 models.BatchMixStatusReport
	if rf, ok := ret.Get(0).(func([]models.PersistedMixStatus) models.BatchMixStatusReport); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Get(0).(models.BatchMixStatusReport)
	}

	return r0
}

// SaveStatusReport provides a mock function with given fields: status
func (_m *IService) SaveStatusReport(status models.PersistedMixStatus) models.MixStatusReport {
	ret := _m.Called(status)

	var r0 models.MixStatusReport
	if rf, ok := ret.Get(0).(func(models.PersistedMixStatus) models.MixStatusReport); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Get(0).(models.MixStatusReport)
	}

	return r0
}

// SetReputation provides a mock function with given fields: id, newRep
func (_m *IService) SetReputation(id string, newRep int64) bool {
	ret := _m.Called(id, newRep)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, int64) bool); ok {
		r0 = rf(id, newRep)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// StartupPurge provides a mock function with given fields:
func (_m *IService) StartupPurge() {
	_m.Called()
}

// UnregisterNode provides a mock function with given fields: id
func (_m *IService) UnregisterNode(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

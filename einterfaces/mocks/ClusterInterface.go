// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	einterfaces "github.com/mattermost/mattermost-server/v6/einterfaces"
	mock "github.com/stretchr/testify/mock"

	model "github.com/mattermost/mattermost-server/v6/model"
)

// ClusterInterface is an autogenerated mock type for the ClusterInterface type
type ClusterInterface struct {
	mock.Mock
}

// ConfigChanged provides a mock function with given fields: previousConfig, newConfig, sendToOtherServer
func (_m *ClusterInterface) ConfigChanged(previousConfig *model.Config, newConfig *model.Config, sendToOtherServer bool) *model.AppError {
	ret := _m.Called(previousConfig, newConfig, sendToOtherServer)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Config, *model.Config, bool) *model.AppError); ok {
		r0 = rf(previousConfig, newConfig, sendToOtherServer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// GetClusterId provides a mock function with given fields:
func (_m *ClusterInterface) GetClusterId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetClusterInfos provides a mock function with given fields:
func (_m *ClusterInterface) GetClusterInfos() []*model.ClusterInfo {
	ret := _m.Called()

	var r0 []*model.ClusterInfo
	if rf, ok := ret.Get(0).(func() []*model.ClusterInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ClusterInfo)
		}
	}

	return r0
}

// GetClusterStats provides a mock function with given fields:
func (_m *ClusterInterface) GetClusterStats() ([]*model.ClusterStats, *model.AppError) {
	ret := _m.Called()

	var r0 []*model.ClusterStats
	if rf, ok := ret.Get(0).(func() []*model.ClusterStats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ClusterStats)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func() *model.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: page, perPage
func (_m *ClusterInterface) GetLogs(page int, perPage int) ([]string, *model.AppError) {
	ret := _m.Called(page, perPage)

	var r0 []string
	if rf, ok := ret.Get(0).(func(int, int) []string); ok {
		r0 = rf(page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(int, int) *model.AppError); ok {
		r1 = rf(page, perPage)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetMyClusterInfo provides a mock function with given fields:
func (_m *ClusterInterface) GetMyClusterInfo() *model.ClusterInfo {
	ret := _m.Called()

	var r0 *model.ClusterInfo
	if rf, ok := ret.Get(0).(func() *model.ClusterInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ClusterInfo)
		}
	}

	return r0
}

// GetPluginStatuses provides a mock function with given fields:
func (_m *ClusterInterface) GetPluginStatuses() (model.PluginStatuses, *model.AppError) {
	ret := _m.Called()

	var r0 model.PluginStatuses
	if rf, ok := ret.Get(0).(func() model.PluginStatuses); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PluginStatuses)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func() *model.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// HealthScore provides a mock function with given fields:
func (_m *ClusterInterface) HealthScore() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// IsLeader provides a mock function with given fields:
func (_m *ClusterInterface) IsLeader() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NotifyMsg provides a mock function with given fields: buf
func (_m *ClusterInterface) NotifyMsg(buf []byte) {
	_m.Called(buf)
}

// QueryLogs provides a mock function with given fields: page, perPage
func (_m *ClusterInterface) QueryLogs(page int, perPage int) (map[string][]string, *model.AppError) {
	ret := _m.Called(page, perPage)

	var r0 map[string][]string
	if rf, ok := ret.Get(0).(func(int, int) map[string][]string); ok {
		r0 = rf(page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(int, int) *model.AppError); ok {
		r1 = rf(page, perPage)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// RegisterClusterMessageHandler provides a mock function with given fields: event, crm
func (_m *ClusterInterface) RegisterClusterMessageHandler(event model.ClusterEvent, crm einterfaces.ClusterMessageHandler) {
	_m.Called(event, crm)
}

// SendClusterMessage provides a mock function with given fields: msg
func (_m *ClusterInterface) SendClusterMessage(msg *model.ClusterMessage) {
	_m.Called(msg)
}

// SendClusterMessageToNode provides a mock function with given fields: nodeID, msg
func (_m *ClusterInterface) SendClusterMessageToNode(nodeID string, msg *model.ClusterMessage) error {
	ret := _m.Called(nodeID, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *model.ClusterMessage) error); ok {
		r0 = rf(nodeID, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartInterNodeCommunication provides a mock function with given fields:
func (_m *ClusterInterface) StartInterNodeCommunication() {
	_m.Called()
}

// StopInterNodeCommunication provides a mock function with given fields:
func (_m *ClusterInterface) StopInterNodeCommunication() {
	_m.Called()
}

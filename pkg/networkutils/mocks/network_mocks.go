// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-cni-k8s/pkg/networkutils (interfaces: NetworkAPIs)

// Package mock_networkutils is a generated GoMock package.
package mock_networkutils

import (
	net "net"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	netlink "github.com/vishvananda/netlink"
)

// MockNetworkAPIs is a mock of NetworkAPIs interface
type MockNetworkAPIs struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkAPIsMockRecorder
}

// MockNetworkAPIsMockRecorder is the mock recorder for MockNetworkAPIs
type MockNetworkAPIsMockRecorder struct {
	mock *MockNetworkAPIs
}

// NewMockNetworkAPIs creates a new mock instance
func NewMockNetworkAPIs(ctrl *gomock.Controller) *MockNetworkAPIs {
	mock := &MockNetworkAPIs{ctrl: ctrl}
	mock.recorder = &MockNetworkAPIsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkAPIs) EXPECT() *MockNetworkAPIsMockRecorder {
	return m.recorder
}

// DeleteRuleListBySrc mocks base method
func (m *MockNetworkAPIs) DeleteRuleListBySrc(arg0 net.IPNet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRuleListBySrc", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRuleListBySrc indicates an expected call of DeleteRuleListBySrc
func (mr *MockNetworkAPIsMockRecorder) DeleteRuleListBySrc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRuleListBySrc", reflect.TypeOf((*MockNetworkAPIs)(nil).DeleteRuleListBySrc), arg0)
}

// GetExcludeSNATCIDRs mocks base method
func (m *MockNetworkAPIs) GetExcludeSNATCIDRs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExcludeSNATCIDRs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetExcludeSNATCIDRs indicates an expected call of GetExcludeSNATCIDRs
func (mr *MockNetworkAPIsMockRecorder) GetExcludeSNATCIDRs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExcludeSNATCIDRs", reflect.TypeOf((*MockNetworkAPIs)(nil).GetExcludeSNATCIDRs))
}

// GetLinkByMac mocks base method
func (m *MockNetworkAPIs) GetLinkByMac(arg0 string, arg1 time.Duration) (netlink.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLinkByMac", arg0, arg1)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLinkByMac indicates an expected call of GetLinkByMac
func (mr *MockNetworkAPIsMockRecorder) GetLinkByMac(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLinkByMac", reflect.TypeOf((*MockNetworkAPIs)(nil).GetLinkByMac), arg0, arg1)
}

// GetRuleList mocks base method
func (m *MockNetworkAPIs) GetRuleList() ([]netlink.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRuleList")
	ret0, _ := ret[0].([]netlink.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRuleList indicates an expected call of GetRuleList
func (mr *MockNetworkAPIsMockRecorder) GetRuleList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuleList", reflect.TypeOf((*MockNetworkAPIs)(nil).GetRuleList))
}

// GetRuleListBySrc mocks base method
func (m *MockNetworkAPIs) GetRuleListBySrc(arg0 []netlink.Rule, arg1 net.IPNet) ([]netlink.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRuleListBySrc", arg0, arg1)
	ret0, _ := ret[0].([]netlink.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRuleListBySrc indicates an expected call of GetRuleListBySrc
func (mr *MockNetworkAPIsMockRecorder) GetRuleListBySrc(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuleListBySrc", reflect.TypeOf((*MockNetworkAPIs)(nil).GetRuleListBySrc), arg0, arg1)
}

// SetupENINetwork mocks base method
func (m *MockNetworkAPIs) SetupENINetwork(arg0, arg1 string, arg2 int, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupENINetwork", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupENINetwork indicates an expected call of SetupENINetwork
func (mr *MockNetworkAPIsMockRecorder) SetupENINetwork(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupENINetwork", reflect.TypeOf((*MockNetworkAPIs)(nil).SetupENINetwork), arg0, arg1, arg2, arg3)
}

// SetupHostNetwork mocks base method
func (m *MockNetworkAPIs) SetupHostNetwork(arg0, arg1 []string, arg2 string, arg3 *net.IP, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupHostNetwork", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupHostNetwork indicates an expected call of SetupHostNetwork
func (mr *MockNetworkAPIsMockRecorder) SetupHostNetwork(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupHostNetwork", reflect.TypeOf((*MockNetworkAPIs)(nil).SetupHostNetwork), arg0, arg1, arg2, arg3, arg4)
}

// UpdateRuleListBySrc mocks base method
func (m *MockNetworkAPIs) UpdateRuleListBySrc(arg0 []netlink.Rule, arg1 net.IPNet, arg2 []string, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRuleListBySrc", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRuleListBySrc indicates an expected call of UpdateRuleListBySrc
func (mr *MockNetworkAPIsMockRecorder) UpdateRuleListBySrc(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRuleListBySrc", reflect.TypeOf((*MockNetworkAPIs)(nil).UpdateRuleListBySrc), arg0, arg1, arg2, arg3)
}

// UseExternalSNAT mocks base method
func (m *MockNetworkAPIs) UseExternalSNAT() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseExternalSNAT")
	ret0, _ := ret[0].(bool)
	return ret0
}

// UseExternalSNAT indicates an expected call of UseExternalSNAT
func (mr *MockNetworkAPIsMockRecorder) UseExternalSNAT() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseExternalSNAT", reflect.TypeOf((*MockNetworkAPIs)(nil).UseExternalSNAT))
}

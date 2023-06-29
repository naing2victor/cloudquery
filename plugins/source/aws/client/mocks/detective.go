// Code generated by MockGen. DO NOT EDIT.
// Source: detective.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	detective "github.com/aws/aws-sdk-go-v2/service/detective"
	gomock "github.com/golang/mock/gomock"
)

// MockDetectiveClient is a mock of DetectiveClient interface.
type MockDetectiveClient struct {
	ctrl     *gomock.Controller
	recorder *MockDetectiveClientMockRecorder
}

// MockDetectiveClientMockRecorder is the mock recorder for MockDetectiveClient.
type MockDetectiveClientMockRecorder struct {
	mock *MockDetectiveClient
}

// NewMockDetectiveClient creates a new mock instance.
func NewMockDetectiveClient(ctrl *gomock.Controller) *MockDetectiveClient {
	mock := &MockDetectiveClient{ctrl: ctrl}
	mock.recorder = &MockDetectiveClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDetectiveClient) EXPECT() *MockDetectiveClientMockRecorder {
	return m.recorder
}

// BatchGetGraphMemberDatasources mocks base method.
func (m *MockDetectiveClient) BatchGetGraphMemberDatasources(arg0 context.Context, arg1 *detective.BatchGetGraphMemberDatasourcesInput, arg2 ...func(*detective.Options)) (*detective.BatchGetGraphMemberDatasourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &detective.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to BatchGetGraphMemberDatasources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetGraphMemberDatasources", varargs...)
	ret0, _ := ret[0].(*detective.BatchGetGraphMemberDatasourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetGraphMemberDatasources indicates an expected call of BatchGetGraphMemberDatasources.
func (mr *MockDetectiveClientMockRecorder) BatchGetGraphMemberDatasources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetGraphMemberDatasources", reflect.TypeOf((*MockDetectiveClient)(nil).BatchGetGraphMemberDatasources), varargs...)
}

// BatchGetMembershipDatasources mocks base method.
func (m *MockDetectiveClient) BatchGetMembershipDatasources(arg0 context.Context, arg1 *detective.BatchGetMembershipDatasourcesInput, arg2 ...func(*detective.Options)) (*detective.BatchGetMembershipDatasourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &detective.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to BatchGetMembershipDatasources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetMembershipDatasources", varargs...)
	ret0, _ := ret[0].(*detective.BatchGetMembershipDatasourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetMembershipDatasources indicates an expected call of BatchGetMembershipDatasources.
func (mr *MockDetectiveClientMockRecorder) BatchGetMembershipDatasources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetMembershipDatasources", reflect.TypeOf((*MockDetectiveClient)(nil).BatchGetMembershipDatasources), varargs...)
}

// DescribeOrganizationConfiguration mocks base method.
func (m *MockDetectiveClient) DescribeOrganizationConfiguration(arg0 context.Context, arg1 *detective.DescribeOrganizationConfigurationInput, arg2 ...func(*detective.Options)) (*detective.DescribeOrganizationConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &detective.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeOrganizationConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeOrganizationConfiguration", varargs...)
	ret0, _ := ret[0].(*detective.DescribeOrganizationConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeOrganizationConfiguration indicates an expected call of DescribeOrganizationConfiguration.
func (mr *MockDetectiveClientMockRecorder) DescribeOrganizationConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeOrganizationConfiguration", reflect.TypeOf((*MockDetectiveClient)(nil).DescribeOrganizationConfiguration), varargs...)
}

// GetMembers mocks base method.
func (m *MockDetectiveClient) GetMembers(arg0 context.Context, arg1 *detective.GetMembersInput, arg2 ...func(*detective.Options)) (*detective.GetMembersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &detective.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetMembers")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMembers", varargs...)
	ret0, _ := ret[0].(*detective.GetMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembers indicates an expected call of GetMembers.
func (mr *MockDetectiveClientMockRecorder) GetMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembers", reflect.TypeOf((*MockDetectiveClient)(nil).GetMembers), varargs...)
}

// ListDatasourcePackages mocks base method.
func (m *MockDetectiveClient) ListDatasourcePackages(arg0 context.Context, arg1 *detective.ListDatasourcePackagesInput, arg2 ...func(*detective.Options)) (*detective.ListDatasourcePackagesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &detective.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListDatasourcePackages")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDatasourcePackages", varargs...)
	ret0, _ := ret[0].(*detective.ListDatasourcePackagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDatasourcePackages indicates an expected call of ListDatasourcePackages.
func (mr *MockDetectiveClientMockRecorder) ListDatasourcePackages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDatasourcePackages", reflect.TypeOf((*MockDetectiveClient)(nil).ListDatasourcePackages), varargs...)
}

// ListGraphs mocks base method.
func (m *MockDetectiveClient) ListGraphs(arg0 context.Context, arg1 *detective.ListGraphsInput, arg2 ...func(*detective.Options)) (*detective.ListGraphsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &detective.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListGraphs")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGraphs", varargs...)
	ret0, _ := ret[0].(*detective.ListGraphsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGraphs indicates an expected call of ListGraphs.
func (mr *MockDetectiveClientMockRecorder) ListGraphs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGraphs", reflect.TypeOf((*MockDetectiveClient)(nil).ListGraphs), varargs...)
}

// ListInvitations mocks base method.
func (m *MockDetectiveClient) ListInvitations(arg0 context.Context, arg1 *detective.ListInvitationsInput, arg2 ...func(*detective.Options)) (*detective.ListInvitationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &detective.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListInvitations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInvitations", varargs...)
	ret0, _ := ret[0].(*detective.ListInvitationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInvitations indicates an expected call of ListInvitations.
func (mr *MockDetectiveClientMockRecorder) ListInvitations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInvitations", reflect.TypeOf((*MockDetectiveClient)(nil).ListInvitations), varargs...)
}

// ListMembers mocks base method.
func (m *MockDetectiveClient) ListMembers(arg0 context.Context, arg1 *detective.ListMembersInput, arg2 ...func(*detective.Options)) (*detective.ListMembersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &detective.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListMembers")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMembers", varargs...)
	ret0, _ := ret[0].(*detective.ListMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMembers indicates an expected call of ListMembers.
func (mr *MockDetectiveClientMockRecorder) ListMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockDetectiveClient)(nil).ListMembers), varargs...)
}

// ListOrganizationAdminAccounts mocks base method.
func (m *MockDetectiveClient) ListOrganizationAdminAccounts(arg0 context.Context, arg1 *detective.ListOrganizationAdminAccountsInput, arg2 ...func(*detective.Options)) (*detective.ListOrganizationAdminAccountsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &detective.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListOrganizationAdminAccounts")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrganizationAdminAccounts", varargs...)
	ret0, _ := ret[0].(*detective.ListOrganizationAdminAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizationAdminAccounts indicates an expected call of ListOrganizationAdminAccounts.
func (mr *MockDetectiveClientMockRecorder) ListOrganizationAdminAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizationAdminAccounts", reflect.TypeOf((*MockDetectiveClient)(nil).ListOrganizationAdminAccounts), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockDetectiveClient) ListTagsForResource(arg0 context.Context, arg1 *detective.ListTagsForResourceInput, arg2 ...func(*detective.Options)) (*detective.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &detective.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTagsForResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*detective.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockDetectiveClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockDetectiveClient)(nil).ListTagsForResource), varargs...)
}

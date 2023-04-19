// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	todo "github.com/B-danik/SecondTopic/todo"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthorization) Create(user todo.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAuthorizationMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthorization)(nil).Create), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(accessToken string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), accessToken)
}

// MockBook is a mock of Book interface.
type MockBook struct {
	ctrl     *gomock.Controller
	recorder *MockBookMockRecorder
}

// MockBookMockRecorder is the mock recorder for MockBook.
type MockBookMockRecorder struct {
	mock *MockBook
}

// NewMockBook creates a new mock instance.
func NewMockBook(ctrl *gomock.Controller) *MockBook {
	mock := &MockBook{ctrl: ctrl}
	mock.recorder = &MockBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBook) EXPECT() *MockBookMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBook) Create(name string, price int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, price)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBookMockRecorder) Create(name, price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBook)(nil).Create), name, price)
}

// Delete mocks base method.
func (m *MockBook) Delete(ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBookMockRecorder) Delete(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBook)(nil).Delete), ID)
}

// Get mocks base method.
func (m *MockBook) Get(ID int) (todo.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ID)
	ret0, _ := ret[0].(todo.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBookMockRecorder) Get(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBook)(nil).Get), ID)
}

// GetList mocks base method.
func (m *MockBook) GetList() ([]todo.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList")
	ret0, _ := ret[0].([]todo.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockBookMockRecorder) GetList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockBook)(nil).GetList))
}

// MockRent is a mock of Rent interface.
type MockRent struct {
	ctrl     *gomock.Controller
	recorder *MockRentMockRecorder
}

// MockRentMockRecorder is the mock recorder for MockRent.
type MockRentMockRecorder struct {
	mock *MockRent
}

// NewMockRent creates a new mock instance.
func NewMockRent(ctrl *gomock.Controller) *MockRent {
	mock := &MockRent{ctrl: ctrl}
	mock.recorder = &MockRentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRent) EXPECT() *MockRentMockRecorder {
	return m.recorder
}

// CraeteRent mocks base method.
func (m *MockRent) CraeteRent(user_id, book_id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CraeteRent", user_id, book_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CraeteRent indicates an expected call of CraeteRent.
func (mr *MockRentMockRecorder) CraeteRent(user_id, book_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CraeteRent", reflect.TypeOf((*MockRent)(nil).CraeteRent), user_id, book_id)
}

// GetRent mocks base method.
func (m *MockRent) GetRent(id int) ([]todo.Rent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRent", id)
	ret0, _ := ret[0].([]todo.Rent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRent indicates an expected call of GetRent.
func (mr *MockRentMockRecorder) GetRent(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRent", reflect.TypeOf((*MockRent)(nil).GetRent), id)
}

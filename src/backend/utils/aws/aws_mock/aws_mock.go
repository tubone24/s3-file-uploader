// Code generated by MockGen. DO NOT EDIT.
// Source: s3.go
// Source: dynamodb.go

// Package mock_aws is a generated GoMock package.
package aws_mock

import (
	gomock "github.com/golang/mock/gomock"
	aws "github.com/tubone24/s3-file-uploader/src/backend/utils/aws"
	dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	reflect "reflect"
)

// MockS3 is a mock of S3 interface
type MockS3 struct {
	ctrl     *gomock.Controller
	recorder *MockS3MockRecorder
}

// MockS3MockRecorder is the mock recorder for MockS3
type MockS3MockRecorder struct {
	mock *MockS3
}

// NewMockS3 creates a new mock instance
func NewMockS3(ctrl *gomock.Controller) *MockS3 {
	mock := &MockS3{ctrl: ctrl}
	mock.recorder = &MockS3MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockS3) EXPECT() *MockS3MockRecorder {
	return m.recorder
}

// UploadFile mocks base method
func (m *MockS3) UploadFile(bucket, filePath, key, contentType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", bucket, filePath, key, contentType)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile
func (mr *MockS3MockRecorder) UploadFile(bucket, filePath, key, contentType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockS3)(nil).UploadFile), bucket, filePath, key, contentType)
}

// ListObject mocks base method
func (m *MockS3) ListObject(bucket, prefix string) ([]aws.S3FileObjectInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObject", bucket, prefix)
	ret0, _ := ret[0].([]aws.S3FileObjectInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObject indicates an expected call of ListObject
func (mr *MockS3MockRecorder) ListObject(bucket, prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObject", reflect.TypeOf((*MockS3)(nil).ListObject), bucket, prefix)
}

// DownloadFile mocks base method
func (m *MockS3) DownloadFile(bucket, key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", bucket, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadFile indicates an expected call of DownloadFile
func (mr *MockS3MockRecorder) DownloadFile(bucket, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockS3)(nil).DownloadFile), bucket, key)
}

// DeleteObject mocks base method
func (m *MockS3) DeleteObject(bucket, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", bucket, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject
func (mr *MockS3MockRecorder) DeleteObject(bucket, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockS3)(nil).DeleteObject), bucket, key)
}

// MockDynamoDB is a mock of DynamoDB interface
type MockDynamoDB struct {
	ctrl     *gomock.Controller
	recorder *MockDynamoDBMockRecorder
}

// MockDynamoDBMockRecorder is the mock recorder for MockDynamoDB
type MockDynamoDBMockRecorder struct {
	mock *MockDynamoDB
}

// NewMockDynamoDB creates a new mock instance
func NewMockDynamoDB(ctrl *gomock.Controller) *MockDynamoDB {
	mock := &MockDynamoDB{ctrl: ctrl}
	mock.recorder = &MockDynamoDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDynamoDB) EXPECT() *MockDynamoDBMockRecorder {
	return m.recorder
}

// Scan mocks base method
func (m *MockDynamoDB) Scan(arg0 string) ([]map[string]*dynamodb.AttributeValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", arg0)
	ret0, _ := ret[0].([]map[string]*dynamodb.AttributeValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scan indicates an expected call of Scan
func (mr *MockDynamoDBMockRecorder) Scan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockDynamoDB)(nil).Scan), arg0)
}

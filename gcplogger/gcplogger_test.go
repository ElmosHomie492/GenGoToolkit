package gcplogger

import (
	"cloud.google.com/go/logging"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockLoggerClient is a mock of LoggerClient
type MockLoggerClient struct {
	mock.Mock
}

func (m *MockLoggerClient) Logger(logID string) *logging.Logger {
	args := m.Called(logID)
	return args.Get(0).(*logging.Logger)
}

func (m *MockLoggerClient) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestLogger_Log(t *testing.T) {
	// Mock context and logger client
	ctx := context.Background()
	mockClient := new(MockLoggerClient)
	mockLogger := new(logging.Logger) // Mock Logger
	mockClient.On("Logger", mock.Anything).Return(mockLogger)

	logger, err := New(ctx, mockClient, "test-log")
	assert.NoError(t, err)
	assert.NotNil(t, logger)

	// Test logging
	//logger.Log(logging.Info, "Test log")
}

func TestLogger_NewError(t *testing.T) {
	ctx := context.Background()

	// Test initialization with nil client
	logger, err := New(ctx, nil, "test-log")
	assert.Error(t, err)
	assert.Nil(t, logger)
}

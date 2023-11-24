package gcplogger

import (
	"cloud.google.com/go/logging"
	"context"
	"fmt"
)

// LoggerClient defines the interface for a logging client
type LoggerClient interface {
	Logger(logID string) *logging.Logger
	Close() error
}

// Logger represents the GCP logger
type Logger struct {
	client LoggerClient
	logger *logging.Logger
}

func CreateLogger(ctx context.Context, projectID string) (*logging.Logger, error) {
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create logging client: %w", err)
	}

	return client.Logger("GenGoToolkit"), nil
}

// New initializes a new GCP logger
func New(ctx context.Context, client LoggerClient, logName string) (*Logger, error) {
	if client == nil {
		return nil, fmt.Errorf("client cannot be nil")
	}

	return &Logger{
		client: client,
		logger: client.Logger(logName),
	}, nil
}

// Log logs a message with the specified severity
func (l *Logger) Log(severity logging.Severity, message string) {
	l.logger.Log(logging.Entry{Severity: severity, Payload: message})
}

// Close closes the logger client
func (l *Logger) Close() error {
	return l.client.Close()
}

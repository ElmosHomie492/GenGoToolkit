package gcplogger

import (
	"cloud.google.com/go/logging"
	"context"
	"fmt"
)

// Logger represents the GCP logger
type Logger struct {
	client *logging.Client
	logger *logging.Logger
}

// New initializes a new GCP logger
func New(projectID, logName string) (*Logger, error) {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create logging client: %w", err)
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

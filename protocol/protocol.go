package protocol

import (
	"context"
)

// ConnectionProtocol interface
type ConnectionProtocol interface {
	Initialize(ctx context.Context) error
	SupportedCommands() ([]Command, error)

	Execute(command Command) error
}

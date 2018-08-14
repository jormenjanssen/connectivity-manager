package protocol

import (
	"context"
)

// Protocol interface
type Protocol interface {
	Initialize(ctx context.Context) error
	SupportedCommands() ([]Command, error)

	Execute(command Command)
}

package supplicant

import (
	"context"
	"fmt"

	"github.com/jormenjanssen/connectivity-manager/protocol"
)

// Protocol type for supplicant protocol
type Protocol struct {
}

// Initialize the supplicant protocol
func (p Protocol) Initialize(ctx context.Context) error {
	fmt.Printf("Initializing supplicant protocol")
	return nil
}

// SupportedCommands gets a list of supported commands
func (p Protocol) SupportedCommands() ([]protocol.Command, error) {
	cmdList := make([]protocol.Command, 0)
	return cmdList, nil
}

// Execute a command on the protocol
func (p Protocol) Execute(command protocol.Command) error {
	return nil
}

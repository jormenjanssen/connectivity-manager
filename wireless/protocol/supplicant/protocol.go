package supplicant

import (
	"context"

	"github.com/jormenjanssen/connectivity-manager/protocol"
)

// Initialize the supplicant protocol
func Initialize(ctx context.Context) error {
	return nil
}

// SupportedCommands gets a list of supported commands
func SupportedCommands() ([]protocol.Command, error) {
	cmdList := make([]protocol.Command, 0)
	return cmdList, nil
}

// Execute a command on the protocol
func Execute(command protocol.Command) {

}

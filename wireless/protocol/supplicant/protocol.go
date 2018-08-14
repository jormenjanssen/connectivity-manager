package supplicant

import (
	"context"

	"github.com/jormenjanssen/connectivity-manager/protocol"
)

// Protocol type for supplicant protocol
type Protocol struct {
}

// Initialize the supplicant protocol
func (protocol *Protocol) Initialize(ctx context.Context) error {
	return nil
}

// SupportedCommands gets a list of supported commands
func (protocol *Protocol) SupportedCommands() ([]protocol.Command, error) {
	cmdList := make([]protocol., 0)
	return cmdList, nil
}

// Execute a command on the protocol
func (protocol *Protocol) Execute(command protocol.Command) {

}

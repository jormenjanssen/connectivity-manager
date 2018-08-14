package supplicant

import (
	"context"
	"fmt"
	"time"

	"github.com/jormenjanssen/connectivity-manager/protocol"
)

// Protocol type for supplicant protocol
type Protocol struct {
}

// Initialize the supplicant protocol
func (p Protocol) Initialize(ctx context.Context) error {
	fmt.Printf("Initializing supplicant protocol")
	go p.initializeBackground(ctx)
	return nil
}

func (p Protocol) initializeBackground(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Stopping")
		default:
			{
				time.Sleep(1 * time.Second)
				fmt.Println("Running supplicant protocol connection")
			}
		}
	}
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

package supplicant

import "context"

func Initialize(ctx context.Context) error {
	return nil
}

func SupportedCommands() ([]Command, error) {
	return []Command, nil
}

func Execute(command Command) {

}

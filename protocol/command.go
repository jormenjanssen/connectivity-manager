package protocol

// Command interface
type Command interface {
	Name() string
	execute(Parameter string) CommandResult
}

// CommandResult interface
type CommandResult interface {
	Result() (interface{}, error)
}

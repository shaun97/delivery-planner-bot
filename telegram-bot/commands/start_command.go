package commands

type StartCommand struct {
	// Add dependencies if needed
}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (c *StartCommand) Execute(args []string) (string, error) {
	return "Welcome to the bot!", nil
}

package commands

type Command interface {
	Execute(args []string) (string, error)
}

package commands

var registry = make(map[string]Command)

func Register(name string, cmd Command) {
	registry[name] = cmd
}

func Get(name string) Command {
	return registry[name]
}

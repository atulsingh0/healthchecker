package models

type Commands struct {
	cmds []Command
}

type Command struct {
	command  string
	args     string
	timeout  int
	exitcode int
	output   string
	extra    string
}

func NewCommands(cmds []Command) *Commands {
	return &Commands{
		cmds: cmds,
	}
}

func NewCommand(command, args string, timeout, exitcode int, output, extra string) *Command {
	return &Command{
		command:  command,
		args:     args,
		timeout:  timeout,
		exitcode: exitcode,
		output:   output,
		extra:    extra,
	}
}

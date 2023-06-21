package data

type Input struct {
	cmds []Cmds
}

type Cmds struct {
	command  string
	args     string
	timeout  int
	exitcode int
	output   string
	extra    string
}

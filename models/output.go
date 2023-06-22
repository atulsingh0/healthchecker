package models

type Outputs struct {
	outs []Output
}

type Output struct {
	command  string
	exitcode int
	output   string
}

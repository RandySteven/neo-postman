package main

type (
	CliRuler interface {
		createModel()
		createRepository()
	}

	cli struct {
		command []string
	}
)

func NewCli(command []string) *cli {
	return &cli{
		command: command,
	}
}

var _ CliRuler = &cli{}

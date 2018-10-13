package rcon

import "github.com/multiplay/go-source"

type Command struct {
	Cmd string
	Args interface{}
}

type CommandBuilder interface {
	SetCmd(command string)
	SetArgs(args ...interface{})
	Command() *source.Cmd
}

type SourceCommandBuilder struct {
	Cmd string
	Args interface{}
}

func (c *SourceCommandBuilder) SetCmd(command string) {
	c.Cmd = command
}

func (c *SourceCommandBuilder) SetArgs(args ...interface{}) {
	c.Args = args
}

func (c *SourceCommandBuilder) Command() *source.Cmd {
	if c.Args == nil {
		return source.NewCmd(c.Cmd)
	}
	return source.NewCmd(c.Cmd).WithArgs(c.Args)
}
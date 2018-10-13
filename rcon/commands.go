package rcon

import (
	"fmt"
	"github.com/multiplay/go-source"
)

type Commander interface {
	BuildCommand(builder CommandBuilder) *source.Cmd
	executeCommand(client *source.Client) (string, error)
}

type Broadcast struct {
	Message string
}

func (b *Broadcast) BuildCommand(builder CommandBuilder) *source.Cmd {
	cmd := fmt.Sprintf("AdminBroadcast %s", b.Message)
	builder.SetCmd(cmd)
	return builder.Command()
}

func (b *Broadcast) executeCommand(c *source.Client) (string, error) {
	cmd := b.BuildCommand(&SourceCommandBuilder{})
	return c.ExecCmd(cmd)
}
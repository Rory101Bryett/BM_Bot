package rcon

import (
	"errors"
	"fmt"
	"github.com/multiplay/go-source"
)

type Commander interface {
	BuildCommand(builder CommandBuilder) (*source.Cmd, error)
	executeCommand(client *source.Client) (string, error)
}

type Broadcast struct {
	Message string
}

func (b *Broadcast) BuildCommand(builder CommandBuilder) (*source.Cmd, error) {
	if b.Message == "" {
		return nil, errors.New("attempting to broadcast empty message")
	}
	cmd := fmt.Sprintf("AdminBroadcast %s", b.Message)
	builder.SetCmd(cmd)
	return builder.Command(), nil
}

func (b *Broadcast) executeCommand(c *source.Client) (string, error) {
	if cmd, err := b.BuildCommand(&SourceCommandBuilder{}); err == nil {
		return c.ExecCmd(cmd)
	} else {
		return "", err
	}
}

type Kick struct {
	Name string
	Id string
	KickReason string
}

func (k *Kick) BuildCommand(builder CommandBuilder) (*source.Cmd, error) {
	var cmd string
	if k.Id != "" {
		cmd = fmt.Sprintf("AdminKick %s %s", k.Id, k.KickReason)
	} else if k.Name != "" {
		cmd = fmt.Sprintf("AdminKick %s %s", k.Name, k.KickReason)
	} else {
		return nil, errors.New("no name or id given")
	}
	builder.SetCmd(cmd)
	return builder.Command(), nil
}

func (k *Kick) executeCommand(c *source.Client) (string, error) {
	if cmd, err := k.BuildCommand(&SourceCommandBuilder{}); err == nil {
		return c.ExecCmd(cmd)
	} else {
		return "", err
	}
}
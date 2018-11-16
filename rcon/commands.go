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

type ListPlayers struct {}

func (l *ListPlayers) BuildCommand(builder CommandBuilder) (*source.Cmd, error) {
	cmd := "ListPlayers"
	builder.SetCmd(cmd)
	return builder.Command(), nil
}

func (l *ListPlayers) executeCommand(c *source.Client) (string, error) {
	if cmd, err := l.BuildCommand(&SourceCommandBuilder{}); err == nil {
		return c.ExecCmd(cmd)
	} else {
		return "", err
	}
}

type TeamChange struct {
	Name string
	Id string
}

func (t *TeamChange) BuildCommand(builder CommandBuilder) (*source.Cmd, error) {
	var cmd string
	if t.Id != "" {
		cmd = fmt.Sprintf("AdminForceTeamChange %s", t.Id)
	} else if t.Name != "" {
		cmd = fmt.Sprintf("AdminForceTeamChange %s", t.Name)
	} else {
		return nil, errors.New("no name or id given")
	}
	builder.SetCmd(cmd)
	return builder.Command(), nil
}

func (t *TeamChange) executeCommand(c *source.Client) (string, error) {
	if cmd, err := t.BuildCommand(&SourceCommandBuilder{}); err == nil {
		return c.ExecCmd(cmd)
	} else {
		return "", err
	}
}

type AllowInvalidPlacement struct {
	Enable bool
}

func (a *AllowInvalidPlacement) BuildCommand(builder CommandBuilder) (*source.Cmd, error) {
	var cmd string
	if a.Enable {
		cmd = fmt.Sprintf("AdminAlwaysValidPlacement 1")
	} else if !a.Enable {
		cmd = fmt.Sprintf("AdminAlwaysValidPlacement 0")
	} else {
		return nil, errors.New("must set bool")
	}
	builder.SetCmd(cmd)
	return builder.Command(), nil
}

func (a *AllowInvalidPlacement) executeCommand(c *source.Client) (string, error) {
	if cmd, err := a.BuildCommand(&SourceCommandBuilder{}); err == nil {
		return c.ExecCmd(cmd)
	} else {
		return "", err
	}
}

type MatchPause struct {
	Pause bool
}

func (m *MatchPause) BuildCommand(builder CommandBuilder) (*source.Cmd, error) {
	var cmd string
	if m.Pause {
		cmd = fmt.Sprintf("AdminPauseMatch")
	} else if !m.Pause {
		cmd = fmt.Sprintf("AdminUnpauseMatch")
	} else {
		return nil, errors.New("must set bool")
	}
	builder.SetCmd(cmd)
	return builder.Command(), nil
}

func (m *MatchPause) executeCommand(c *source.Client) (string, error) {
	if cmd, err := m.BuildCommand(&SourceCommandBuilder{}); err == nil {
		return c.ExecCmd(cmd)
	} else {
		return "", err
	}
}

type AllowAllKits struct {
	Enable bool
}

func (a *AllowAllKits) BuildCommand(builder CommandBuilder) (*source.Cmd, error) {
	var cmd string
	if a.Enable {
		cmd = fmt.Sprintf("AdminAllKitsAvailable 1")
	} else if !a.Enable {
		cmd = fmt.Sprintf("AdminAllKitsAvailable 0")
	} else {
		return nil, errors.New("must set bool")
	}
	builder.SetCmd(cmd)
	return builder.Command(), nil
}

func (a *AllowAllKits) executeCommand(c *source.Client) (string, error) {
	if cmd, err := a.BuildCommand(&SourceCommandBuilder{}); err == nil {
		return c.ExecCmd(cmd)
	} else {
		return "", err
	}
}
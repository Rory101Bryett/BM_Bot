package rcon

import (
	"github.com/Rory101Bryett/BM_Bot/config"
	"testing"
)

func TestAccRconConn(t *testing.T) {
	cfg := config.New()
	cfg.LoadConfig("../squadmin_bot.yml")
	c := New(cfg.Rcon)
	response := c.ExecuteCmd(&Broadcast{"Hello World"})
	t.Log(response)
}

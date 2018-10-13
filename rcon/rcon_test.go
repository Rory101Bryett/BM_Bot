package rcon

import (
	"github.com/Rory101Bryett/BM_Bot/config"
	"testing"
)

func TestAccRCONConn(t *testing.T) {
	cfg := config.New()
	cfg.LoadConfig("../squadmin_bot.yml")
	c := New(cfg.Rcon)
	response := c.ExecuteCmd(&Broadcast{"Hello World"})
	t.Log(response)
	response = c.ExecuteCmd(&Kick{Name: "Gunter", KickReason: "Testing"})
	t.Log(response)
}

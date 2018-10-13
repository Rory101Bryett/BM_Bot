package config

import "testing"

func TestLoadConfig(t *testing.T) {
	cfg := New()
	cfg.LoadConfig("../squadmin_bot.yml")
	if cfg.Rcon.Address == "" || cfg.Rcon.Password == "" {
		t.Fatal("Config was not loaded successfully")
	}
	t.Logf("Rcon Address is: '%s'", cfg.Rcon.Address)
	t.Logf("Rcon Password is: '%s'", cfg.Rcon.Password)
}
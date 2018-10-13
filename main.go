package main

import (
	"fmt"
	"github.com/Rory101Bryett/BM_Bot/config"
	"github.com/Rory101Bryett/BM_Bot/rcon"
)

func main() {
	cfg := config.New()
	cfg.LoadConfig("squadmin_bot.yml")
	c := rcon.New(cfg.Rcon)
	response := c.ExecuteCmd(&rcon.Broadcast{"Hello World"})
	fmt.Print(response)
}
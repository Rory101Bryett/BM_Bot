package main

import (
	"github.com/Rory101Bryett/BM_Bot/config"
	"github.com/Rory101Bryett/BM_Bot/discord"
	"github.com/Rory101Bryett/BM_Bot/rcon"
)

func main() {
	cfg := config.New()
	cfg.LoadConfig("squadmin_bot.yml")

	r := rcon.New(cfg.Rcon)
	//response := r.ExecuteCmd(&rcon.Broadcast{"Hello World"})

	discord.Start(cfg.Discord, r)

	//bm := battlemetrics.New(cfg.BattleMetrics)
	//bm.GetServerId()
	//bm.FindPlayer()
}
package discord

import (
	"fmt"
	"github.com/Rory101Bryett/BM_Bot/config"
	"github.com/Rory101Bryett/BM_Bot/rcon"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Bot struct {
	client *discordgo.Session
	rcon *rcon.RCON
}

var bot = &Bot{}

func Start(cfg *config.DiscordConfig, r *rcon.RCON) {
	token := getToken(cfg.Token)
	dg, err := discordgo.New(token)
	if err != nil {
		log.Fatalf("Could not connect to discord %v", err)
	}
	bot.client = dg
	bot.rcon = r
	bot.BotHandlers()
	bot.Listen()
	fmt.Print("Discord Bot is now Running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.client.Close()
}

func getToken(token string) string {
	if !strings.HasPrefix("Bot", token) {
		t := strings.Join([]string{"Bot", token}, " ")
		return t
	}
	return token
}

func (b *Bot) Listen() {
	err := b.client.Open()
	if err != nil {
		log.Fatalf("Could not open websocket: %s", err)
	}
}

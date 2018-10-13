package discord

import (
	"fmt"
	"github.com/Rory101Bryett/BM_Bot/rcon"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func (b *Bot) BotHandlers() {
	b.client.AddHandler(b.ready)
	b.client.AddHandler(b.broadcast)
	b.client.AddHandler(b.kick)
}

func (b *Bot) ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Squadmin")
}

func (b *Bot) broadcast(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "!broadcast") {
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		}

		msg := parseMessage(m.Content)

		s.ChannelMessageSend(c.ID, fmt.Sprintf("Broadcasting Message: %s", msg))
		response := b.rcon.ExecuteCmd(&rcon.Broadcast{Message: msg})
		s.ChannelMessageSend(c.ID, response)
	}
}

func (b *Bot) kick(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "!kick") {
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		}

		command := parseMessageToSlice(m.Content)
		player := strings.Join(command[:1], " ")
		reason := strings.Join(command[1:], " ")

		s.ChannelMessageSend(c.ID, fmt.Sprintf("Kicking Player '%s' for '%s'", player, reason))
		response := b.rcon.ExecuteCmd(&rcon.Kick{Name: player, KickReason: reason})
		s.ChannelMessageSend(c.ID, response)
	}
}

func parseMessage(message string) string {
	msg := strings.Split(message, " ")
	msg = msg[1:]
	newMsg := strings.Join(msg, " ")
	return newMsg
}

func parseMessageToSlice(message string) []string {
	msg := strings.Split(message, " ")
	msg = msg[1:]
	return msg
}
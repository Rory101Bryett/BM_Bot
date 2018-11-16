package discord

import (
	"github.com/Rory101Bryett/BM_Bot/rcon"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func (b *Bot) BotHandlers() {
	b.client.AddHandler(b.ready)
	b.client.AddHandler(b.broadcast)
	b.client.AddHandler(b.kick)
	b.client.AddHandler(b.listPlayers)
	b.client.AddHandler(b.forceTeamChange)
	b.client.AddHandler(b.enableValidPlacement)
	b.client.AddHandler(b.disableValidPlacement)
	b.client.AddHandler(b.pauseMatch)
	b.client.AddHandler(b.unpauseMatch)
	b.client.AddHandler(b.allowAllKits)
	b.client.AddHandler(b.disallowAllKits)
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

		response := b.rcon.ExecuteCmd(&rcon.Kick{Name: player, KickReason: reason})
		s.ChannelMessageSend(c.ID, response)
	}
}

func (b *Bot) listPlayers(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "!list_players") {
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		}
		response := b.rcon.ExecuteCmd(&rcon.ListPlayers{})
		s.ChannelMessageSend(c.ID, response)
	}
}

func (b *Bot) forceTeamChange(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "!team_change") {
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		}

		name := parseMessage(m.Content)

		response := b.rcon.ExecuteCmd(&rcon.TeamChange{Name: name})
		s.ChannelMessageSend(c.ID, response)
	}
}

func (b *Bot) enableValidPlacement(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "!enable_valid_placements") {
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		}

		response := b.rcon.ExecuteCmd(&rcon.AllowInvalidPlacement{Enable: true})
		s.ChannelMessageSend(c.ID, response)
	}
}

func (b *Bot) disableValidPlacement(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "!disable_valid_placements") {
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		}

		response := b.rcon.ExecuteCmd(&rcon.AllowInvalidPlacement{Enable: false})
		s.ChannelMessageSend(c.ID, response)
	}
}

func (b *Bot) pauseMatch(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "!pause_match") {
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		}

		response := b.rcon.ExecuteCmd(&rcon.MatchPause{Pause: true})
		s.ChannelMessageSend(c.ID, response)
	}
}

func (b *Bot) unpauseMatch(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "!unpause_match") {
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		}

		response := b.rcon.ExecuteCmd(&rcon.MatchPause{Pause: false})
		s.ChannelMessageSend(c.ID, response)
	}
}

func (b *Bot) allowAllKits(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "!allow_all_kits") {
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		}

		response := b.rcon.ExecuteCmd(&rcon.AllowAllKits{Enable: true})
		s.ChannelMessageSend(c.ID, response)
	}
}

func (b *Bot) disallowAllKits(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "!disallow_all_kits") {
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		}

		response := b.rcon.ExecuteCmd(&rcon.AllowAllKits{Enable: false})
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
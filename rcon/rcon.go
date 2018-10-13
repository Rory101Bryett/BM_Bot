package rcon

import (
	"github.com/Rory101Bryett/BM_Bot/config"
	"github.com/multiplay/go-source"
	"log"
)


type RCON struct {
	Client *source.Client
}

/*
Creates a new rcon client with Multipacket responses disabled
*/
func New(cfg *config.RconConfig) *RCON {
	rconClient, err := source.NewClient(cfg.Address, source.Password(cfg.Password), source.DisableMultiPacket())
	if err != nil {
		log.Fatalf("Couldn't create a new RCON client connection")
	}
	return &RCON{Client: rconClient}
}

func (r *RCON) ExecuteCmd(c Commander) string {
	resp, err := c.executeCommand(r.Client)
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
	return resp
}
package battlemetrics

import (
	"fmt"
	"github.com/Rory101Bryett/BM_Bot/config"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const battlemetrics_api = "https://api.battlemetrics.com"
const search_server_endpoint = "/servers?filter[game]=squad&filter[search]=BB"

var battlemetrics_server_id string

type Client struct {
	cfg *config.BattleMetricsConfig
	BaseURL *url.URL
	httpClient *http.Client
}

func New(cfg *config.BattleMetricsConfig) *Client {
	baseUrl, err := url.Parse(battlemetrics_api)
	if err != nil {
		log.Fatal("Could not parse url '%s'", battlemetrics_api)
	}
	client := &http.Client{}
	return &Client{cfg: cfg, BaseURL: baseUrl, httpClient: client}
}

type Servers struct {
	Server []Server `json:"data"`
}

type Server struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes ServerAttribute `json:"attributes"`
}

type ServerAttribute struct {
	Name string `json:"name"`
	Ip string `json:"ip"`
}

func (c *Client) GetServerId() error {
	rel, err := url.Parse(search_server_endpoint)
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	var servers Servers
	_, err = c.do(req, &servers)
	if err != nil {
		return err
	}

	for _, s := range servers.Server {
		if strings.Contains(s.Attributes.Name, c.cfg.ServerName) {
			battlemetrics_server_id = s.Id
			fmt.Printf("Blood Bound Server Id: '%s'", battlemetrics_server_id)
		}
	}
	return nil
}

type Players struct {
	Player []Player `json:"data"`
}

type Player struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes PlayerAttribute
}

type PlayerAttribute struct {
	Id string `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (c *Client) FindPlayer() {
	path := "/players/76561197991476636"
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		log.Fatalf("Unexpected Error: %s", err)
	}

	var players Players
	_, err = c.do(req, &players)
	for _, p := range players.Player {
		fmt.Println(p.Attributes.Name)
	}
}


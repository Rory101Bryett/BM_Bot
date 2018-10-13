package config

import (
"log"

"github.com/spf13/viper"
)

type Config struct {
	Rcon *RconConfig
	BattleMetrics *BattleMetricsConfig
	Discord *DiscordConfig
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadConfig(filePath string) {
	viper.SetConfigFile(filePath)
	viper.AddConfigPath(".") // Look for config in current working directory
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Could not load configurations: %v", err)
	}
	err = viper.Unmarshal(c)
	if err != nil {
		log.Fatalf("Unable to decode into struct %v", err)
	}
}

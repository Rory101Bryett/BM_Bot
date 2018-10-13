package config

type RconConfig struct {
	Address string `mapstructure:"address"`
	Password string `mapstructure:"password"`
}
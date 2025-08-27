package config

import (
	"encoding/json"
	"os"
)

type GuildConfig struct {
	WelcomeChannel string `json:"welcomeChannel"`
	LogChannel     string `json:"logChannel"`
}

type Config struct {
	Guilds map[string]GuildConfig `json:"guilds"`
}

var Cfg Config

func LoadConfig(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &Cfg)
	if err != nil {
		return err
	}
	return nil
}

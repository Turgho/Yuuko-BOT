package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type GuildConfig struct {
	ID             string `json:"ID"`
	WelcomeChannel string `json:"WelcomeChannel"`
	RulesMessageID string `json:"RulesMessageID"`
	RoleMemberID   string `json:"RoleMemberID"`
}

// ConfigStruct representa a estrutura do JSON (slice de guilds)
type ConfigStruct struct {
	Guilds []GuildConfig `json:"Guilds"`
}

// CfgMap é o map global: guildID -> GuildConfig
var CfgMap map[string]GuildConfig

// LoadConfig carrega o JSON e retorna um map
func LoadConfig(path string) (map[string]GuildConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler config: %w", err)
	}

	var cfg ConfigStruct
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, fmt.Errorf("erro ao fazer unmarshal do config.json: %w", err)
	}

	CfgMap = make(map[string]GuildConfig)
	for _, g := range cfg.Guilds {
		CfgMap[g.ID] = g
	}

	return CfgMap, nil
}

func SaveConfig(path string) error {
	// Cria a estrutura completa com slice
	cfg := ConfigStruct{
		Guilds: make([]GuildConfig, 0, len(CfgMap)),
	}

	for _, g := range CfgMap {
		cfg.Guilds = append(cfg.Guilds, g)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao serializar config: %w", err)
	}

	return os.WriteFile(path, data, 0644)
}

package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// GetGuildInfo tenta obter as informações do servidor a partir do estado cacheado,
// e caso não encontre, faz requisição direta à API.
// Retorna o objeto Guild, a URL do ícone (se houver) e erro, se ocorrer.
func GetGuildInfo(s *discordgo.Session, guildID string) (*discordgo.Guild, string, error) {
	guild, err := s.State.Guild(guildID)
	if err != nil {
		guild, err = s.Guild(guildID)
		if err != nil {
			return nil, "", err
		}
	}

	iconURL := ""
	if guild.Icon != "" {
		iconURL = fmt.Sprintf("https://cdn.discordapp.com/icons/%s/%s.png", guild.ID, guild.Icon)
	}

	return guild, iconURL, nil
}

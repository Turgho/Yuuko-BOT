package register

import (
	"Turgho/Yuuko-BOT/config"

	"github.com/bwmarrin/discordgo"
)

func RegisterAllCommands(s *discordgo.Session, appID string, guilds map[string]config.GuildConfig) {
	for guildID := range guilds {
		RegisterPublicCommands(s, appID, guildID)
		RegisterAdminCommands(s, appID, guildID)
		RegistesGamesCommands(s, appID, guildID)
	}
}

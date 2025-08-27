package register

import (
	"Turgho/Yuuko-BOT/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegisterAllCommands(s *discordgo.Session, appID string, guilds map[string]config.GuildConfig) {
	log.Printf(">> Guilds carregadas: %d\n", len(guilds))
	for guildID := range guilds {
		log.Printf("Registrando comandos para guild: %s\n", guildID)
		RegisterPublicCommands(s, appID, guildID)
		RegisterAdminCommands(s, appID, guildID)
		RegistesGamesCommands(s, appID, guildID)
	}
}

package register

import "github.com/bwmarrin/discordgo"

func RegisterAllCommands(s *discordgo.Session, appID, guildID string) {
	RegisterPublicCommands(s, appID, guildID)
	RegisterAdminCommands(s, appID, guildID)
}

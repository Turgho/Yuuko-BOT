package commands

import (
	"Turgho/Yuuko-BOT/commands/router"
	"Turgho/Yuuko-BOT/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix = "!"

func HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	content := m.Content
	if !strings.HasPrefix(content, prefix) {
		return
	}

	args := strings.Split(strings.TrimPrefix(content, prefix), " ")
	cmd := args[0]

	// Primeiro checa se é comando de admin
	if handler, ok := router.AdminCommands[cmd]; ok {
		if utils.IsAdmin(s, m.GuildID, m.Author.ID) {
			handler(s, m)
		} else {
			s.ChannelMessageSendReply(m.ChannelID, "❌ Você não tem permissão para esse comando.", &discordgo.MessageReference{
				MessageID: m.ID,
				ChannelID: m.ChannelID,
			})
		}
		return // para aqui se for admin
	}

	// Depois checa se é comando público
	if handler, ok := router.PublicCommands[cmd]; ok {
		handler(s, m)
		return
	}
}

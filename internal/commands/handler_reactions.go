package commands

import (
	"log"

	"Turgho/Yuuko-BOT/config"
	"Turgho/Yuuko-BOT/internal/services/logger"

	"github.com/bwmarrin/discordgo"
)

// HandleReactionAdd é chamado quando um usuário adiciona uma reação em uma mensagem.
func HandleReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// Ignora reações do próprio bot
	if r.UserID == s.State.User.ID {
		return
	}

	for _, guildCfg := range config.CfgMap {
		if r.MessageID == guildCfg.RulesMessageID && r.Emoji.Name == "✅" {
			logger.LogReaction(s, r) // log automático
			err := s.GuildMemberRoleAdd(r.GuildID, r.UserID, guildCfg.RoleMemberID)
			if err != nil {
				log.Printf("❌ Erro ao adicionar cargo ao usuário %s na guild %s: %v", r.UserID, r.GuildID, err)
				return
			}

			_, err = s.ChannelMessageSend(r.ChannelID, "<@"+r.UserID+"> leu e aceitou as regras! ✅")
			if err != nil {
				log.Printf("❌ Erro ao enviar mensagem de confirmação: %v", err)
			}
			return
		}
	}
}

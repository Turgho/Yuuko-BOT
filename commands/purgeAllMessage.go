package commands

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func PurgeAllCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	member, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Erro ao verificar permissões.")
		return
	}

	// Verificar se o autor é admin ou tem permissão de gerenciar mensagens
	isAdmin := false
	for _, roleID := range member.Roles {
		role, err := s.State.Role(m.GuildID, roleID)
		if err != nil {
			continue
		}
		if role.Permissions&(discordgo.PermissionAdministrator|discordgo.PermissionManageMessages) != 0 {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		s.ChannelMessageSend(m.ChannelID, "🚫 Você não tem permissão para usar este comando.")
		return
	}

	for {
		messages, err := s.ChannelMessages(m.ChannelID, 100, "", "", "")
		if err != nil || len(messages) == 0 {
			break
		}

		var deletable []string
		for _, msg := range messages {
			if time.Since(msg.Timestamp.Local()) < 14*24*time.Hour {
				deletable = append(deletable, msg.ID)
			} else {
				// Deleta individualmente se for mais antiga
				_ = s.ChannelMessageDelete(m.ChannelID, msg.ID)
			}
		}

		if len(deletable) > 1 {
			err = s.ChannelMessagesBulkDelete(m.ChannelID, deletable)
			if err != nil {
				log.Println("Erro ao apagar mensagens em lote:", err)
			}
		} else if len(deletable) == 1 {
			_ = s.ChannelMessageDelete(m.ChannelID, deletable[0])
		}
	}

	s.ChannelMessageSend(m.ChannelID, "🧹 Canal limpo com sucesso!")
}

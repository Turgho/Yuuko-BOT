package admin

import (
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

// PurgeSlashCommand remove todas as mensagens de um canal em blocos de até 100 mensagens.
// Mensagens com mais de 14 dias são excluídas individualmente, pois o Discord limita exclusões em massa.
func PurgeSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Confirmação visual rápida
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	totalDeleted := 0

	for {
		// Busca até 100 mensagens do canal
		messages, err := s.ChannelMessages(i.ChannelID, 100, "", "", "")
		if err != nil || len(messages) == 0 {
			break
		}

		var deletable []string
		for _, msg := range messages {
			// Ignora mensagens fixadas
			if msg.Pinned {
				continue
			}

			// Verifica se tem menos de 14 dias
			if time.Since(msg.Timestamp) < 14*24*time.Hour {
				deletable = append(deletable, msg.ID)
			} else {
				// Exclui individualmente
				_ = s.ChannelMessageDelete(i.ChannelID, msg.ID)
				totalDeleted++
			}
		}

		// Exclui em lote se possível
		if len(deletable) > 1 {
			err = s.ChannelMessagesBulkDelete(i.ChannelID, deletable)
			if err != nil {
				log.Println("Erro ao apagar mensagens em lote:", err)
			}
			totalDeleted += len(deletable)
		} else if len(deletable) == 1 { // Deleta 1 por 1
			_ = s.ChannelMessageDelete(i.ChannelID, deletable[0])
			totalDeleted++
		}
	}

	// Envia resposta final ao usuário
	s.ChannelMessageSend(i.ChannelID, "🧹 Foram apagadas **"+formatInt(totalDeleted)+"** mensagens com sucesso!")
}

// formatInt formata um int como string
func formatInt(n int) string {
	return strconv.Itoa(n)
}

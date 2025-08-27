package public

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func PingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	latency := s.HeartbeatLatency().Milliseconds()
	content := fmt.Sprintf("🏓 Pong!\nLatência: `%dms`", latency)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	if err != nil {
		log.Println("Erro ao enviar resposta do ping:", err)
		return
	}

	// Obter nome do servidor para log (guild pode ser nil se for DM)
	if i.GuildID != "" {
		guild, err := s.State.Guild(i.GuildID)
		if err != nil {
			guild, err = s.Guild(i.GuildID)
			if err != nil {
				log.Printf("Erro ao buscar guild: %v", err)
				return
			}
		}
		log.Printf("Usuário %s executou ping no servidor %s (ID: %s)", i.Member.User.Username, guild.Name, i.GuildID)
	}
}

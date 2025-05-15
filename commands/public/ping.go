package public

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	latency := s.HeartbeatLatency().Milliseconds()
	content := fmt.Sprintf("🏓 Pong!\nLatência: `%dms`", latency)

	// Envia a mensagem como resposta
	_, err := s.ChannelMessageSendReply(m.ChannelID, content, &discordgo.MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	})
	if err != nil {
		log.Println("Erro ao enviar resposta do ping:", err)
		return
	}

	// Tenta obter o nome do servidor
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		guild, err = s.Guild(m.GuildID) // fallback direto da API
		if err != nil {
			log.Printf("Erro ao buscar guild: %v", err)
			return
		}
	}

	log.Printf("Usuário %s executou ping no servidor %s (ID: %s)", m.Author.Username, guild.Name, m.GuildID)
}

package admin

import (
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ShutdownSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Responde à interação avisando que vai desligar
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "🛑 Desligando o bot...",
			Flags:   discordgo.MessageFlagsEphemeral, // mensagem visível só para quem executou
		},
	})
	if err != nil {
		// Se não conseguir responder, tenta enviar mensagem no canal
		s.ChannelMessageSend(i.ChannelID, "A mimir, patrão... 😴")
	}

	log.Println("Bot desligado às ", time.Now().Format(time.RFC3339))

	// Fecha a sessão para desconectar do gateway
	s.Close()

	// Finaliza o processo
	os.Exit(0)
}

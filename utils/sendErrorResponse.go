package utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// SendErrorResponse envia uma mensagem de erro padrão como resposta a uma interação
func SendErrorResponse(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
			Flags:   discordgo.MessageFlagsEphemeral, // mensagem visível só para o usuário
		},
	})
	if err != nil {
		log.Printf("Erro ao enviar resposta de erro: %v", err)
	}
}

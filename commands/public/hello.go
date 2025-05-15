package public

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func HelloSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Obtém o usuário que executou o comando
	var user *discordgo.User
	if i.Member != nil {
		user = i.Member.User
	} else {
		user = i.User
	}

	message := fmt.Sprintf("Olá, %s 👋", user.Username)

	// Envia a resposta da interação
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})

	if err != nil {
		log.Println("Erro ao enviar resposta do hello:", err)
	}
}

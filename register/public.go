package register

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegisterPublicCommands(s *discordgo.Session, appID, guildID string) {
	publicCommands := []*discordgo.ApplicationCommand{
		{
			Name:        "hello",
			Description: "Diz olá ao usuário",
		},
		{
			Name:        "ping",
			Description: "Comando para testar a latência de resposta do bot",
		},
	}

	for _, cmd := range publicCommands {
		_, err := s.ApplicationCommandCreate(appID, guildID, cmd)
		if err != nil {
			log.Printf("Erro ao registrar comando público /%s: %v", cmd.Name, err)
		}
	}
}

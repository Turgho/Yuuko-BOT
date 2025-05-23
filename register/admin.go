package register

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegisterAdminCommands(s *discordgo.Session, appID, guildID string) {
	adminCommands := []*discordgo.ApplicationCommand{
		{
			Name:        "rules",
			Description: "Mostra as regras do servidor",
		},
		{
			Name:        "purge",
			Description: "Limpa as mensagens do canal todo",
		},
		{
			Name:        "restart",
			Description: "Reinicia o bot",
		},
		{
			Name:        "shutdown",
			Description: "Desliga o bot",
		},
	}

	for _, cmd := range adminCommands {
		_, err := s.ApplicationCommandCreate(appID, guildID, cmd)
		if err != nil {
			log.Printf("Erro ao registrar comando admin /%s: %v", cmd.Name, err)
		}
	}
}

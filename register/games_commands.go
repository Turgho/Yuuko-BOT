package register

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegistesGamesCommands(s *discordgo.Session, appID, guildID string) {
	gamesCommands := []*discordgo.ApplicationCommand{
		{
			Name:        "coinflip",
			Description: "Gira uma moeda com cara ou coroa",
		},
	}

	for _, cmd := range gamesCommands {
		_, err := s.ApplicationCommandCreate(appID, guildID, cmd)
		if err != nil {
			log.Printf("Erro ao registrar comando de games /%s: %v", cmd.Name, err)
		}
	}
}

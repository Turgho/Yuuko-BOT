package register

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegistesGamesCommands(s *discordgo.Session, appID, guildID string) {
	gamesCommands := []*discordgo.ApplicationCommand{
		{
			Name:        "givepoints",
			Description: "Adiciona pontos a um usuário",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "usuario",
					Description: "Usuário para dar pontos",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "quantidade",
					Description: "Quantidade de pontos a adicionar",
					Required:    true,
				},
			},
		},
		{
			Name:        "getpoints",
			Description: "Busca a quantidade de pontos de um usuário",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "usuario",
					Description: "Usuário para buscar a quantidade de pontos",
					Required:    true,
				},
			},
		},
	}

	for _, cmd := range gamesCommands {
		_, err := s.ApplicationCommandCreate(appID, guildID, cmd)
		if err != nil {
			log.Printf("Erro ao registrar comando de games /%s: %v", cmd.Name, err)
		}
	}
}

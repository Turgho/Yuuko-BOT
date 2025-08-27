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
		{
			Name:        "kick",
			Description: "Expulsa um membro do servidor",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "Usuário para kickar",
					Required:    true,
				},
			},
		},
		{
			Name:        "setmember",
			Description: "Define cargo de membro do servidor",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "rolename",
					Description: "Cargo para adicionar",
					Required:    true,
				},
			},
		},
		{
			Name:        "setwelcome",
			Description: "Define o canal de boas-vindas da guilda",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel",
					Description: "Canal de boas-vindas",
					Required:    true,
				},
			},
		},
	}

	for _, cmd := range adminCommands {
		created, err := s.ApplicationCommandCreate(appID, guildID, cmd)

		if err != nil {
			log.Printf("❌ Erro ao registrar comando /%s: %v", cmd.Name, err)
		} else {
			log.Printf("✅ Comando /%s registrado! GuildID: %s | AppID: %s | CmdID: %s",
				created.Name, guildID, appID, created.ID)
		}
	}
}

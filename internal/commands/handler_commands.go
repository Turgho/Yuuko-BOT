package commands

import (
	"Turgho/Yuuko-BOT/internal/router"
	"Turgho/Yuuko-BOT/internal/services/logger"
	"Turgho/Yuuko-BOT/internal/services/utils"

	"github.com/bwmarrin/discordgo"
)

// HandleInteraction é o handler geral para todas as interações
func HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		handleSlashCommand(s, i)
	case discordgo.InteractionMessageComponent:
		handleButtonInteraction(s, i)
	}
}

// ------------------------ Handlers internos ------------------------

func handleSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	cmdName := i.ApplicationCommandData().Name

	respondEphemeral := func(content string) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}

	// Admin commands
	if handler, ok := router.AdminCommands[cmdName]; ok {
		if utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
			logger.LogCommand(s, i) // log automático
			handler(s, i)
		} else {
			respondEphemeral("❌ Você não tem permissão para esse comando.")
		}
		return
	}

	// Outras categorias (public/games)
	commandMaps := []map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
		router.PublicCommands,
		router.GamesCommands,
	}

	for _, m := range commandMaps {
		if handler, ok := m[cmdName]; ok {
			logger.LogCommand(s, i) // log automático
			handler(s, i)
			return
		}
	}

	// Comando não encontrado
	respondEphemeral("❌ Comando não conhecido.")
}

func handleButtonInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i == nil || i.Type != discordgo.InteractionMessageComponent {
		return
	}

	respondEphemeral(s, i, "Botão não configurado para este comando.")
}

// ------------------------ Funções utilitárias ------------------------

func respondEphemeral(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

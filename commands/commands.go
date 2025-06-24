package commands

import (
	"Turgho/Yuuko-BOT/commands/router"
	"Turgho/Yuuko-BOT/utils"

	"github.com/bwmarrin/discordgo"
)

// HandleInteraction é o handler geral para interações (slash commands, botões, etc).
func HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		HandleSlashCommand(s, i)

	case discordgo.InteractionMessageComponent:
		HandleButtonInteraction(s, i)

	// Outros tipos podem ser tratados aqui se quiser
	default:
		// Ignora outros tipos
	}
}

// HandleSlashCommand processa comandos do tipo slash command.
func HandleSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Só trata comandos do tipo ApplicationCommand
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	cmd := i.ApplicationCommandData().Name

	respondEphemeral := func(content string) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}

	if handler, ok := router.AdminCommands[cmd]; ok {
		if utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
			handler(s, i)
		} else {
			respondEphemeral("❌ Você não tem permissão para esse comando.")
		}
		return
	}

	commandCategories := []map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
		router.PublicCommands,
		router.MonsterHunterCommands,
		router.GamesCommands,
	}

	for _, category := range commandCategories {
		if handler, ok := category[cmd]; ok {
			handler(s, i)
			return
		}
	}

	respondEphemeral("❌ Comando não conhecido.")
}

func HandleButtonInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i == nil {
		return
	}

	if i.Type != discordgo.InteractionMessageComponent {
		// Ignora interações que não são de botão
		return
	}

	data := i.MessageComponentData()
	if data.CustomID == "" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Dados inválidos no componente.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}
	elements := []string{"Fire", "Water", "Thunder", "Ice", "Dragon", "Poison", "Sleep", "Paralysis", "Blast", "Stun"}

	customID := data.CustomID

	// Verifica se o customID é um dos elementos
	isElement := false
	for _, e := range elements {
		if customID == e {
			isElement = true
			break
		}
	}

	switch {
	case customID == "attack" || customID == "run":
		// monsterhunter.HandleMonsterHunterButtons(s, i)
	case isElement:
		// monsterhunter.HandleElementButton(s, i)
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Botão desconhecido.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}

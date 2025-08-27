package admin

import (
	"log"

	"Turgho/Yuuko-BOT/config"
	"Turgho/Yuuko-BOT/internal/services/logger"

	"github.com/bwmarrin/discordgo"
)

// SetWelcomeChannelSlashCommand define o canal de boas-vindas da guilda
func SetWelcomeChannelSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Você precisa fornecer um canal.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	channelID, ok := options[0].Value.(string)
	if !ok || channelID == "" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Canal inválido.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	// Atualiza WelcomeChannel no config
	guildCfg, ok := config.CfgMap[i.GuildID]
	if !ok {
		guildCfg = config.GuildConfig{ID: i.GuildID}
	}
	guildCfg.WelcomeChannel = channelID
	config.CfgMap[i.GuildID] = guildCfg

	// Salva config.json
	if err := config.SaveConfig("config/config.json"); err != nil {
		log.Println("Erro ao salvar config.json:", err)
	}

	// Resposta para o usuário
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "✅ Canal de boas-vindas definido com sucesso!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	// Log automático
	logger.LogCommand(s, i)
}

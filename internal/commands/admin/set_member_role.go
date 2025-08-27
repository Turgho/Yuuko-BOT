package admin

import (
	"log"

	"Turgho/Yuuko-BOT/config"

	"github.com/bwmarrin/discordgo"
)

// SetMemberRoleSlashCommand define o MemberRoleID da guilda
func SetMemberRoleSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Você precisa fornecer um cargo.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	roleID, ok := options[0].Value.(string)
	if !ok || roleID == "" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Cargo inválido.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	// Atualiza MemberRoleID no config
	guildCfg, ok := config.CfgMap[i.GuildID]
	if !ok {
		guildCfg = config.GuildConfig{ID: i.GuildID}
	}
	guildCfg.RoleMemberID = roleID
	config.CfgMap[i.GuildID] = guildCfg

	if err := config.SaveConfig("config/config.json"); err != nil {
		log.Println("Erro ao salvar config.json:", err)
	}

	// Resposta para o usuário
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "✅ Cargo de membro definido com sucesso!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

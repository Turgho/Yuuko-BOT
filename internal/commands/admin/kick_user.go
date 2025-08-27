package admin

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func KickUserSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Verifica se a opção "user" foi passada
	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Você precisa mencionar um usuário para kickar.",
				Flags:   discordgo.MessageFlagsEphemeral, // mensagem só para o autor
			},
		})
		return
	}

	// Pega o usuário alvo
	user := options[0].UserValue(s)

	// Tenta kickar
	err := s.GuildMemberDeleteWithReason(i.GuildID, user.ID, "Kick via comando")
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Erro ao kickar %s: %v", user.Username, err),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	// Responde com sucesso
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("👢 %s foi expulso com sucesso.", user.Username),
		},
	})
}

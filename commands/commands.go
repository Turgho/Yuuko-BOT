package commands

import (
	"Turgho/Yuuko-BOT/commands/router"
	"Turgho/Yuuko-BOT/utils"

	"github.com/bwmarrin/discordgo"
)

// HandleSlashCommand processa interações do tipo slash command recebidas pelo bot.
// Ele identifica o comando invocado, verifica permissões e executa o handler apropriado.
func HandleSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Obtém o nome do comando invocado pelo usuário
	cmd := i.ApplicationCommandData().Name

	// Primeiro, verifica se o comando está na lista de comandos de administrador
	if handler, ok := router.AdminCommands[cmd]; ok {
		// Verifica se o usuário que chamou o comando tem permissão de admin
		if utils.IsAdmin(s, i.GuildID, i.Member.User.ID) {
			// Executa o handler do comando para administradores
			handler(s, i)
		} else {
			// Caso o usuário não tenha permissão, envia uma mensagem efêmera informando
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "❌ Você não tem permissão para esse comando.",
					Flags:   discordgo.MessageFlagsEphemeral, // só o usuário vê a mensagem
				},
			})
		}
		return // Finaliza o fluxo, pois já encontrou o comando
	}

	// Se não for comando admin, verifica se é um comando público
	if handler, ok := router.PublicCommands[cmd]; ok {
		// Executa o handler do comando público
		handler(s, i)
		return
	}

	// Comando não reconhecido — responde informando
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "❌ Comando não conhecido.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

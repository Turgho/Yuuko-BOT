package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var RulesMessageID string

func RulesCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	rules := `
Bem-vindos ao nosso servidor! Para manter um ` + "`ambiente saudável e agradável para todos`" + `, é importante seguir as regras abaixo. Lembre-se, ` + "`o respeito mútuo é fundamental para garantir uma boa convivência`.\n\n" +

		`- ` + "`Regra 1`" + `
👐 ┆ **Respeito Mútuo**:
— Trate todos os membros com respeito. Ofensas, discriminações ou assédios de qualquer tipo não serão tolerados.

- ` + "`Regra 2`" + `
🗣️ ┆ **Linguagem Apropriada**:
— Evite palavrões e linguagem ofensiva. Mantenha as conversas em um nível apropriado para todos os membros.

- ` + "`Regra 3`" + `
🚫 ┆ **Conteúdo Adequado**:
— Não compartilhe conteúdo impróprio (pornografia, violência, etc.).

- ` + "`Regra 4`" + `
📨 ┆ **Spam e Flood**:
— Não envie mensagens repetitivas ou irrelevantes.

- ` + "`Regra 5`" + `
📋 ┆ **Canais e Tópicos**:
— Use os canais adequadamente conforme o tópico.

- ` + "`Regra 6`" + `
⚖️ ┆ **Proibições Legais**:
— Atividades ilegais são estritamente proibidas.

- ` + "`Regra 7`" + `
📢 ┆ **Divulgação e Publicidade**:
— Não divulgue links ou produtos sem permissão.

- ` + "`Regra 8`" + `
🔒 ┆ **Privacidade**:
— Não compartilhe informações pessoais de outros membros.

- ` + "`Regra 9`" + `
🎙️ ┆ **Comportamento em Chamadas**:
— Respeite os outros durante chamadas de voz.

- ` + "`Regra 10`" + `
👮 ┆ **Siga as Instruções da Equipe**:
— Os moderadores estão aqui para ajudar. Siga suas instruções.`

	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		// Se não estiver no cache, tenta buscar do Discord
		guild, err = s.Guild(m.GuildID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Não foi possível obter as informações do servidor.")
			return
		}
	}

	if guild.Icon == "" {
		s.ChannelMessageSend(m.ChannelID, "O servidor não tem um ícone definido.")
		return
	}

	iconURL := fmt.Sprintf("https://cdn.discordapp.com/icons/%s/%s.png", guild.ID, guild.Icon)
	thumbnailURL := "https://media.tenor.com/xzT-FtV3-MEAAAAM/xxxholic-yuuko.gif"

	embed := &discordgo.MessageEmbed{
		Title:       "📜 **REGRAS DO SERVIDOR**",
		Description: rules,
		Color:       0xff4600,
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: m.Author.AvatarURL(""),
			Name:    m.Author.Username,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: thumbnailURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    guild.Name,
			IconURL: iconURL,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err == nil {
		RulesMessageID = msg.ID
		s.MessageReactionAdd(m.ChannelID, msg.ID, "✅")
		s.ChannelMessageSend(m.ChannelID, "📢 **APÓS LER AS REGRAS, REAJA A ESTA MENSAGEM COM `✅` PARA LIBERAR OS OUTROS CANAIS DO SERVIDOR.**")
	}
}

package admin

import (
	"Turgho/Yuuko-BOT/internal/services/utils"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

// RulesMessageID armazena o ID da mensagem das regras (usado para reações, etc)
var RulesMessageID string

// RulesSlashCommand envia um embed com as regras do servidor
// e orienta os membros a reagirem para obter acesso total.
func RulesSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Determina o usuário que executou o comando
	var user *discordgo.User
	if i.Member != nil {
		user = i.Member.User
	} else {
		user = i.User
	}

	// Conteúdo das regras (utiliza crases para formato inline no Discord)
	rules := `
Bem-vindo ao nosso servidor! Para manter um ` + "`ambiente saudável`" + `, siga as regras abaixo:

- ` + "`Regra 1`" + ` 👐 **Respeito Mútuo**  
Trate todos com respeito. Ofensas, discriminações ou assédio não são tolerados.

- ` + "`Regra 2`" + ` 🗣️ **Linguagem Apropriada**  
Evite palavrões e linguagem ofensiva.

- ` + "`Regra 3`" + ` 🚫 **Conteúdo Adequado**  
Proibido conteúdo impróprio.

- ` + "`Regra 4`" + ` 📨 **Spam e Flood**  
Evite mensagens repetitivas ou irrelevantes.

- ` + "`Regra 5`" + ` 📋 **Canais e Tópicos**  
Use os canais corretamente.

- ` + "`Regra 6`" + ` ⚖️ **Proibições Legais**  
Atividades ilegais são proibidas.

- ` + "`Regra 7`" + ` 📢 **Publicidade**  
Não divulgue links ou produtos sem permissão.

- ` + "`Regra 8`" + ` 🔒 **Privacidade**  
Não compartilhe informações pessoais de terceiros.

- ` + "`Regra 9`" + ` 🎙️ **Chamadas de Voz**  
Mantenha o respeito nas chamadas.

- ` + "`Regra 10`" + ` 👮 **Equipe de Moderação**  
Siga as instruções dos moderadores.`

	// Obtém informações do servidor
	guild, iconURL, err := utils.GetGuildInfo(s, i.GuildID)
	if err != nil {
		utils.SendErrorResponse(s, i, "❌ Não foi possível obter as informações do servidor.")
		log.Println("Erro ao obter informações do servidor ", i.GuildID)
		return
	}

	// Cria o embed das regras
	embed := &discordgo.MessageEmbed{
		Title:       "📜 **REGRAS DO SERVIDOR**",
		Description: rules,
		Color:       0xff4600,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    user.Username,
			IconURL: user.AvatarURL(""),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/xzT-FtV3-MEAAAAM/xxxholic-yuuko.gif",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    guild.Name,
			IconURL: iconURL,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// Envia o embed como resposta à interação
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
	if err != nil {
		log.Println("Erro ao responder slash:", err)
		return
	}

	// Busca a última mensagem enviada do bot (embed das regras)
	// para adicionar a reação
	msgs, _ := s.ChannelMessages(i.ChannelID, 1, "", "", "")
	if len(msgs) > 0 {
		RulesMessageID = msgs[0].ID
		s.MessageReactionAdd(i.ChannelID, RulesMessageID, "✅")
		s.ChannelMessageSend(i.ChannelID, "📢 **APÓS LER AS REGRAS, REAJA COM `✅` PARA LIBERAR OS CANAIS.**")
	}
}

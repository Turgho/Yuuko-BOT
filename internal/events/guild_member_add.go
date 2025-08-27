package events

import (
	"Turgho/Yuuko-BOT/internal/services/utils"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func OnGuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	log.Println("➡️ Evento GuildMemberAdd acionado:", m.User.Username)
	if m.User == nil {
		log.Println("Evento GuildMemberAdd recebido, mas m.User é nil")
		return
	}

	channelID := os.Getenv("WELCOME_CHANNEL_ID")
	if channelID == "" {
		log.Println("WELCOME_CHANNEL_ID não encontrado no .env")
		return
	}

	welcomeMsg := "Bem-vindo ao servidor, <@" + m.User.ID + ">!\n\n" +
		"Leia o canal <#" + channelID + "> para mais informações."

	guild, iconURL, err := utils.GetGuildInfo(s, m.GuildID)
	if err != nil {
		log.Println("Erro ao obter informações do servidor", m.GuildID, err)
		return
	}

	avatarURL := m.User.AvatarURL("")

	embed := &discordgo.MessageEmbed{
		Title:       "🎉 UM NOVO MEMBRO CHEGOU!",
		Description: welcomeMsg,
		Color:       0xff4600,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: avatarURL,
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://giffiles.alphacoders.com/192/192655.gif",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    guild.Name,
			IconURL: iconURL,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err = s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		log.Println("Erro ao enviar mensagem de boas-vindas:", err)
	}
}

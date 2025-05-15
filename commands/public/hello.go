package public

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HelloCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := fmt.Sprintf("Olá, %s 👋", m.Author.Username)
	s.ChannelMessageSend(m.ChannelID, message)
}

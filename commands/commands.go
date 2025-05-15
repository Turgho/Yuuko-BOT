package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var prefix = "!"

func HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	content := m.Content
	if !strings.HasPrefix(content, prefix) {
		return
	}

	args := strings.Split(strings.TrimPrefix(content, prefix), " ")
	cmd := args[0]

	if handler, ok := CommandMap[cmd]; ok {
		handler(s, m)
	}
}

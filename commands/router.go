package commands

import (
	"github.com/bwmarrin/discordgo"
)

var CommandMap = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate){
	"ping":  PingCommand,
	"hello": HelloCommand,
	"rules": RulesCommand,
	"purge": PurgeAllCommand,
	// "outro": OutroCommand,
}

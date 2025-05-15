package router

import (
	"Turgho/Yuuko-BOT/commands/admin"
	"Turgho/Yuuko-BOT/commands/public"

	"github.com/bwmarrin/discordgo"
)

var PublicCommands = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate){
	"ping":  public.PingCommand,
	"hello": public.HelloCommand,
	// outros comandos públicos
}

var AdminCommands = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate){
	"purge":   admin.PurgeAllCommand,
	"rules":   admin.RulesCommand,
	"restart": admin.RestartCommand,
	// outros comandos restritos
}

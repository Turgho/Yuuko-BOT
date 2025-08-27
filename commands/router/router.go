package router

import (
	"Turgho/Yuuko-BOT/commands/admin"
	"Turgho/Yuuko-BOT/commands/public"

	"github.com/bwmarrin/discordgo"
)

var PublicCommands = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	// comandos public slash
	"ping":     public.PingCommand,
	"hello":    public.HelloSlashCommand,
	"coinflip": public.CoinflipSlashCommand,
}

var AdminCommands = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	// comandos admin slash
	"rules":    admin.RulesSlashCommand,
	"purge":    admin.PurgeSlashCommand,
	"restart":  admin.RestartSlashCommand,
	"shutdown": admin.ShutdownSlashCommand,
	"kick":     admin.KickUserSlashCommand,
}

var GamesCommands = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}

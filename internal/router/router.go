package router

import (
	"Turgho/Yuuko-BOT/internal/commands/admin"
	"Turgho/Yuuko-BOT/internal/commands/games"
	"Turgho/Yuuko-BOT/internal/commands/public"

	"github.com/bwmarrin/discordgo"
)

var PublicCommands = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	// comandos public slash
	"ping":    public.PingCommand,
	"hello":   public.HelloSlashCommand,
	"weather": public.WeatherSlashCommand,
}

var AdminCommands = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	// comandos admin slash
	"rules":      admin.RulesSlashCommand,
	"purge":      admin.PurgeSlashCommand,
	"restart":    admin.RestartSlashCommand,
	"shutdown":   admin.ShutdownSlashCommand,
	"kick":       admin.KickUserSlashCommand,
	"setmember":  admin.SetMemberRoleSlashCommand,
	"setwelcome": admin.SetWelcomeChannelSlashCommand,
}

var GamesCommands = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"coinflip": games.CoinflipSlashCommand,
}

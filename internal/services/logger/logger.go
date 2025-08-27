package logger

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

// LogCommand registra informações de execução de um comando
func LogCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.Member.User
	guildID := i.GuildID
	channelID := i.ChannelID
	commandName := i.ApplicationCommandData().Name
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	log.Printf("[COMMAND] %s | GuildID: %s | ChannelID: %s | User: %s#%s (%s) | Command: %s\n",
		timestamp,
		guildID,
		channelID,
		user.Username,
		user.Discriminator,
		user.ID,
		commandName,
	)
}

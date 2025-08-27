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

// LogReaction registra informações de uma reação adicionada
func LogReaction(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Pega informações do usuário
	user, err := s.User(r.UserID)
	if err != nil {
		log.Printf("[REACTION] %s | GuildID: %s | ChannelID: %s | UserID: %s | MsgID: %s | Emoji: %s (erro ao obter usuário)\n",
			timestamp, r.GuildID, r.ChannelID, r.UserID, r.MessageID, r.Emoji.Name)
		return
	}

	log.Printf("[REACTION] %s | GuildID: %s | ChannelID: %s | User: %s#%s (%s) | MsgID: %s | Emoji: %s\n",
		timestamp,
		r.GuildID,
		r.ChannelID,
		user.Username,
		user.Discriminator,
		user.ID,
		r.MessageID,
		r.Emoji.Name,
	)
}

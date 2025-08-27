package events

import "github.com/bwmarrin/discordgo"

func RegisterEventsHandler(s *discordgo.Session) {
	s.AddHandler(OnGuildMemberAdd)
}

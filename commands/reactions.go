package commands

import (
	"log"
	"os"

	"Turgho/Yuuko-BOT/commands/admin"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func HandleReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	// Pega token do discord
	roleID := os.Getenv("CARGO_MEMBRO")
	if roleID == "" {
		log.Fatal("CARGO_MEMBRO não encontrado no .env")
	}

	// Ignora reações do próprio bot
	if r.UserID == s.State.User.ID {
		return
	}

	// Verifica se a reação é no ID da mensagem de regras e se é o emoji de check
	if r.MessageID == admin.RulesMessageID && r.Emoji.Name == "✅" {
		err := s.GuildMemberRoleAdd(r.GuildID, r.UserID, roleID)
		if err != nil {
			log.Printf("Erro ao adicionar cargo ao usuário %s: %v", r.UserID, err)
			return
		}

		s.ChannelMessageSend(r.ChannelID, "<@"+r.UserID+"> aceitou as regras!")
	}
}

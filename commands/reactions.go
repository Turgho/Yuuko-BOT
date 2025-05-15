package commands

import (
	"log"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	loadEnvOnce    sync.Once
	roleID         string
	rulesMessageID string
)

// HandleReactionAdd é chamado quando um usuário adiciona uma reação em uma mensagem.
// Caso seja a mensagem de regras com o emoji ✅, o usuário recebe o cargo de membro.
func HandleReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// Carrega .env apenas uma vez
	loadEnvOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Erro ao carregar .env")
		}

		roleID = os.Getenv("CARGO_MEMBRO")
		if roleID == "" {
			log.Fatal("CARGO_MEMBRO não definido no .env")
		}

		rulesMessageID = os.Getenv("RULES_MESSAGE_ID")
		if roleID == "" {
			log.Fatal("RULES_MESSAGE_ID não definido no .env")
		}
	})

	// Ignora reações do próprio bot
	if r.UserID == s.State.User.ID {
		return
	}

	// Verifica se a reação é na mensagem correta com o emoji correto
	if r.MessageID == rulesMessageID && r.Emoji.Name == "✅" {
		err := s.GuildMemberRoleAdd(r.GuildID, r.UserID, roleID)
		if err != nil {
			log.Printf("❌ Erro ao adicionar cargo ao usuário %s: %v", r.UserID, err)
			return
		}

		// Envia confirmação pública no canal
		s.ChannelMessageSend(r.ChannelID, "<@"+r.UserID+"> leu e aceitou as regras! ✅")
	}
}

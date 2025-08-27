package games

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func CoinflipSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	coinflip := []string{"cara", "coroa"}

	message := fmt.Sprintf("A moeda caiu em %s 🪙", coinflip[rand.Intn(2)])

	// Envia a resposta da interação
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})

	if err != nil {
		log.Println("Erro ao enviar resposta do coinflip:", err)
		return
	}
}

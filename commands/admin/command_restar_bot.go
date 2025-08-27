package admin

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/bwmarrin/discordgo"
)

func RestartSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Envia resposta à interação (em vez de mensagem solta)
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "🔄 Reiniciando o bot...",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Println("Erro ao responder interação de reinício:", err)
	}

	// Executa o script de reinício
	cmd := exec.Command("bash", "./restart.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		log.Println("Erro ao iniciar novo processo:", err)
		return
	}

	log.Println("Bot reiniciado às ", time.Now().Format(time.RFC3339))

	// Fecha a sessão para desconectar do gateway
	s.Close()

	// Encerra o processo atual após iniciar o novo
	os.Exit(0)
}

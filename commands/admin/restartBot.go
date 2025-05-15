package admin

import (
	"log"
	"os"
	"os/exec"

	"github.com/bwmarrin/discordgo"
)

func RestartCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "🔄 Reiniciando o bot...")
	if err != nil {
		log.Println("Erro ao enviar mensagem de reinício:", err)
	}

	cmd := exec.Command("bash", "restart.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		log.Println("Erro ao iniciar novo processo:", err)
		s.ChannelMessageSend(m.ChannelID, "❌ Erro ao reiniciar.")
		return
	}

	os.Exit(0)
}

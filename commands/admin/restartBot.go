package admin

import (
	"log"
	"os"
	"os/exec"

	"github.com/bwmarrin/discordgo"
)

func RestartCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	execPath, err := os.Executable() // caminho do executável atual
	if err != nil {
		log.Println("Erro ao obter executável:", err)
		s.ChannelMessageSend(m.ChannelID, "❌ Erro ao reiniciar o bot.")
		return
	}

	cmd := exec.Command(execPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		log.Println("Erro ao iniciar novo processo:", err)
		s.ChannelMessageSend(m.ChannelID, "❌ Erro ao reiniciar o bot.")
		return
	}

	s.ChannelMessageSend(m.ChannelID, "🔄 Reiniciando o bot...")

	log.Println("Novo processo iniciado com PID:", cmd.Process.Pid)

	// Sai do processo atual para que o novo processo tome lugar
	os.Exit(0)
}

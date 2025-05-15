package main

import (
	"Turgho/Yuuko-BOT/commands"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func initLog() {
	// Criar pasta "logs" se não existir
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if err := os.Mkdir("logs", 0755); err != nil {
			log.Fatal("Erro ao criar pasta logs:", err)
		}
	}

	// Abrir arquivo para escrita e adicionar logs nele
	file, err := os.OpenFile("logs/yuuko_bot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Erro ao abrir arquivo de log:", err)
	}

	// Redireciona todos os logs para o arquivo
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Adiciona data, hora e linha
}

func main() {
	// Inicia o arquivo de logs
	initLog()

	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN não encontrado no .env")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Erro ao criar sessão:", err)
	}

	const prefix = "!"

	// Adiciona handler de mensagens
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot {
			return
		}

		content := strings.TrimSpace(m.Content)

		// Verifica se começa com o prefixo
		if !strings.HasPrefix(content, prefix) {
			return
		}

		// Remove o prefixo
		command := strings.TrimPrefix(content, prefix)

		// Pega só a primeira palavra (caso venha com argumentos depois)
		fields := strings.Fields(command)
		if len(fields) == 0 {
			return
		}
		cmdName := fields[0]

		// Busca e executa o comando
		if handler, ok := commands.CommandMap[cmdName]; ok {
			handler(s, m)
		}
	})

	// Adiciona handler de reações
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
		if r.UserID == s.State.User.ID {
			return // Ignora reações do próprio bot
		}

		// Verifica se é o emoji de check
		if r.MessageID == commands.RulesMessageID && r.Emoji.Name == "✅" {
			roleID := "1256066701548720209"

			// Reação canal de regras
			s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("<@%s> aceitou as regras!", r.UserID))
			err := s.GuildMemberRoleAdd(r.GuildID, r.UserID, roleID)
			if err != nil {
				log.Printf("Erro ao adicionar cargo ao usuário: %s\n", r.Emoji.User.Username)
			}
		}
	})

	err = dg.Open()
	if err != nil {
		log.Fatal("Erro ao conectar-se ao Discord:", err)
	}
	defer dg.Close()

	log.Printf("Logado como: %s#%s!\n", dg.State.User.Username, dg.State.User.Discriminator)

	// Esperar até CTRL+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	log.Println("Desligando bot.")
}

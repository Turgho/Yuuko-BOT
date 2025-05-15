package main

import (
	"Turgho/Yuuko-BOT/commands"
	"log"
	"os"
	"os/signal"
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

	// Pega token do discord
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN não encontrado no .env")
	}

	// Inicia uma nova sessão
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Erro ao criar sessão:", err)
	}

	// Adiciona handler de mensagens
	dg.AddHandler(commands.HandleCommand)

	// Adiciona handler de reações
	dg.AddHandler(commands.HandleReactionAdd)

	// Inicia a conexão com api do discord
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

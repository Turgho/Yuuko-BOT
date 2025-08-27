package main

import (
	"Turgho/Yuuko-BOT/commands"
	"Turgho/Yuuko-BOT/config"
	"Turgho/Yuuko-BOT/events"
	"Turgho/Yuuko-BOT/register"
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

	// Carregar config.json
	if err := config.LoadConfig("config/config.json"); err != nil {
		log.Fatal("Erro ao carregar config.json: ", err)
	}

	// Inicia uma nova sessão
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Erro ao criar sessão:", err)
	}

	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers | discordgo.IntentGuildMessageReactions

	// Adiciona handler de mensagens
	dg.AddHandler(commands.HandleSlashCommand)

	// Adiciona handler de reações
	dg.AddHandler(commands.HandleReactionAdd)

	dg.AddHandler(commands.HandleButtonInteraction)

	// Registra os eventos
	events.RegisterEventsHandler(dg)

	// Inicia a conexão com api do discord
	err = dg.Open()
	if err != nil {
		log.Fatal("Erro ao conectar-se ao Discord:", err)
	}
	defer dg.Close()

	// Registrar os comandos
	appID := dg.State.User.ID
	register.RegisterAllCommands(dg, appID, config.Cfg.Guilds)

	log.Printf("Logado como: %s#%s!\n", dg.State.User.Username, dg.State.User.Discriminator)

	// Esperar até CTRL+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	log.Println("Desligando bot.")
}

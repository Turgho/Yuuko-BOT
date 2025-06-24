package main

import (
	"Turgho/Yuuko-BOT/commands"
	"Turgho/Yuuko-BOT/db"
	"Turgho/Yuuko-BOT/events"
	"Turgho/Yuuko-BOT/models/monsterhunter/armor"
	"Turgho/Yuuko-BOT/models/monsterhunter/weapon"
	"Turgho/Yuuko-BOT/register"
	"Turgho/Yuuko-BOT/server"
	"fmt"
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
	server.ServeStaticIcons()

	fmt.Println("Servidor de ícones rodando em http://localhost:8000/icons/")

	// Inicia o arquivo de logs
	initLog()

	// Inicia a Database
	db.InitDB()

	if err := weapon.LoadWeapons(); err != nil {
		log.Fatalf("Falha crítica ao carregar armas: %v", err)
	}

	if err := armor.LoadArmors(); err != nil {
		log.Fatalf("Falha crítica ao carregar armaduras: %v", err)
	}

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

	guildID := os.Getenv("GUILD_ID")
	if guildID == "" {
		log.Fatal("GUILD_ID não encontrado no .env")
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
	register.RegisterAllCommands(dg, appID, guildID)

	log.Printf("Logado como: %s#%s!\n", dg.State.User.Username, dg.State.User.Discriminator)

	// Esperar até CTRL+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	log.Println("Desligando bot.")
}

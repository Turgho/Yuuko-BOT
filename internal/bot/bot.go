package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"Turgho/Yuuko-BOT/config"
	"Turgho/Yuuko-BOT/internal/commands"
	"Turgho/Yuuko-BOT/internal/events"
	"Turgho/Yuuko-BOT/register"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// StartBot inicializa e conecta o bot ao Discord
func StartBot() {
	// Inicializa logging
	initLog()

	// Carrega variáveis de ambiente do .env
	loadEnv()

	// Carrega arquivo de configuração config.json
	loadConfig()

	// Cria sessão do Discord
	dg := createDiscordSession()

	// Registra handlers de comandos e eventos
	registerHandlers(dg)

	// Conecta ao Discord
	openSession(dg)
	defer dg.Close()

	// Registra os comandos nas guilds
	registerCommands(dg)

	log.Printf("🤖 Bot logado como: %s#%s\n", dg.State.User.Username, dg.State.User.Discriminator)

	// Mantém o bot ativo até receber sinal de interrupção
	waitForShutdown()
}

// ---------------------------- Funções auxiliares ----------------------------

// initLog cria a pasta logs (se necessário) e redireciona logs para arquivo
func initLog() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if err := os.Mkdir("logs", 0755); err != nil {
			log.Fatal("Erro ao criar pasta logs:", err)
		}
	}

	file, err := os.OpenFile("logs/yuuko_bot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Erro ao abrir arquivo de log:", err)
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// loadEnv carrega as variáveis do arquivo .env
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	if os.Getenv("DISCORD_TOKEN") == "" {
		log.Fatal("DISCORD_TOKEN não encontrado no .env")
	}
}

// loadConfig carrega o arquivo config.json
func loadConfig() {
	if err := config.LoadConfig("config/config.json"); err != nil {
		log.Fatal("Erro ao carregar config.json:", err)
	}
}

// createDiscordSession cria a sessão Discord com intents necessárias
func createDiscordSession() *discordgo.Session {
	token := os.Getenv("DISCORD_TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Erro ao criar sessão Discord:", err)
	}

	dg.Identify.Intents = discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMembers |
		discordgo.IntentsGuildMessageReactions

	return dg
}

// registerHandlers registra todos os handlers de comandos e eventos
func registerHandlers(dg *discordgo.Session) {
	// Comandos do tipo slash
	dg.AddHandler(commands.HandleInteraction)

	// Reações adicionadas
	dg.AddHandler(commands.HandleReactionAdd)

	// Outros eventos
	events.RegisterEventsHandler(dg)
}

// openSession abre a conexão com a API do Discord
func openSession(dg *discordgo.Session) {
	if err := dg.Open(); err != nil {
		log.Fatal("Erro ao conectar-se ao Discord:", err)
	}
}

// registerCommands registra todos os comandos nas guilds configuradas
func registerCommands(dg *discordgo.Session) {
	appID := dg.State.User.ID
	register.RegisterAllCommands(dg, appID, config.Cfg.Guilds)
}

// waitForShutdown mantém o bot ativo até receber CTRL+C ou sinal de desligamento
func waitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
	log.Println("⚡ Desligando bot.")
}

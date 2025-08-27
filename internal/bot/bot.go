package bot

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

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

const logDir = "logs"
const maxDays = 7

// initLog cria a pasta logs (se necessário) e redireciona logs para arquivo
func initLog() {
	// Cria pasta logs se não existir
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			log.Fatal("Erro ao criar pasta logs:", err)
		}
	}

	// Nome do arquivo baseado na data
	today := time.Now().Format("2006-01-02")
	logFile := filepath.Join(logDir, today+".log")

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Erro ao abrir arquivo de log:", err)
	}

	// Redireciona logs
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Remove logs antigos
	removeOldLogs()
}

// removeOldLogs deleta arquivos de log com mais de maxDays
func removeOldLogs() {
	files, err := os.ReadDir(logDir)
	if err != nil {
		log.Printf("Erro ao ler pasta de logs: %v", err)
		return
	}

	cutoff := time.Now().AddDate(0, 0, -maxDays)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		filePath := filepath.Join(logDir, f.Name())
		info, err := f.Info()
		if err != nil {
			continue
		}

		if info.ModTime().Before(cutoff) {
			os.Remove(filePath)
		}
	}
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
	// Carrega config.json e popula o map global CfgMap
	configMap, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatal("Erro ao carregar config.json:", err)
	}

	// Salva no global
	config.CfgMap = configMap

	// Se quiser, pode logar as guilds carregadas
	for id := range configMap {
		log.Printf("✅ Guild carregada: %s", id)
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
	register.RemoveObsoleteCommandsAllGuilds(dg)
	register.RegisterAllCommands(dg, appID, config.CfgMap)
}

// waitForShutdown mantém o bot ativo até receber CTRL+C ou sinal de desligamento
func waitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
	log.Println("⚡ Desligando bot.")
}

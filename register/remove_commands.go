package register

import (
	"Turgho/Yuuko-BOT/config"
	"Turgho/Yuuko-BOT/internal/router"
	"log"

	"github.com/bwmarrin/discordgo"
)

// RemoveCommandByName deleta um comando específico
func RemoveCommandByName(dg *discordgo.Session, appID, guildID, cmdName string) error {
	commands, err := dg.ApplicationCommands(appID, guildID)
	if err != nil {
		return err
	}

	for _, cmd := range commands {
		if cmd.Name == cmdName {
			return dg.ApplicationCommandDelete(appID, guildID, cmd.ID)
		}
	}

	return nil // comando não encontrado
}

// RemoveObsoleteCommands deleta todos os comandos que não estão no router
func RemoveObsoleteCommandsAllGuilds(dg *discordgo.Session) {
	validCommands := []string{}
	for name := range router.AdminCommands {
		validCommands = append(validCommands, name)
	}
	for name := range router.PublicCommands {
		validCommands = append(validCommands, name)
	}
	for name := range router.GamesCommands {
		validCommands = append(validCommands, name)
	}

	for guildID := range config.CfgMap {
		commands, err := dg.ApplicationCommands(dg.State.User.ID, guildID)
		if err != nil {
			log.Println("Erro ao listar comandos da guild:", err)
			continue
		}

		validMap := make(map[string]bool)
		for _, name := range validCommands {
			validMap[name] = true
		}

		for _, cmd := range commands {
			if !validMap[cmd.Name] {
				err := dg.ApplicationCommandDelete(dg.State.User.ID, guildID, cmd.ID)
				if err != nil {
					log.Printf("Erro ao remover comando %s na guild %s: %v", cmd.Name, guildID, err)
				} else {
					log.Printf("✅ Comando removido: %s na guild %s", cmd.Name, guildID)
				}
			}
		}
	}
}

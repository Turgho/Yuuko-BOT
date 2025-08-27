package public

import (
	"Turgho/Yuuko-BOT/internal/services/utils"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// HandleWeatherCommand é o handler do slash command /weather
func WeatherSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		respondEphemeral(s, i, "❌ Você precisa informar uma cidade.")
		return
	}

	city := options[0].StringValue()
	weather, err := utils.GetWeather(city)
	if err != nil {
		log.Println("Erro ao buscar clima:", err)
		respondEphemeral(s, i, "❌ Não foi possível obter a previsão.")
		return
	}

	weatherType := utils.WeatherCodeToText(weather.Current.WeatherCode)

	msg := fmt.Sprintf("🌤 **Previsão para %s**:\n- Temperatura: %.1f°C\n- Velocidade do vento: %.1f km/h\n- Clima: %s",
		city, weather.Current.Temperature, weather.Current.Windspeed, weatherType)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
}

// respondEphemeral envia mensagem visível só para quem chamou o comando
func respondEphemeral(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

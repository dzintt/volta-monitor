package discord

import (
	"strconv"
	"strings"
	"time"
	"volta-monitor/models"

	"github.com/bwmarrin/discordgo"
)

func SendToDiscord(webhook string, station models.ChargerStation) error {
	discordID := strings.Split(webhook, "/")[5]
	discordToken := strings.Split(webhook, "/")[6]
	discordSession, err := discordgo.New("")
	if err != nil {
		return err
	}

	var color int
	var title string
	if station.Availability.State == "IDLE" || station.Availability.State == "PLUGGED_OUT" {
		color = 0x5af269
		title = "🟢 Volta Charging Station Available 🟢"
	} else {
		color = 0xfa556b
		title = "🔴 Volta Charging Station In Use 🔴"
	}

	discordSession.WebhookExecute(discordID, discordToken, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       title,
				Description: station.Description,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Station Name",
						Value:  station.Name,
						Inline: true,
					},
					{
						Name:   "Time Limit",
						Value:  strconv.Itoa(int(station.TimeLimit)/60) + " minutes",
						Inline: true,
					},
					{
						Name:   "State",
						Value:  station.Availability.State,
						Inline: true,
					},
					{
						Name:   "Connectors",
						Value:  strings.Join(station.Availability.Connectors, "\n"),
						Inline: true,
					},
				},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: station.Image,
				},
				Color: color,
				Footer: &discordgo.MessageEmbedFooter{
					Text: "Volta Monitor | Station: " + station.ID,
				},
				Timestamp: time.Now().Format(time.RFC3339),
			},
		},
		Username: "Volta Monitor",
	})

	return nil
}

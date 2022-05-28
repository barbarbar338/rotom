package commands

import (
	"fmt"
	"rotom/db"
	"rotom/logger"

	"github.com/bwmarrin/discordgo"
)

var SeviyeData = discordgo.ApplicationCommand {
	Name: "seviye",
	Description: "Pokémon Unite/GO oyun içi seviyenizi ayarlar",
	Options: []*discordgo.ApplicationCommandOption {
		{
			Name: "go",
			Description: "Pokémon GO oyun içi seviyenizi ayarlar",
			Type: discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption {
				{
					Name: "seviye",
					Description: "Pokémon GO oyun içi seviyeniz",
					Required: true,
					Type: discordgo.ApplicationCommandOptionInteger,
				},
			},
		},
		{
			Name: "unite",
			Description: "Pokémon Unite oyun içi seviyenizi ayarlar",
			Type: discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption {
				{
					Name: "seviye",
					Description: "Pokémon Unite oyun içi seviyeniz",
					Required: true,
					Type: discordgo.ApplicationCommandOptionInteger,
				},
			},
		},
	},
}

func SeviyeHandler(s *discordgo.Session, i *discordgo.InteractionCreate, u *db.UserModel) {
	options := i.ApplicationCommandData().Options
	option := options[0]
	level := uint64(option.Options[0].Value.(float64))

	switch option.Name {
		case "go":
			u.PokemonGOLevel = level
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData {
					Content: fmt.Sprintf("Pokémon GO seviyeniz `%v` olarak ayarlandı! `/profil` komutu ile profilinizi görün.", level),
				},
			})
		case "unite":
			u.PokemonUniteLevel = level
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData {
					Content: fmt.Sprintf("Pokémon Unite seviyeniz `%v` olarak ayarlandı! `/profil` komutu ile profilinizi görün.", level),
				},
			})
	}

	_, err := db.Database.Model(u).WherePK().Update()
	if err != nil {
		logger.Logger.ErrorF("Error updating user: %v", err)
	}
}

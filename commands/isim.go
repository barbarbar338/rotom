package commands

import (
	"fmt"
	"rotom/db"
	"rotom/logger"

	"github.com/bwmarrin/discordgo"
)

var IsimData = discordgo.ApplicationCommand {
	Name: "isim",
	Description: "Pokémon Unite/GO oyun içi isminizi ayarlar",
	Options: []*discordgo.ApplicationCommandOption {
		{
			Name: "go",
			Description: "Pokémon GO oyun içi isminizi ayarlar",
			Type: discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption {
				{
					Name: "isim",
					Description: "Pokémon GO oyun içi isminiz",
					Required: true,
					Type: discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{
			Name: "unite",
			Description: "Pokémon Unite oyun içi isminizi ayarlar",
			Type: discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption {
				{
					Name: "isim",
					Description: "Pokémon Unite oyun içi isminiz",
					Required: true,
					Type: discordgo.ApplicationCommandOptionString,
				},
			},
		},
	},
}

func IsimHandler(s *discordgo.Session, i *discordgo.InteractionCreate, u *db.UserModel) {
	options := i.ApplicationCommandData().Options
	option := options[0]
	name := option.Options[0].Value.(string)

	switch option.Name {
		case "go":
			u.PokemonGOUsername = name
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData {
					Content: fmt.Sprintf("Pokémon GO oyun içi isminiz `%v` olarak ayarlandı! `/profil` komutu ile profilinizi görün.", name),
				},
			})
		case "unite":
			u.PokemonUniteUsername = name
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData {
					Content: fmt.Sprintf("Pokémon Unite oyun içi isminiz `%v` olarak ayarlandı! `/profil` komutu ile profilinizi görün.", name),
				},
			})
	}

	_, err := db.Database.Model(u).WherePK().Update()
	if err != nil {
		logger.Logger.ErrorF("Error updating user: %v", err)
	}
}

package commands

import (
	"fmt"
	"rotom/db"
	"rotom/logger"

	"github.com/bwmarrin/discordgo"
)

var KodData = discordgo.ApplicationCommand {
	Name: "kod",
	Description: "Pokémon Unite/GO arkadaşlık kodunuzu ayarlar",
	Options: []*discordgo.ApplicationCommandOption {
		{
			Name: "go",
			Description: "Pokémon GO arkadaşlık kodunuzu ayarlar",
			Type: discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption {
				{
					Name: "kod",
					Description: "Pokémon GO arkadaşlık kodunuz",
					Required: true,
					Type: discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{
			Name: "unite",
			Description: "Pokémon Unite arkadaşlık kodunuzu ayarlar",
			Type: discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption {
				{
					Name: "kod",
					Description: "Pokémon Unite arkadaşlık kodunuz",
					Required: true,
					Type: discordgo.ApplicationCommandOptionString,
				},
			},
		},
	},
}

func KodHandler(s *discordgo.Session, i *discordgo.InteractionCreate, u *db.UserModel) {
	options := i.ApplicationCommandData().Options
	option := options[0]
	code := option.Options[0].Value.(string)

	switch option.Name {
		case "go":
			u.PokemonGOFriendCode = code
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData {
					Content: fmt.Sprintf("Pokémon GO arkadaş kodunuz `%v` olarak ayarlandı! `/profil` komutu ile profilinizi görün.", code),
				},
			})
		case "unite":
			u.PokemonUniteFriendCode = code
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData {
					Content: fmt.Sprintf("Pokémon Unite arkadaş kodunuz `%v` olarak ayarlandı! `/profil` komutu ile profilinizi görün.", code),
				},
			})
	}

	_, err := db.Database.Model(u).WherePK().Update()
	if err != nil {
		logger.Logger.ErrorF("Error updating user: %v", err)
	}
}

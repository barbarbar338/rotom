package commands

import (
	"rotom/db"

	"github.com/bwmarrin/discordgo"
)

var PingData = discordgo.ApplicationCommand {
	Name: "ping",
	Description: "Replies with pong!",
}

func PingHandler(s *discordgo.Session, i *discordgo.InteractionCreate, u *db.UserModel) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData {
			Content: ":ping_pong: Pong!",
		},
	})
}

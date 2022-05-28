package events

import (
	"rotom/commands"
	"rotom/db"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if cmd, ok := commands.Handlers[i.ApplicationCommandData().Name]; ok {	
   		u := db.GetUserModel(i.Member.User.ID)
		cmd(s, i, u)
	}
}

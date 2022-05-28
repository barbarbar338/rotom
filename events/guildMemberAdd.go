package events

import (
	"rotom/config"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberAdd(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	if event.GuildID != config.GuildID {
		return
	}

	if event.User.Bot {
		return
	}

	member, err := session.GuildMember(config.GuildID, event.User.ID)
	if err != nil {
		return
	}

	err = session.GuildMemberRoleAdd(config.GuildID, member.User.ID, config.AutoRole)
	if err != nil {
		return
	}
}

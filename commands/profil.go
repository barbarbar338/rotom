package commands

import (
	"fmt"
	"rotom/db"
	"rotom/logger"
	"rotom/utils"

	"github.com/bwmarrin/discordgo"
)

var ProfilData = discordgo.ApplicationCommand {
	Name: "profil",
	Description: "Sunucu profilinizi görmenizi sağlar",
	Options: []*discordgo.ApplicationCommandOption {
		{
			Name: "kullanıcı",
			Description: "Profiline bakılacak kullanıcı",
			Required: false,
			Type: discordgo.ApplicationCommandOptionUser,
		},
	},
}

func ProfilHandler(s *discordgo.Session, i *discordgo.InteractionCreate, u *db.UserModel) {
	member := i.Member
	options := i.ApplicationCommandData().Options
	if len(options) > 0 {
		user := options[0].UserValue(s)
		guild, err := s.Guild(i.GuildID)
		if err != nil {
			logger.Logger.ErrorF("Could not get guild: %s", err)
		} else {
			m, err := s.GuildMember(guild.ID, user.ID)
			if err != nil {
				logger.Logger.ErrorF("Could not get member: %s", err)
			} else {
				member = m
				u = db.GetUserModel(m.User.ID)
			}
		}

	}

	instinct := utils.ContainsS(member.Roles, "978310967668654120")
	valor := utils.ContainsS(member.Roles, "978310838328901703")
	mystic := utils.ContainsS(member.Roles, "978310925872414791")

	var icon string
	var color int
	if (instinct) {
		icon = "https://static.pokenav.app/images/team-logos/png/128/instinct.png"
		color = 0xFFE212 // yellow
	} else if (valor) {
		icon = "https://static.pokenav.app/images/team-logos/png/128/valor.png"
		color = 0xFF0000 // red
	} else if (mystic) {
		icon = "https://static.pokenav.app/images/team-logos/png/128/mystic.png"
		color = 0x005DFF // blue
	} else {
		icon = "https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Pok%C3%A9_Ball_icon.svg/2052px-Pok%C3%A9_Ball_icon.png"
		color = 0x01DC47 // green
	}

	fields := []*discordgo.MessageEmbedField {}
	
	if u.PokemonGOUsername != "" {
		fields = append(fields, &discordgo.MessageEmbedField {
			Name: "Pokémon GO Oyuncu Adı",
			Value: fmt.Sprintf("%v", u.PokemonGOUsername),
			Inline: true,
		})
	}
	if u.PokemonGOFriendCode != "" {
		fields = append(fields, &discordgo.MessageEmbedField {
			Name: "Pokémon GO Arkadaş Kodu",
			Value: fmt.Sprintf("%v", u.PokemonGOFriendCode),
			Inline: true,
		})
	}
	if u.PokemonGOLevel != 0 {
		fields = append(fields, &discordgo.MessageEmbedField {
			Name: "Pokémon GO Seviyesi",
			Value: fmt.Sprintf("%v", u.PokemonGOLevel),
			Inline: true,
		})
	}

	if u.PokemonUniteUsername != "" {
		fields = append(fields, &discordgo.MessageEmbedField {
			Name: "Pokémon Unite Oyuncu Adı",
			Value: fmt.Sprintf("%v", u.PokemonUniteUsername),
			Inline: true,
		})
	}
	if u.PokemonUniteFriendCode != "" {
		fields = append(fields, &discordgo.MessageEmbedField {
			Name: "Pokémon Unite Arkadaş Kodu",
			Value: fmt.Sprintf("%v", u.PokemonUniteFriendCode),
			Inline: true,
		})
	}
	if u.PokemonUniteLevel != 0 {
		fields = append(fields, &discordgo.MessageEmbedField {
			Name: "Pokémon Unite Seviyesi",
			Value: fmt.Sprintf("%v", u.PokemonUniteLevel),
			Inline: true,
		})
	}

	description := "%v kullanıcısının Pokémon GO ve Pokémon Unite profili hakkında aşağıdan bilgi alabilirsiniz."
	if len(fields) < 1 {
		description = "%v kullanıcısının Pokémon GO ve Pokémon Unite profili hakkında hiçbir bilgi bulunmamaktadır. Bu kullanıcıya profilini düzenlemesini söyleyebilirsiniz."
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData {
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Content: fmt.Sprintf(":pencil: **%v** kullanıcısının profili", member.User.Username),
			Embeds: []*discordgo.MessageEmbed {
				{
					Title: fmt.Sprintf(":pencil: %v kullanıcısının profili", member.User.Username),
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: icon,
					},
					Description: fmt.Sprintf(description, member.User.Username),
					Fields: fields,
					Color: color,
				},
			},
		},
	})
}

package commands

import (
	"rotom/db"
	"rotom/logger"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		&PingData,
		&IsimData,
		&KodData,
		&SeviyeData,
		&YardimData,
		&ProfilData,
	}
	
	Handlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate, *db.UserModel) {
		"ping": PingHandler,
		"isim": IsimHandler,
		"kod": KodHandler,
		"seviye": SeviyeHandler,
		"yardim": YardimHandler,
		"yardım": YardimHandler,
		"profil": ProfilHandler,
	}
)

func Migrate(s *discordgo.Session) {
	logger.Logger.Info("Migrating commands")
	for _, command := range commands {
		logger.Logger.InfoF("Posting command: %v", command.Name)
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
		if err != nil {
			logger.Logger.WarningF("Failed to create command: %s", err.Error())
		}
	}
}

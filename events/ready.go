package events

import (
	"time"

	"rotom/logger"
	"rotom/utils"

	"github.com/bwmarrin/discordgo"
)

const (
	statusIntervalPeriot	= 30 * time.Second
)

var (
	statusTexts 			= []string {
		"Yarım için /yardım komutunu kullanın!",
		"Pokémon GO seviyenizi /seviye go <seviye> komutu ile ayarlayın",
		"Pokémon Unite seviyenizi /seviye unite <seviye> komutu ile ayarlayın",
		"Pokémon GO isminizi /isim go <isim> komutu ile ayarlayın",
		"Pokémon Unite isminizi /isim unite <isim> komutu ile ayarlayın",
		"Pokémon GO arkadaşlık kodunuzu /kod go <kod> komutu ile ayarlayın",
		"Pokémon Unite arkadaşlık kodunuzu /kod unite <kod> komutu ile ayarlayın",
		"Profilinizi görmek için /profil komutunu kullanın",
	}
	StatusInterval			chan bool
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {	
	err := session.UpdateGameStatus(0, utils.RandomString(statusTexts))
	if err != nil {
		logger.Logger.WarningF("Failed to update game status: %s", err.Error())
	}

	StatusInterval = utils.SetInterval(func() {
		err := session.UpdateGameStatus(0, utils.RandomString(statusTexts))
		if err != nil {
			logger.Logger.WarningF("Failed to update game status: %s", err.Error())
		}
	}, statusIntervalPeriot)

	logger.Logger.InfoF("[%s:%s] Ready to serve!", session.State.User.Username, session.State.User.ID)
}

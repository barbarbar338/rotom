package main

import (
	"flag"
	"fmt"
	"os"

	"rotom/commands"
	"rotom/config"
	"rotom/db"
	"rotom/events"
	"rotom/logger"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func main() {
	var flagMigrateDB, flagMigrateCommands, flagMigrateAll bool

	flag.BoolVar(&flagMigrateDB, "db", false, "Migrate database")
	flag.BoolVar(&flagMigrateCommands, "commands", false, "Migrate commands")
	flag.BoolVar(&flagMigrateAll, "all", false, "Migrate commands and database")

	flag.Parse()

	s, err := discordgo.New(fmt.Sprintf("Bot %v", config.Token))
	if err != nil {
		logger.Logger.FatalF("Failed to create Discord session: %v", err.Error())
	}

	s.AddHandler(events.Ready)
	s.AddHandler(events.InteractionCreate)
	s.AddHandler(events.GuildMemberAdd)

	err = s.Open()
	if err != nil {
		logger.Logger.FatalF("Failed to connect to Discord API: %v", err.Error())
	}

	if flagMigrateAll {
		db.Migrate()
		commands.Migrate(s)
		os.Exit(0)
		return
	}

	if flagMigrateDB {
		db.Migrate()
		os.Exit(0)
		return
	}

	if flagMigrateCommands {
		commands.Migrate(s)
		os.Exit(0)
		return
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%v", config.PORT))
}

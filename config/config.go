package config

import (
	"os"

	"rotom/logger"

	"github.com/joho/godotenv"
)

var (
	Token			string
	RedisHost 		string
	RedisPort 		string
	RedisPassword	string
	PGHost 			string
	PGPort 			string
	PGUser 			string
	PGPassword 		string
	PGDatabase 		string
	PORT 			string
)

func init() {
	production := os.Getenv("PRODUCTION")

	if production != "true" {
		err := godotenv.Load()
		if err != nil {
			logger.Logger.Fatal("Error loading .env file: " + err.Error())
		}
	}

	Token = os.Getenv("BOT_TOKEN")
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	PGHost = os.Getenv("POSTGRE_HOST")
	PGPort = os.Getenv("POSTGRE_PORT")
	PGUser = os.Getenv("POSTGRE_USER")
	PGPassword = os.Getenv("POSTGRE_PASSWORD")
	PGDatabase = os.Getenv("POSTGRE_DATABASE")
	PORT = os.Getenv("PORT")
}

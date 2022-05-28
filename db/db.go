package db

import (
	"fmt"

	"rotom/config"
	"rotom/logger"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var (
	Database *pg.DB
)

func init() {
    d := pg.Connect(&pg.Options {
        Addr: fmt.Sprintf("%s:%s", config.PGHost, config.PGPort),
		User: config.PGUser,
		Password: config.PGPassword,
		Database: config.PGDatabase,
    })
	Database = d
}

func Migrate() {
    logger.Logger.Info("Migrating database")

    models := []interface{} {
        (*UserModel)(nil),
    }

    for _, model := range models {
        err := Database.Model(model).CreateTable(&orm.CreateTableOptions{
            Temp: false,
        })
    
		if err != nil {
            logger.Logger.FatalF("Failed to create schema: %s", err)
        }
    }
}

func GetUserModel(id string) (*UserModel) {
    	u := &UserModel {
			ID: id,
		}
		err := Database.Model(u).WherePK().Select()
		if err != nil {
			u = &UserModel {
				ID: id,
				PokemonGOLevel: 0,
				PokemonGOUsername: "",
				PokemonGOFriendCode: "",
				PokemonUniteLevel: 0,
				PokemonUniteUsername: "",
				PokemonUniteFriendCode: "",
			}
			Database.Model(u).Insert()
		}
        return u
}

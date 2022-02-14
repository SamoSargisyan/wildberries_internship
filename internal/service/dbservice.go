package service

import (
	"github.com/jmoiron/sqlx"
	"l0/config"
	"l0/pkg/db/postgres"
	"log"
	"path/filepath"
)

func getDb() *sqlx.DB {
	mainConfigFile, _ := filepath.Abs("./config/services.yml")

	cfg, err := config.GetDatabaseConfigurations(mainConfigFile)
	if err != nil {
		log.Fatalln(err)
	}

	db, err := postgres.NewConnection(cfg.Database.Combine)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

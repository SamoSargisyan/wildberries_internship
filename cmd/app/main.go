package main

import (
	"l0/config"
	"l0/internal/repository/postgres"
	"log"
)

func main() {
	// не нравится хардкодить строку. Требует исправления
	cfg, err := config.ReadConfigYML("/Users/samvelsargisan/golang/wildberries/l0/config/main.yml")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = postgres.NewConnection(cfg.Database.Combine)
	if err != nil {
		log.Fatalln(err)
	}
}

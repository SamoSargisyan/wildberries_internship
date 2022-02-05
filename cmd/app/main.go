package main

import (
	"l0/config"
	"l0/internal/service"
	"log"
	"path/filepath"
)

func main() {
	mainConfigFile, _ := filepath.Abs("./config/main.yml")

	cfg, err := config.GetDatabaseConfigurations(mainConfigFile)
	if err != nil {
		log.Fatalln(err)
	}

	//log.Fatal(cfg.Database.Combine)
	_, err = service.Bootstrap(cfg, "1")
	if err != nil {
		log.Fatalln(err)
	}
}

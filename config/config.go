package config

import (
	"fmt"
	_ "github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `env:"DB_PASSWORD"`
	Db       string `yaml:"dbname"`
	SslMode  string `yaml:"sslmode"`
	Combine  string
}

type Config struct {
	Database Database `yaml:"database"`
}

func closer(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("close error: %v", err)
	}
}

func GetDatabaseConfigurations(filePath string) (cfg *Config, err error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return cfg, err
	}

	defer closer(file)

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)

	if err != nil {
		return nil, err
	}
	if err != nil {
		return cfg, err
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// password removed
	cfg.Database.Combine = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		os.Getenv("DB_PASSWORD"),
		cfg.Database.Db,
		cfg.Database.SslMode,
	)

	return cfg, nil
}

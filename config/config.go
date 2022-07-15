package config

import (
	"log"

	environment "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

var Conf Service

type Service struct {
	Server   Server
	Database Database
}

type Server struct {
	Port string `env:"PORT"`
}

type Database struct {
	DatabaseFileName string `env:"DATABASE_FILE_NAME"`
}

var isLoaded = false

func Load() (Service, error) {
	if isLoaded {
		return Conf, nil
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading config file: ", err)
	}

	_, err = environment.UnmarshalFromEnviron(&Conf)
	if err != nil {
		log.Fatal("Error loading config file: ", err)
	}

	isLoaded = true
	return Conf, err
}

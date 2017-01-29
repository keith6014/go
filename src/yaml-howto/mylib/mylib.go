package mylib

import (
	"log"
)

func AAA() bool {
	return true
}

type Config struct {
	Info string
}

func ValidateConfig(myconfig *Config) bool {

	if len(myconfig.Info) == 0 {
		log.Print("info is not set")
		return false
	}

	return true
}

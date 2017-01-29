package mylib

import (
	"log"
)

func A() bool {
	return true
}

type Config struct {
	Info string
}

func ValidateConfig(myconfig *Config) bool {

	if len(myconfig.Info) == 0 {
		log.Fatal("info is not set")
	}

	return true
}

package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"yaml-howto/mylib"
)

type Config struct {
	Info string
}

func validateConfig(myconfig *Config) bool {

	if len(myconfig.Info) == 0 {
		log.Fatal("info is not set")
	}

	return true
}

func main() {
	reader, _ := os.Open("data/data.yaml")
	buf, _ := ioutil.ReadAll(reader)

	log.Printf("%v\n", mylib.A())
	var conf Config
	yaml.Unmarshal(buf, &conf)

	log.Printf("%v\n", validateConfig(&conf))

	log.Printf("%v\n", conf)

}

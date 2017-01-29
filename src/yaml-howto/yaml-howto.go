package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"yaml-howto/mylib"
)

func main() {
	reader, _ := os.Open("data/data.yaml")
	buf, _ := ioutil.ReadAll(reader)

	log.Printf("%v\n", mylib.A())
	var conf mylib.Config
	yaml.Unmarshal(buf, &conf)

	log.Printf("%v\n", mylib.ValidateConfig(&conf))

	log.Printf("%v\n", conf)

}

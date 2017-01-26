package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func run(cmd string) []byte {
	fmt.Println(strings.Fields(cmd))
	out, err := exec.Command(strings.Fields(cmd)).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func main() {
	fmt.Println("ok")
	run("/bin/ps")
}

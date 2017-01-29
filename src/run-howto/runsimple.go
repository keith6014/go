package main

import (
	"log"
	"os/exec"
	"strings"
)

func run(cmd string, verbose bool) []byte {
	if verbose {
		log.Printf(cmd + "\n")
	}
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func main() {
	log.Printf("ok\n")
	run("/bin/ps", false)
}

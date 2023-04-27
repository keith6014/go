package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type CLI struct {
	Args    string `arg:"" name:"path" help:"paths and paths" help:my args type:"path"`
	Verbose bool   `help:"Verbose"`
}

func main() {
	cli := CLI{}

	_ = kong.Parse(&cli)

	fmt.Println(cli.Args)
	if cli.Verbose {
		fmt.Println("verbose set")
	}
}

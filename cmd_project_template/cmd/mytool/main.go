package main

import "fmt"
import "gopkg.in/alecthomas/kingpin.v2"
import "local/mylib"

var (
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	name    = kingpin.Arg("mytool", "Tool").Required().String()
)

func main() {
	fmt.Println(mylib.Output())
	kingpin.Parse()
	fmt.Printf("%v,%s\n", *verbose, *name)
}

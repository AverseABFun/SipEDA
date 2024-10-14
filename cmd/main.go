package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/averseabfun/sipeda"
)

var commands []string

func main() {
	flag.Parse()

	commands = flag.Args()

	var f, err = os.Open("../example/example.sch.sip")
	if err != nil {
		panic(err)
	}
	l, err := sipeda.ParseFile(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(sipeda.ValidateFile(l))
}

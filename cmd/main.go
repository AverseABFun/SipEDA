package main

import (
	"fmt"
	"os"

	"github.com/averseabfun/sipeda"
)

func main() {
	var f, err = os.Open("../example/example.fop.sip")
	if err != nil {
		panic(err)
	}
	l, err := sipeda.ParseFile(f)
	fmt.Printf("%#v, %#v\n", l, err)
}

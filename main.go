package main

import (
	"fmt"

	wf "github.com/bwiggs/go-weatherflow/weatherflow"
)

func main() {
	fmt.Println("starting weatherflow listener")
	wf.Listen()
}

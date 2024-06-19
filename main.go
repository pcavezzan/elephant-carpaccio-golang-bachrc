package main

import (
	"log"

	"github.com/bachrc/test-stonal/cmd"
)


func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
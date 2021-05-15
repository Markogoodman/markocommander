package main

import (
	"log"

	"github.com/Markogoodman/markocommander/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}

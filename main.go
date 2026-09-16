package main

import (
	"log"

	"github.com/ioanrobertrosu/todo/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
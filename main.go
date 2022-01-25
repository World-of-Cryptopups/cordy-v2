package main

import (
	"log"

	"github.com/World-of-Cryptopups/cordy-v2/commands"
	"github.com/World-of-Cryptopups/minidis"
)

func main() {
	// start bot
	if err := minidis.Execute(commands.Bot); err != nil {
		log.Fatal(err)
	}
}

package commands

import (
	"os"

	"github.com/World-of-Cryptopups/minidis"
)

var Bot *minidis.Minidis

func init() {
	Bot = minidis.New(os.Getenv("TOKEN"))
}

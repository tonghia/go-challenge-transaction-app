package main

import (
	"log"
	"os"

	"github.com/tonghia/go-challenge-transaction-app/internal/app"
)

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

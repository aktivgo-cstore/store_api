package main

import (
	"log"
	"store_api/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"

	"github.com/exdev-studio/requests-dashboard-api/internal/apiserver"
)

func main() {
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}

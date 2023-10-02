package main

import (
	"log"

	"github.com/pavhov/packs-calculate-test/api"
)

func main() {
	log.Fatal(api.StartServer())
}

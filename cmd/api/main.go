package main

import (
	"log"

	"github.com/chars-mc/encryptor-api/cmd/api/bootstrap"
)

func main() {
	log.Fatal(bootstrap.Start())
}

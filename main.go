package main

import (
	"fmt"
	"log"
	"os"

	"github.com/blockcypher/gobcy"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize godotenv for reading secrets stored in .env files.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	btc := gobcy.API{os.Getenv("TOKEN"), "btc", "main"}
	chain, err := btc.GetChain()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("%+v\n", chain)
}

package main

import (
	"log"

	"github.com/alireza12prom/SimpleChain/internal/api"
	"github.com/alireza12prom/SimpleChain/internal/domain"
	"github.com/dgraph-io/badger/v4"
)

func main() {
	// Database Initialization
	db, err := badger.Open(badger.DefaultOptions(".db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Blockchain Initialization
	blockchain := domain.NewBlockchain(db)
	blockchain.Init()

	// Server
	api.Run(blockchain)
}

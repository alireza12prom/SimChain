package main

import (
	"log"

	"github.com/alireza12prom/SimpleChain/internal/core"
	"github.com/alireza12prom/SimpleChain/internal/server"
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
	blockchain := core.NewBlockchain(db)
	blockchain.Init()

	// Server
	server.Run(blockchain)
}

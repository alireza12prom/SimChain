package main

import (
	"log"

	"github.com/alireza12prom/SimpleChain/internal/core"
	"github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
)

func main() {
	// Database Initialization
	db, err := badger.Open(badger.DefaultOptions("./simchain.db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Blockchain Initialization
	blockchain := core.NewBlockchain(db)
	blockchain.Init()

	// Server
	r := gin.Default()

	{
		g := r.Group("/transaction")
		g.POST("/new", func(c *gin.Context) {})
	}

	r.Run()
}

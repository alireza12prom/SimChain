package main

import (
	"log"
	"net/http"

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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

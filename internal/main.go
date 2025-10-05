package main

import (
	"fmt"
	"log"

	"github.com/alireza12prom/SimpleChain/internal/core"
	"github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
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
	r := gin.Default()

	{
		g := r.Group("/transaction")

		type NewTransactionDto struct {
			From   string  `json:"from" binding:"required"`
			To     string  `json:"to" binding:"required"`
			Amount float64 `json:"amount" binding:"required"`
		}

		g.POST("/new", func(c *gin.Context) {
			var body NewTransactionDto
			if err := c.ShouldBindBodyWithJSON(&body); err != nil {
				c.JSON(400, `invalid input.`)
			}

			trx := core.NewTransaction(body.From, body.To, body.Amount)
			blockchain.AddTransaction(trx)

			c.JSON(200, gin.H{"msg": "transaction added to the pending pool."})
		})

		g.GET("/pending", func(c *gin.Context) {
			pending := blockchain.GetPendingTransaction()
			c.JSON(200, gin.H{"data": pending, "count": len(pending)})
		})
	}

	{
		r.GET("/history", func(c *gin.Context) {
			history := blockchain.GetBlocks()
			c.JSON(200, gin.H{"data": history, "count": len(history)})
		})

		r.POST("/mine", func(c *gin.Context) {
			block, err := blockchain.MineBlock()
			if err != nil {
				c.JSON(400, gin.H{"reason": err.Error()})
				return
			}

			c.JSON(200, gin.H{
				"msg": fmt.Sprintf("block #%d with %d transactions.", block.Index, block.Size()),
			})
		})
	}

	r.Run()
}

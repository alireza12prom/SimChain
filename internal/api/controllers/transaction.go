package controllers

import (
	"github.com/alireza12prom/SimpleChain/internal/domain"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	Blockchain *domain.Blockchain
}

func (tc *TransactionController) Create(c *gin.Context) {
	var body NewTransactionDto
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, `invalid input.`)
	}

	trx := domain.NewTransaction(body.From, body.To, body.Amount)
	tc.Blockchain.AddTransaction(trx)

	c.JSON(200, gin.H{"msg": "transaction added to the pending pool."})
}

func (tc *TransactionController) GetPending(c *gin.Context) {
	pending := tc.Blockchain.GetPendingTransaction()
	c.JSON(200, gin.H{"data": pending, "count": len(pending)})
}

type NewTransactionDto struct {
	From   string  `json:"from" binding:"required"`
	To     string  `json:"to" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

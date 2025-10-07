package controllers

import (
	"fmt"

	"github.com/alireza12prom/SimpleChain/internal/core"
	"github.com/gin-gonic/gin"
)

type BlockController struct {
	Blockchain *core.Blockchain
}

func (bc *BlockController) GetHistory(c *gin.Context) {
	history := bc.Blockchain.GetBlocks()
	c.JSON(200, gin.H{"data": history, "count": len(history)})
}

func (bc *BlockController) Mine(c *gin.Context) {
	block, err := bc.Blockchain.MineBlock()
	if err != nil {
		c.JSON(400, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"msg": fmt.Sprintf("block #%d with %d transactions.", block.Index, block.Size()),
	})
}

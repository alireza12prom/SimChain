package main

import (
	"github.com/alireza12prom/SimpleChain/internal/api"
	"github.com/alireza12prom/SimpleChain/internal/blockchain"
	"github.com/alireza12prom/SimpleChain/internal/domain"
)

func main() {
	// Blockchain Initialization
	blockStore := blockchain.NewBadgerStore("./.db")
	defer blockStore.Close()

	blockchainConfig := domain.BlockchainConfig{
		Difficulty:   4,
		MaxBlockSize: 10,
	}

	blockchain := blockchain.NewBlockchain(blockchainConfig, blockStore)

	// Server
	api.Run(blockchain)
}

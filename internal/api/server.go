package api

import (
	"github.com/alireza12prom/SimpleChain/internal/api/controllers"
	"github.com/alireza12prom/SimpleChain/internal/domain"
	"github.com/gin-gonic/gin"
)

func Run(blockchain domain.IBlockchain) {
	r := gin.Default()

	{
		controller := controllers.TransactionController{
			Blockchain: blockchain,
		}

		g := r.Group("/transaction")
		g.POST("/new", controller.Create)
		g.GET("/pending", controller.GetPending)
	}

	{
		controller := controllers.BlockController{
			Blockchain: blockchain,
		}

		g := r.Group("/block")
		g.GET("/history", controller.GetHistory)
		g.POST("/create", controller.Create)
	}

	r.Run(":8081")
}

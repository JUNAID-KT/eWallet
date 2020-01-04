package router

import (
	"github.com/JUNAID-KT/eWallet/controller"

	"github.com/JUNAID-KT/eWallet/util"
	"github.com/gin-gonic/gin"
)

// Set routers
func SetRoutes(router *gin.Engine) *gin.Engine {

	v1 := router.Group(util.ApiVersion)
	{
		wallet := v1.Group(util.ApiPrefix)
		{
			// routes for transaction data
			wallet.GET(util.TransactionsApi, controller.GetTransactions)
		}
	}
	return router
}

package controller

import (
	"net/http"

	"github.com/JUNAID-KT/eWallet/util"
	"github.com/gin-gonic/gin"
)

func GetTransactions(context *gin.Context) {

	context.JSON(http.StatusOK, util.SetStatus(http.StatusOK, util.SuccessDesc, "resource data fetched"))
}

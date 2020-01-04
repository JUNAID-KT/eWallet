package controller

import (
	"net/http"

	"github.com/JUNAID-KT/eWallet/models"
	se "github.com/JUNAID-KT/eWallet/search_engine"
	"github.com/JUNAID-KT/eWallet/util"
	"github.com/gin-gonic/gin"
)

var GetTransactin = "GetTransactions"

// Handler for getting transactions
func GetTransactions(context *gin.Context) {
	var apiRequest models.TransactionByUser
	if err := context.BindJSON(&apiRequest); err != nil {
		util.ErrorResponder(err, GetTransactin, util.FailureDesc, util.BindingFailedMsg, http.StatusBadRequest, context)
		return
	}
	es := se.GetESInstance()

	err, transactions := es.GetTransactions(apiRequest.User)
	if err != nil {
		util.ErrorResponder(err, GetTransactin, util.FailureDesc, err.Error(), http.StatusBadRequest, context)
		return
	}
	context.JSON(http.StatusOK, models.ListTransactionResponse{
		Status: util.SetStatusResponse(http.StatusOK,
			http.StatusText(http.StatusOK),
			"Transactions fetched"),
		Data: transactions,
	})
}

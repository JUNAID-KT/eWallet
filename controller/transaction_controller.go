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
	var request models.TransactionByUser
	es := se.GetESInstance()
	if err := context.ShouldBindJSON(&request); err != nil {
		util.ErrorResponder(err, GetTransactin, util.FailureDesc, util.BindingFailedMsg, http.StatusBadRequest, context)
		return
	}
	if err := es.Validate.Struct(request); err != nil {
		util.ErrorResponder(err, GetTransactin, util.FailureDesc, util.ValidationFailedMsg, http.StatusBadRequest, context)
		return
	}

	err, transactions := es.GetTransactions(request.User)
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

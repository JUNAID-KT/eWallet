package util

const (
	ApiV1           = "/v1.0"
	ApiVersion      = ApiV1
	ApiPrefix       = "/eWallet"
	TransactionsApi = "transaction"
	MaxRetries      = 5
	SearchLimit     = 10000

	FailureDesc = "FAILURE"
	SuccessDesc = "SUCCESS"

	TransactionIndexName = "ethereum_transactions"
	TransactionTypeName  = "transactions"
	ValidationFailedMsg  = "Invalid Request"
	BindingFailedMsg     = "Request binding failed"
)

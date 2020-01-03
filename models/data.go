package models

type Transaction struct {
	From, To        string
	BlockNumber     uint64
	TransactionHash string
}
type ListTransactionResponse struct {
	Status StatusResponse
	Data   []Transaction
}

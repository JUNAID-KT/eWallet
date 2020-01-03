package models

type Transaction struct {
	From            string `json:"from"`
	To              string `json:"to"`
	BlockNumber     uint64 `json:"block_number"`
	TransactionHash string `json:"transaction_hash"`
}
type ListTransactionResponse struct {
	Status StatusResponse `json:"status_response"`
	Data   []Transaction  `json:"data"`
}

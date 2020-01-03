package worker

import (
	"context"

	"github.com/JUNAID-KT/eWallet/models"
	se "github.com/JUNAID-KT/eWallet/search_engine"
	log "github.com/Sirupsen/logrus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	SaveTrans = "SaveTransactions"
	GetBlock  = "GetBlocks"
)

// Initialize worker
func Init() {
	var blocks = make(chan *types.Block)
	go SaveTransactions(blocks)
	go GetBlocks(blocks)
}

// Fetch Ethereum blocks
func GetBlocks(blocks chan *types.Block) {
	defer close(blocks)
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.WithFields(log.Fields{"method": GetBlock, "description": err.Error()})
		return
	}
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.WithFields(log.Fields{"method": GetBlock, "description": err.Error()})
		return
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			blocks <- block

		}
	}
}
func SaveTransactions(blocks chan *types.Block) {
	elasticClient := se.GetESInstance()
	var transaction models.Transaction
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.WithFields(log.Fields{"method": SaveTrans, "description": err.Error()})
		return
	}
	for block := range blocks {
		transaction.BlockNumber = block.Number().Uint64()
		for _, tx := range block.Transactions() {
			transaction.To = tx.To().Hex()
			transaction.TransactionHash = tx.Hash().Hex()
			chainID, err := client.NetworkID(context.Background())
			if err != nil {
				log.WithFields(log.Fields{"method": SaveTrans, "description": err.Error()})
				return
			}
			if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
				transaction.From = msg.From().Hex()
			}
			err = elasticClient.SaveTransactions(transaction)
			if err != nil {
				log.WithFields(log.Fields{"method": SaveTrans, "description": err.Error()})
				return
			}
		}
	}
}

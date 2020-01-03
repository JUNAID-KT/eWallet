package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Init() {
	//var wg sync.WaitGroup
	var blocks = make(chan *types.Block)
	//wg.Add(2)
	go SaveData(blocks)
	go GetBlocks(blocks)
	fmt.Println("worker started..................")
	//wg.Wait()
	//GetTrans(blocks)
	//time.Sleep(time.Second * 1000)
}
func GetBlocks(blocks chan *types.Block) {
	defer close(blocks)
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("BLOCK")
			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
			blocks <- block
			//SaveData(block)
		}
	}
}
func SaveData(blocks chan *types.Block) {
	for block := range blocks {
		for _, tx := range block.Transactions() {
			fmt.Println("TRANSACTION")
			fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
			fmt.Println(tx.Value().String())    // 10000000000000000
			fmt.Println(tx.Gas())               // 105000
			fmt.Println(tx.GasPrice().Uint64()) // 102000000000
			fmt.Println(tx.Nonce())             // 110644
			fmt.Println(tx.Data())              // []
			fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e
		}
	}
}

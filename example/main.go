package main

import (
	"log"
	"os"

	"github.com/go-develop/coinapi/config"

	"github.com/go-develop/coinapi/wallet"
)

func main() {

	factory := wallet.Factory{Environment: config.Env_DEVELOPMENT}
	client := factory.Get(config.CoinId_USDT)

	//address := client.NewAddress("")
	//log.Printf("钱包地址：%s\n", address)

	balance := client.GetBalance("mipvXNuvsoD9JvxCiG6MeKB4noyp3e88ME")
	log.Printf("查询余额：%s\n", balance)

	//BTC手续费：mt8Lvqmik6w4ZimnqDb8pKUupYyXFpjBQJ
	txid := client.SendTransaction("mipvXNuvsoD9JvxCiG6MeKB4noyp3e88ME", "mgBg7ZRXaxybbCQeFqKnGXvVCECoPV4LYb", "0.005", "mt8Lvqmik6w4ZimnqDb8pKUupYyXFpjBQJ")
	log.Println(txid)

	tx := client.GetTxById("0ef9b5e05f68480dd745768bac73382d2b5ee2f2c914bcec51f0d79e05374707")
	log.Println(tx)

	index := client.GetBlockCount()
	log.Println(index)

	txs := client.GetBlockTxs(1560835)
	log.Println(txs)
	log.Println(txs[0])

	txs = client.GetPendingTxs("")
	log.Println(txs)

	println(os.Getenv("HOME"))
	println(os.Getenv("PATH"))
	println(os.Getenv("GOPATH"))

}

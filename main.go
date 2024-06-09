package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// api for use external blockchain
var infuraURL = "https://mainnet.infura.io/v3/de18283a1cd8402da6ec6c8a7148d2f9"

// local blockcahin
var ganacheURL = "http://127.0.0.1:8545"
 
func main(){

	// connect to network
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("Error to create a eth client: %v", err)
	}
	defer client.Close()

	// get info
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get the block: %v", err)
	}
	fmt.Println(block.Number())

	// specific wallet
	addr := "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"
	address := common.HexToAddress(addr)

	// check balance
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance: %v", err)
	}
	fmt.Println(balance)

	// whei --> eht
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println(fBalance)

	balanceEth := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(balanceEth)
}
















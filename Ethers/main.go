package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const url_key string = "https://sepolia.infura.io/v3/77ad6d7f679f4787876bb366224c9a4a"

func TaskOne_BlockWR(senderPk string, toAddr string) {
	client, err := ethclient.Dial(url_key)
	if err != nil {
		panic(err)
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	blockNum := big.Int(*header.Number)
	block, err := client.BlockByNumber(context.Background(), &blockNum)
	if err != nil {
		panic(err)
	}

	fmt.Println("block hash :" + block.Hash().Hex())
	fmt.Println("block time :" + strconv.FormatUint(block.Time(), 10))
	fmt.Println("block transaction count :" + strconv.FormatInt(int64(block.Transactions().Len()), 10))

	fmt.Println("--------- Send A Transaction ---------")

	pKey, err := crypto.HexToECDSA(senderPk) //在此设置发送方	!!!私钥!!!
	if err != nil {
		panic(err)
	}
	pbKey := pKey.Public()
	pbKeyECDSA, ok := pbKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type:Public Key")
	}
	addr := crypto.PubkeyToAddress(*pbKeyECDSA)
	fmt.Println("sender :" + addr.Hex())
	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		panic(err)
	}
	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	gasLimit := 21000
	targetAddr := common.HexToAddress(toAddr) //在此设置 ETH接收方
	fmt.Println("receiver :" + targetAddr.Hex())
	//创建交易数据
	tx := types.NewTransaction(nonce, targetAddr, big.NewInt(10000000000000000), uint64(gasLimit), gasprice, nil)

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	//交易签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), pKey)
	if err != nil {
		panic(err)
	}
	//发送交易
	client.SendTransaction(context.Background(), signedTx)
	fmt.Println("TX Hash :" + signedTx.Hash().Hex())
}

func main() {
	//privateKey := ""
	//toAddr := ""

	//TaskOne_BlockWR(privateKey, toAddr)

	//TaskTwo_ContractInteractive(privateKey)	//deploy

	contractHex := "0x9c3F59F0D1CB7Ea9eb295bfB7Aa9a893c3a1BCb3" //部署后的合约地址
	//ContractDataSet(privateKey, contractHex, "contract interaction test") //write
	ContractDataGet(contractHex) //Read
}

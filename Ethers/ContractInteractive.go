package main

import (
	. "Ethers/contracts/EthDeploy"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TaskTwo_ContractInteractive(senderPK string) {
	client, err := ethclient.Dial(url_key)
	if err != nil {
		panic(err)
	}
	pKey, err := crypto.HexToECDSA(senderPK) //在此设置发送方	!!!私钥!!!
	if err != nil {
		panic(err)
	}
	pbKey := pKey.Public()
	pbKeyECDSA, ok := pbKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type:Public Key")
	}
	addr := crypto.PubkeyToAddress(*pbKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pKey, chainId)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300300)
	auth.GasPrice = gasPrice

	address, tx, instance, err := DeployEthDeploy(auth, client)
	if err != nil {
		panic(err)
	}
	fmt.Println(address.Hex())
	fmt.Println(tx.Hash())
	_ = instance
}
func ContractDataSet(senderPK string, contractHex string, dataValue string) {
	client, err := ethclient.Dial(url_key)
	if err != nil {
		panic(err)
	}
	pKey, err := crypto.HexToECDSA(senderPK)
	if err != nil {
		panic(err)
	}
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	opt, err := bind.NewKeyedTransactorWithChainID(pKey, chainId)
	if err != nil {
		panic(err)
	}

	contractAddr := common.HexToAddress(contractHex)
	ethdeploy, err := NewEthDeploy(contractAddr, client)
	if err != nil {
		panic(err)
	}
	tx, err := ethdeploy.SetData(opt, dataValue)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx hash :" + tx.Hash().Hex())
}

func ContractDataGet(contractHex string) {
	client, err := ethclient.Dial(url_key)
	if err != nil {
		panic(err)
	}
	contractAddr := common.HexToAddress(contractHex)
	ethdeploy, err := NewEthDeploy(contractAddr, client)
	if err != nil {
		panic(err)
	}

	callOpt := &bind.CallOpts{Context: context.Background()}
	dataValue, err := ethdeploy.GetData(callOpt)
	if err != nil {
		panic(err)
	}

	fmt.Println("contract data value :" + dataValue)
}

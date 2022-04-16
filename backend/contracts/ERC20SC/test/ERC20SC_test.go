package snakeSCtest

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/phuongdoan13/BlockchainGame/backend/contracts"
	ERC20SC "github.com/phuongdoan13/BlockchainGame/backend/contracts/ERC20SC"
	"log"
	"math/big"
	"testing"
)

func TestDeploy(t *testing.T) {
	deployerPrivateKey := "d96fcaba769f8413ab98a7b7096e0830aa92daba147e7475075c04ff259aa3c7"
	client := contracts.GetGanacheClient()

	privateKey, err := crypto.HexToECDSA(deployerPrivateKey)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(10000)

	address, tx, instance, err := ERC20SC.DeployERC20SC(auth, client, big.NewInt(1000))
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())

	_, _ = instance, tx

	if 0 == 1 {
		t.Errorf("Dummy test case")
	}
}

// Test case: Test BalanceOf function
func TestBalanceOf(t *testing.T) {
	contractAddress := "0xd782EE5f0aeD35787000b3953a46C087691847AC"
	publicAddress := "0x0b843Dd4c33883d6c191F68279C326fc105644FE"

	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contractAddress)
	instance, err := ERC20SC.NewERC20SC(address, client)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(instance.TotalSupply(nil))
	userPublicAddress := common.HexToAddress(publicAddress)
	balance, err := instance.BalanceOf(nil, userPublicAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
	//assert.Equal(t, balance, big.NewInt(10), "Symbol did not match!")

}

// Test case: Test BalanceOf function
func TestTransfer(t *testing.T) {
	contractAddress := "0xd782EE5f0aeD35787000b3953a46C087691847AC"
	publicAddress := "0x0b843Dd4c33883d6c191F68279C326fc105644FE"
	deployerPrivateKey := "d96fcaba769f8413ab98a7b7096e0830aa92daba147e7475075c04ff259aa3c7"
	// Get connection
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}
	// Get smart contract
	address := common.HexToAddress(contractAddress)
	instance, err := ERC20SC.NewERC20SC(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Get transaction auth
	privateKey, err := crypto.HexToECDSA(deployerPrivateKey)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	// Do transfer
	recipientPublicAddress := common.HexToAddress(publicAddress)
	transaction, err := instance.Transfer(auth, recipientPublicAddress, big.NewInt(200))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transaction)
}

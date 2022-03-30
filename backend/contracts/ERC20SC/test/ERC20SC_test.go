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

	client := contracts.GetGanacheClient()

	privateKey, err := crypto.HexToECDSA("db454228402b38f7e0ce9e6df42d7b04c640fae680dad83c06b8dea9d1b3c4ed")
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
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x0A4c3AD7cE0c28CF3AD65B2E5643bb5a5043Ee44")
	instance, err := ERC20SC.NewERC20SC(address, client)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(instance.TotalSupply(nil))
	userPublicAddress := common.HexToAddress("0x8a2a3a1dacF2B4b57734eB3DB71c33d3EBe263B6")
	balance, err := instance.BalanceOf(nil, userPublicAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
	//assert.Equal(t, balance, big.NewInt(10), "Symbol did not match!")

}

// Test case: Test BalanceOf function
func TestTransfer(t *testing.T) {
	// Get connection
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}
	// Get smart contract
	address := common.HexToAddress("0x0A4c3AD7cE0c28CF3AD65B2E5643bb5a5043Ee44")
	instance, err := ERC20SC.NewERC20SC(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// Get transaction auth
	privateKey, err := crypto.HexToECDSA("db454228402b38f7e0ce9e6df42d7b04c640fae680dad83c06b8dea9d1b3c4ed")
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
	recipientPublicAddress := common.HexToAddress("0x941B2bb2CA1E6aF49A728Ed0f0556cc070b23AF7")
	transaction, err := instance.Transfer(auth, recipientPublicAddress, big.NewInt(10))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transaction)
}

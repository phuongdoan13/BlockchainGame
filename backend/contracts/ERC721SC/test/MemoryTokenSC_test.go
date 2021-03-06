package MemoryTokenSC_test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/phuongdoan13/BlockchainGame/backend/contracts"
	"github.com/phuongdoan13/BlockchainGame/backend/contracts/ERC721SC"
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"math/big"
	"testing"
)

func TestDeploy(t *testing.T) {

	client := contracts.GetGanacheClient()

	privateKey, err := crypto.HexToECDSA("3b79e47859611d8d8a87f06c94c84c5a0f8d9f96f13901a1dccaa9fa336daa10")
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

	address, tx, instance, err := ERC721SC.DeployERC721SC(auth, client, "Memory Token", "MEM")
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())

	_, _ = instance, tx

	if 0 == 1 {
		t.Errorf("Dummy test case")
	}
}

// Test case: Test if the instance of the smart contract has the name "Memory Token"
func TestSmartContractHasName(t *testing.T) {
	client := contracts.GetGanacheClient()

	address := common.HexToAddress("0xbF08c6CEcC51A8af59a07e7028076A2fb69Eb8E5")
	instance, err := ERC721SC.NewERC721SC(address, client)
	if err != nil {
		log.Fatal(err)
	}
	tokenName, err := instance.Name(nil)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, tokenName, "Memory Token", "Name did not match")
	fmt.Println("Name: ", tokenName)
}

// Test case: Test if the instance of the smart contract has the symbol "MEM"
func TestSmartContractHasSymbol(t *testing.T) {
	client := contracts.GetGanacheClient()

	address := common.HexToAddress("0xbF08c6CEcC51A8af59a07e7028076A2fb69Eb8E5")
	instance, err := ERC721SC.NewERC721SC(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tokenSymbol, err := instance.Symbol(nil)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, tokenSymbol, "MEM", "Symbol did not match!")
	fmt.Println("Symbol: ", tokenSymbol)
}

// Test case: After deploying a contract,
// the account balance should be smaller than the default 100
// (Unless it joins the game)
func TestAccountBalanceAfterSmartContractDeploy(t *testing.T) {
	// Get the Ethereum client
	client := contracts.GetGanacheClient()

	// Get the deployer's balance after the deployment
	blockNumber := big.NewInt(1)
	account := common.HexToAddress("0x2509942d6be3260cb09bD4d5Bc90bf628B502531")
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		t.Fatal(err)
	}

	// Get the deployer's balance before the deployment and compare with that after the deployment
	var bignum, _ = new(big.Int).SetString("100000000000000000000", 0)
	if balance.Cmp(bignum) >= 0 {
		t.Errorf("Should be smaller than 100.00 Eth")
	}

	fmt.Println(new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18))))
}

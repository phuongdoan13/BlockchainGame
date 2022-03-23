package MemoryTokenSC_test

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"math/big"
	"testing"
)

func TestDeploy(t *testing.T) {

	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}

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

	address, tx, instance, err := ERC721SC.DeployMatching(auth, client, "Memory Token", "MEM")
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
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x3903895653567B95D2DcCff3dAfBF118d22d4a5d")
	instance, err := ERC721SC.NewMatching(address, client)
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
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x3903895653567B95D2DcCff3dAfBF118d22d4a5d")
	instance, err := ERC721SC.NewMatching(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tokenSymbol, err := instance.Symbol(nil)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, tokenSymbol, "MEM", "Symbol did not match!")
}

// Test case: After deploying a contract,
// the account balance should be smaller than the default 100
// (Unless it joins the game)
func TestAccountBalanceAfterSmartContractDeploy(t *testing.T) {
	// Get the Ethereum client
	_, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		//t.Errorf(err)
		t.Fatal(err)
	}

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

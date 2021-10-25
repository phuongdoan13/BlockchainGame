package matchingSCtest

import (
	"context"
	"github.com/phuongdoan13/BlockchainGame/backend/contracts/matching"
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"math/big"

	//"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	//"github.com/phuongdoan13/BlockchainGame/backend/contracts/matching"
	//"math/big"
	"testing"
)




// Test case: After deploying a contract,
// the account balance should be smaller than the default 100
// (Unless it joins the game)
func TestAccountBalanceAfterSmartContractDeploy(t *testing.T){
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		//t.Errorf(err)
		t.Fatal(err)
	}
	blockNumber := big.NewInt(1)
	account := common.HexToAddress("0x2509942d6be3260cb09bD4d5Bc90bf628B502531")
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		t.Fatal(err)
	}

	var bignum, _ = new(big.Int).SetString("100000000000000000000", 0)
	if balance.Cmp(bignum) >= 0{
		t.Errorf("Should be smaller than 100.00 Eth")
	}

	fmt.Println(new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18))))
}

// Test case: Test if the instance of the smart contract has the name "Memory Token"
func TestSmartContractHasName(t *testing.T){
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x3903895653567B95D2DcCff3dAfBF118d22d4a5d")
	instance, err := matching.NewMatching(address, client)
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
func TestSmartContractHasSymbol(t *testing.T){
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x3903895653567B95D2DcCff3dAfBF118d22d4a5d")
	instance, err := matching.NewMatching(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tokenSymbol, err := instance.Symbol(nil)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, tokenSymbol, "MEM", "Symbol did not match!")
}


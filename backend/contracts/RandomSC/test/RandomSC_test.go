package RandomSC_test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	"github.com/phuongdoan13/BlockchainGame/backend/contracts"
	RandomSC "github.com/phuongdoan13/BlockchainGame/backend/contracts/RandomSC"
	"log"
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

	address, tx, instance, err := RandomSC.DeployRandomSC(auth, client)
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
func TestRandMod(t *testing.T) {
	client := contracts.GetGanacheClient()

	address := common.HexToAddress("0x297D52321Ac45E1fF3E5d472AD74B5FdAe62B2d1")
	instance, err := RandomSC.NewRandomSC(address, client)
	if err != nil {
		log.Fatal(err)
	}

	randChoice, err := instance.RandMod(nil, big.NewInt(132))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(randChoice)

}
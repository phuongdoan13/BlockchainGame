package api

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/phuongdoan13/BlockchainGame/backend/contracts"
	"github.com/phuongdoan13/BlockchainGame/backend/contracts/ERC20SC"
	"log"
	"math/big"
	"net/http"
)

func WinERC20_controller(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	var p winERC20_RequestBody
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(p)
	// Perform transaction
	err = winERC20_implementer(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Success"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func winERC20_implementer(p winERC20_RequestBody) error {
	// PERFORM Transfer() method of ERC20 contract
	contractAddress := "0xDFe2c2493a93fb50Da5bD51Adf846c9fF58ba1Ef"
	publicAddress := p.PublicAddress
	deployerPrivateKey := "d96fcaba769f8413ab98a7b7096e0830aa92daba147e7475075c04ff259aa3c7"
	amount := p.Amount

	// Get connection
	client := contracts.GetGanacheClient()

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
	transaction, err := instance.Transfer(auth, recipientPublicAddress, &amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transaction)
	return nil
}

type winERC20_RequestBody struct {
	PublicAddress string
	Amount        big.Int
}

func enableCors(w *http.ResponseWriter) {
	header := (*w).Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
}

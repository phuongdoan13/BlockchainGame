package api

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
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

	// Get connection
	client := contracts.GetGanacheClient()

	// Get smart contract
	address := common.HexToAddress("0x0A4c3AD7cE0c28CF3AD65B2E5643bb5a5043Ee44")
	instance, err := ERC20SC.NewERC20SC(address, client)
	if err != nil {
		return err
	}

	// Get transaction auth
	privateKey, err := crypto.HexToECDSA("db454228402b38f7e0ce9e6df42d7b04c640fae680dad83c06b8dea9d1b3c4ed")
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	// Do transfer
	recipientPublicAddress := common.HexToAddress(p.PublicAddress)
	_, err = instance.Transfer(auth, recipientPublicAddress, &p.Amount)
	if err != nil {
		return err
	}

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

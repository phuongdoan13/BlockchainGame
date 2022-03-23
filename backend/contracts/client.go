package contracts

import "github.com/ethereum/go-ethereum/ethclient"

func GetGanacheClient() *ethclient.Client {
	// Get connection
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		panic(err)
	}
	return client
}

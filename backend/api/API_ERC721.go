package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MatchingContractABI(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		b, err := ioutil.ReadFile("/Users/phuongdoantuminh/Projects/BlockchainManagement/backend/api/util/MatchingSC.json")
		if err != nil {
			fmt.Println(err)
		}
		rawIn := json.RawMessage(string(b))
		var objmap map[string]*json.RawMessage
		err_json := json.Unmarshal(rawIn, &objmap)
		if err_json != nil {
			fmt.Println(err_json)
		}

		json.NewEncoder(w).Encode(objmap)
		fmt.Println(objmap)
	}
}

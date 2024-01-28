package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getWinners(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("INVALID METHOD getWinners")
		return
	}

	payload, err := json.Marshal(map[string]interface{}{
		"contractAddress": contractAddress,
		"blockchainNetwork": blockChainId,
		"method": "lastWinners",
		"params": []string{},
	})
	if err != nil {
		log.Println("[ERROR] json body:", err)
		return
	}
	res, err := makeRequest(viewUrl, r.Method, payload)
	if err != nil {
		log.Println("[ERROR] makeRequest:", err)
		return
	}
	w.Write(res)
}


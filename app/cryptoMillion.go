package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	client = &http.Client{};
	blockChainId int
	contractAddress string
	viewUrl string
	mutableUrl string
	address string
)

func main() {
	initEnv()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if Index().Render(context.Background(), w) != nil {
			return
		}
	})
	http.HandleFunc("/connect", connect)

	http.HandleFunc("/balance", getBalance)
	http.HandleFunc("/buyTicket", buyTicket)
	http.HandleFunc("/tickets", getTickets)
	http.HandleFunc("/winners", getWinners)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initEnv() {
	id, err := strconv.Atoi(os.Getenv("BLOCKCHAIN_ID"))
	if err != nil {
		log.Println("[ERROR] blockChainId:", err)
		os.Exit(1)
	}
	blockChainId = id
	contractAddress = os.Getenv("CONTRACT_ADDRESS")
	if contractAddress == "" {
		log.Println("[ERROR] contractAddress is empty")
		os.Exit(1)
	}
	viewUrl = os.Getenv("VOTTUN_VIEW_URL")
	if viewUrl == "" {
		log.Println("[ERROR] VOTTUN_VIEW_URL is empty")
		os.Exit(1)
	}
	mutableUrl = os.Getenv("VOTTUN_MUTABLE_URL")
	if mutableUrl == "" {
		log.Println("[ERROR] VOTTUN_MUTABLE_URL is empty")
		os.Exit(1)
	}
}

func connect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR] read Body:", err)
		return
	}
	defer r.Body.Close()
	log.Println("body:", string(body))
	address = strings.Split(string(body), "=")[1]
	if Content().Render(context.Background(), w) != nil {
		return
	}
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("[ERROR] Method Not Allowed")
		return
	}
	log.Println("GET/ getBalance")

	payload, err := json.Marshal(map[string]interface{}{
		"contractAddress": contractAddress,
		"blockchainNetwork": blockChainId,
		"method": "getBalance",
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
	log.Println("res:", string(res))
	w.Write(res)
}


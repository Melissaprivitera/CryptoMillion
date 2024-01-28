package main

import (
	"encoding/json"
	"log"
	"strconv"
	"net/http"
)

func buyTicket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("INVALID METHOD buyTicket")
		return
	}
	log.Println("POST/ buyTicket")
	r.ParseForm()

	input := [5]string {
		r.FormValue("number1"),
		r.FormValue("number2"),
		r.FormValue("number3"),
		r.FormValue("number4"),
		r.FormValue("number5"),
	}
	numbers := [5]uint8{}
	for i := 0; i < 5; i++ {
		if input[i] == "" {
			log.Println("[ERROR] empty input")
			return
		}
		n, err := strconv.Atoi(input[i])
		if err != nil {
			log.Println("[ERROR] invalid input:", err)
			return
		}
		if n > 40 || n < 1 {
			log.Println("[ERROR] invalid input:", n)
			return
		}
		numbers[i] = uint8(n)
	}
	payload, err := json.Marshal(map[string]interface{}{
		"contractAddress": contractAddress,
		"blockchainNetwork": blockChainId,
		"sender" : address,
		"method": "buy",
		"params": []string{address, strconv.Itoa(int(packNumbers(numbers)))},
	})
	if err != nil {
		log.Println("[ERROR] json body:", err)
		return
	}
	makeRequest(mutableUrl, r.Method, payload)
}

func getTickets(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("INVALID METHOD getTickets")
		return
	}
	log.Println("GET/ getTickets")
	payload, err := json.Marshal(map[string]interface{}{
		"contractAddress": contractAddress,
		"blockchainNetwork": blockChainId,
		"method": "getTickets",
		"params": []string{address},
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


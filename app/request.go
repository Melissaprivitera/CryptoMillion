package main

import (
	"net/http"
	"os"
	"log"
	"strings"
	"io"
)

func makeRequest(url string, method string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))
	if err != nil {
		log.Println("[ERROR] create request:", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer " + os.Getenv("VOTTUN_API_KEY"))
	req.Header.Add("x-application-vkn", os.Getenv("VOTTUN_APP_ID"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR] do request:", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("[ERROR] read Body:", err)
		return nil, err
	}
	log.Println(string(body))
	// TODO Return Respones
	return body, nil
}


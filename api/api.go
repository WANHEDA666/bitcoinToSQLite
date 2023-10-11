package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testProject/internal/storage/bitcoin"
)

func GetBitcoinData() bitcoin.Response {
	response, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject bitcoin.Response
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

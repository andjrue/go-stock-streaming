package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/andjrue/go-stock-streaming/external"
)

func main() {
	symbol := "AAPL"

	external.Init()
	quote, err := external.GetQuote(symbol)
	if err != nil {
		log.Fatal(err)
	}
	q, jsonErr := json.MarshalIndent(quote, "", " ")
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Printf("Quote received successfully: %v", q)

}

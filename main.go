package main

// It will be interesting to see how this runs in the morning when markets are open

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/andjrue/go-stock-streaming/external"
	"github.com/andjrue/go-stock-streaming/internal"
	"github.com/joho/godotenv"
)

func main() {

	symbol := "NFLX" // Make this something the user can enter?

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env: %v", err)
	}

	external.Init()

	pub := internal.NewPublisher()
	sub := pub.Subscribe()

	go func() {
		count := 0 // Need to track requests, I think limited to 60/min
		for update := range sub {
			count++
			fmt.Println(count)
			data, _ := json.Marshal(update)
			fmt.Printf("update received: %s\n", data)
		}
	}()

	for {

		quote, err := external.GetQuote(symbol)
		if err != nil {
			log.Printf("error getting quote: %v", err)
		}

		update := internal.QuoteUpdate{
			Symbol:    symbol,
			Price:     *quote.C,
			OpenPrice: *quote.O,
		}

		pub.Publish(update)
		time.Sleep(1 * time.Second)
	}

}

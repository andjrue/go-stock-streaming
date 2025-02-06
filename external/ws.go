package external

import (
	"context"
	"fmt"
	"log"
	"os"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/joho/godotenv"
)

var Client *finnhub.DefaultApiService

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Issue loading API key: %v", err)
	}

	key := os.Getenv("FINN_API_KEY")
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", key)
	Client = finnhub.NewAPIClient(cfg).DefaultApi
}

func GetQuote(symbol string) (*finnhub.Quote, error) {
	res, _, err := Client.Quote(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		return nil, fmt.Errorf("error fetching quote for %s: %v", symbol, err)
	}
	fmt.Printf("Response: %v\n", &res)
	return &res, nil
}

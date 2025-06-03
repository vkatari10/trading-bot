package api

// This file works with the market API to get real time market data

import (
    finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
    "context"
    "github.com/joho/godotenv"
    "os"
)

// Could use Alpaca Market API if we want to speed it up on 15m delayed data
// for development purposes instead of waiting 30 minutes for burn in

// GetQuote Gets the current price of a given ticker using the 
// Finnhub API
func GetQuote(ticker string) (float32, error) {

    godotenv.Load()
    apiKey := os.Getenv("FINNHUB_API_KEY")

    cfg := finnhub.NewConfiguration()
    cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)
    finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

    res, _, err := finnhubClient.Quote(context.Background()).Symbol(ticker).Execute()
    if err != nil {
        return -1.0, err
    }

    return *res.C, nil // return current price
} // GetQuote
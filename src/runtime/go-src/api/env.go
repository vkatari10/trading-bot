package api

import (
       "github.com/vkatari10/trading-bot/src/runtime/go-src/api"
       "github.com/joho/godotenv"
       "os"
)

// This file contains environment variables that other files depend
// on to interact with external APIs

var (
    finnhubApi string
    finnhubSec string
    alpacaApi string
    alpacaSec string
)


// init loads the environment variable keys
func init() {

     godotenv.Load()

     finnhubApi = os.Getenv("FINNHUB_API_KEY")
     finnhubSec = os.Getenv("FINNHUB_SECRET_KEY")
     alpacaApi = os.Getenv("ALPACA_API_KEY")
     alpacaSec = os.Getenv("ALPACA_SECRET_KEY")

} // init

func alpacaApi() string {
     return alpacaApi
} // alpacaAPI

func alapcaSecret() string {
     return alpacaSec
} // alpacaSecret

func finnhubApi() string {
     return finnhubApi
} // finnhubApi

func finnhubSecret() string {
     return finnhubSec
} // finnhubSecret

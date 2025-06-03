package api

import (
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

func alpacaApiGet() string {
     return alpacaApi
} // alpacaAPI

func alapcaSecretGet() string {
     return alpacaSec
} // alpacaSecret

func finnhubApiGet() string {
     return finnhubApi
} // finnhubApi

func finnhubSecretGet() string {
     return finnhubSec
} // finnhubSecret

package api

import (
       "github.com/joho/godotenv"
       "os"
)

// This file contains environment variables that other files depend
// on to interact with external APIs

var (
    alpacaApi string
    alpacaSec string
)

// init loads the environment variable keys
func init() {
     godotenv.Load()
     alpacaApi = os.Getenv("ALPACA_API_KEY")
     alpacaSec = os.Getenv("ALPACA_SECRET_KEY")
} // init
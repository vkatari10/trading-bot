package alpaca

import (
       "github.com/joho/godotenv"
       "os"
       //"log"
)

// This file contains environment variables that other files depend
// on to interact with external APIs

var (
    alpacaApi string
    alpacaSec string
    alpacaAccountLink string
    alpacaPositionsLink string
    alpacaOrdersLink string
    alpacaClockLink string
    alpacaMarketDataLink string
)

// init loads the environment variable keys in this package
func init() {
     godotenv.Load("../../../.env")
     alpacaApi = os.Getenv("ALPACA_API")
     alpacaSec = os.Getenv("ALPACA_SECRET")
     alpacaAccountLink = os.Getenv("ALPACA_ACCOUNT_LINK")
     alpacaPositionsLink = os.Getenv("ALPACA_POSITIONS_LINK")
     alpacaOrdersLink = os.Getenv("ALPACA_ORDERS_LINK")
     alpacaClockLink = os.Getenv("ALPACA_CLOCK_LINK")
     alpacaMarketDataLink = os.Getenv("ALPACA_MARKET_DATA_LINK")
  
} // init




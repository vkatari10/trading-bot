package eventloop

// Contains methods to get the environment variables 
// for the runtime engine

import (
	"github.com/joho/godotenv"
    "os"
)

var (

	thisTicker string // keep for now remove soon
	// for when multi asset trading is implemented

	logLink string
	dataLink string
	envLink string
	brokerLink string
)

func init() {
	godotenv.Load("../../../.env")

	// Main Settings
	thisTicker = os.Getenv("LIVE_TRADE_TICKER") // Move to JSON later 

	// API Endpoints
	logLink = os.Getenv("LOG_LINK")
	dataLink = os.Getenv("DATA_LINK")
	envLink = os.Getenv("ENV_LINK")
	brokerLink = os.Getenv("BROKER_LINK")

} // init()

package eventloop

// Contains methods to get the environment variables 
// for the runtime engine

import (
	"github.com/joho/godotenv"
    "os"
)

var (
	logLink string
	dataLink string
	envLink string
	brokerLink string
	mlServerLink string
)

func init() {
	godotenv.Load("../../../.env")

	// API Endpoints
	logLink = os.Getenv("LOG_LINK")
	dataLink = os.Getenv("DATA_LINK")
	envLink = os.Getenv("ENV_LINK")
	brokerLink = os.Getenv("BROKER_LINK")
	mlServerLink = os.Getenv("TEST_ML_API_LINK")
} // init()

package eventloop

// Contains methods to get the environment variables 
// for the runtime engine

import (
	"github.com/joho/godotenv"
    "os"
	"fmt"
	"strconv"
	"time"
)

var (
	refreshRate string
	ticker string
	burnWindow string

	logLink string
	dataLink string
	envLink string
	brokerLink string
)

func init() {
	godotenv.Load(".env")
	ticker = os.Getenv("TICKER")
	refreshRate = os.Getenv("REFRESH_RATE")
	burnWindow = os.Getenv("BURN_WINDOW")

	logLink = os.Getenv("LOG_LINK")
	dataLink = os.Getenv("DATA_LINK")
	envLink = os.Getenv("ENV_LINK")
	brokerLink = os.Getenv("BROKER_LINK")
} // init()

func getRefreshRate() (time.Duration, error) {
	rate, err := strconv.Atoi(refreshRate)
	if err != nil {
		return -1, fmt.Errorf("refresh_rate environment variable") // prevent loop from actaully running
	} // if
	return time.Duration(rate), nil
} // getRefreshRate

func getTicker() string {
	return ticker
} // getTicker

func getBurnWindow() (int, error) {
	period, err := strconv.Atoi(burnWindow)
	if err != nil {
		return -1, fmt.Errorf("burn_window environment variable")
	}
	return period, nil
} // getBurnWindow
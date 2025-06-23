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
	bufferFlushTime string

	logToStdio string
	alwaysRun string
	burnInOverride string

	logLink string
	dataLink string
	envLink string
	brokerLink string

	featuresFile string
)

func init() {
	godotenv.Load("../../../.env")

	// Main Settings
	refreshRate = os.Getenv("REFRESH_RATE_TIME")
	ticker = os.Getenv("LIVE_TRADE_TICKER") // Move to JSON later 
	burnWindow = os.Getenv("BURN_IN_WINDOW_TIME")
	bufferFlushTime = os.Getenv("ENGINE_LOG_AOI_FLUSH_TIME")

	// Dev Tools
	logToStdio = os.Getenv("LOG_TO_STDIO")
	alwaysRun = os.Getenv("ALWAYS_RUN")

	// API Endpoints
	logLink = os.Getenv("LOG_LINK")
	dataLink = os.Getenv("DATA_LINK")
	envLink = os.Getenv("ENV_LINK")
	brokerLink = os.Getenv("BROKER_LINK")

	// User Config File
	featuresFile = os.Getenv("FEATURE_CONFIG_FILE")
} // init()

// getRefreshRate returns the refresh rate to use at runtime
func getRefreshRate() (time.Duration, error) {
	rate, err := strconv.Atoi(refreshRate)
	if err != nil {
		return -1, fmt.Errorf("refreshRate environment variable conversion failed") // prevent loop from actaully running
	} // if
	return time.Duration(rate), nil
} // getRefreshRate

// getFlushTime returns the engine log API flush time
func getBufferFlushTime() (time.Duration, error) {
	rate, err := strconv.Atoi(bufferFlushTime)
	if err != nil {
		return -1, fmt.Errorf("bufferFlushTime environment variable conversion failed")
	} // if
	return time.Duration(rate), nil
} // getBufferFlushTime

// getTicker returns the ticker to use at runtime
func getTicker() string {
	return ticker
} // getTicker

// getBurnWindow gets the amount of time to burn in data before 
// starting live execution 
func getBurnWindow() (int, error) {
	period, err := strconv.Atoi(burnWindow)
	if err != nil {
		return -1, fmt.Errorf("burn_window environment variable")
	}
	return period, nil
} // getBurnWindow

// GetFeatureConfigFile gets the config file to use as features during burn 
// in and runtime
func GetFeatureConfigFile() string {
	return featuresFile
} // getMLFeatureFile

// getPrintToStdio gets the boolean value if the engine should
// print logs to stdio
func getPrintToStdio() bool  {
	return logToStdio == "TRUE"
} // getPrintToStdio

// getAlwaysRun gets the boolean value if the engine should
// run even when the market is closed
func getAlwaysRun() bool {
	return alwaysRun == "TRUE"
} // getAlwaysRun

// getBurnInOverride returns if the override process
// to balance technical indicators should be used else
// replaces the burn in process with dummy data
func getBurnInOverride() bool {
	return burnInOverride == "TRUE"
} // getBurnInOverride
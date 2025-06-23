package eventloop

// Contains methods to get the environment variables 
// for the runtime engine

import (
	"github.com/joho/godotenv"
    "os"
	"fmt"
	"strconv"
)

var (
	refreshRate string
	thisTicker string
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


	thisBurnTime int
	thisRefreshRate int
	thisBufferFlushTime int
	thisAlwaysRun bool
	thisBurnInOverride bool
	thisPrintToStdio bool
)

func init() {
	godotenv.Load("../../../.env")

	// Main Settings
	refreshRate = os.Getenv("REFRESH_RATE_TIME")
	thisTicker = os.Getenv("LIVE_TRADE_TICKER") // Move to JSON later 
	burnWindow = os.Getenv("BURN_IN_WINDOW_TIME")
	bufferFlushTime = os.Getenv("ENGINE_LOG_API_FLUSH_TIME")

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

	// Call methods to intiailize non string values
	burnWindow, err := getBurnWindow() // burn in window
	if err != nil {
		go SendPayload(map[string]any{
			"msg": "could not read BURN_IN_WINDOW_TIME, check .env",
		}, logLink)
	} // if

	thisBurnTime = burnWindow

	refRate, err := getRefreshRate() // refresh rate
	if err != nil {
		go SendPayload(map[string]any{
			"msg": "could not read REFRESH_RATE_TIME, check .env",
		}, logLink)
	} // if	

	thisRefreshRate = refRate

	flushTime, err := getBufferFlushTime()
	if err != nil {
		go SendPayload(map[string]any{
			"msg": fmt.Sprintf("%v", err),
		}, logLink)
	} // if

	thisBufferFlushTime = flushTime
	
	runalways := getAlwaysRun()
	thisAlwaysRun = runalways

	print := getPrintToStdio()
	thisPrintToStdio = print

	override := getBurnInOverride()
	thisBurnInOverride = override

} // init()

// getRefreshRate returns the refresh rate to use at runtime
func getRefreshRate() (int, error) {
	rate, err := strconv.Atoi(refreshRate)
	if err != nil {
		return -1, fmt.Errorf("refreshRate environment variable conversion failed") // prevent loop from actaully running
	} // if
	return rate, nil
} // getRefreshRate

// getFlushTime returns the engine log API flush time
func getBufferFlushTime() (int, error) {
	rate, err := strconv.Atoi(bufferFlushTime)
	if err != nil {
		return -1, fmt.Errorf("bufferFlushTime environment variable conversion failed -> %v", err)
	} // if
	return rate, nil
} // getBufferFlushTime

// getBurnWindow gets the amount of time to burn in data before 
// starting live execution 
func getBurnWindow() (int, error) {
	period, err := strconv.Atoi(burnWindow)
	if err != nil {
		return -1, fmt.Errorf("burn_window environment variable")
	}
	return period, nil
} // getBurnWindow

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

package json

import (
	"encoding/json"
	"fmt"
	"os"
)

// castTo helper method to cast values to a certain 
// type 
// (I have a feeling this does not help at all)
func castTo[T any](x any) (T, bool) {
	y, ok := x.(T)
	return y, ok
} // castTo

// RuntimeSettings is a struct to represent user runtime
// settings
type RuntimeSettings struct {
	settingMap map[string]any
	CycleTime int
	BurnTime int
	LogAPIFlushTime int
	LogToStdout bool
	RunAfterClose bool
	OverrideBurnIn bool
	MLAPIRetryCount int
	// Add more items if needed
} // RuntimeSettings

// GetRuntimeSettingsMap gets the runtime setting object from 
// the user JSON config
func GetRuntimeSettings(file string) (RuntimeSettings, error) {

	jsonData, err := os.ReadFile(file)
	if err != nil {
		return RuntimeSettings{}, fmt.Errorf("%v", err) // figure how else to handle this later another way
	} // if

	var jsonMap map[string]any

	err = json.Unmarshal(jsonData, &jsonMap)

	if err != nil {
		return RuntimeSettings{}, fmt.Errorf("%v", err)
	} // if

	runtimeObject, ok := jsonMap["runtime_settings"].(map[string]any) 
	if !ok {
		return RuntimeSettings{}, fmt.Errorf("%v", ok)
	} // if

	var userSettings RuntimeSettings;

	userSettings.settingMap = runtimeObject

	userSettings, err = initializeSettingsMap(userSettings)
	if err != nil {
		return RuntimeSettings{}, fmt.Errorf("%v", err)
	}
	
	return userSettings, nil
} // GetRuntimeSettings

// initializeSettingsMap asserts and returns the values for each 
// setting in the JSON config
func initializeSettingsMap(rs RuntimeSettings) (RuntimeSettings, error) {

	// fmt.Println(rs.settingMap)

	cycleTime, ok := castTo[float64](rs.settingMap["cycle_time"])
	if !ok {
		return RuntimeSettings{}, fmt.Errorf("%s", makeErrorString("cycle_time"))
	} // if
	rs.CycleTime = int(cycleTime)

	burnTime, ok := castTo[float64](rs.settingMap["burn_window_time"])
	if !ok {
		return RuntimeSettings{}, fmt.Errorf("%s", makeErrorString("burn_window_time"))
	} // if 
	rs.BurnTime = int(burnTime)

	flushTime, ok := castTo[float64](rs.settingMap["log_api_flush_time"])
	if !ok {
		return RuntimeSettings{}, fmt.Errorf("%s", makeErrorString("log_api_flush_time"))
	} // if 
	rs.LogAPIFlushTime = int(flushTime)

	logStdout, ok := castTo[bool](rs.settingMap["log_to_stdout"])
	if !ok {
		return RuntimeSettings{}, fmt.Errorf("%s", makeErrorString(("log_to_stdout")))
	} // if
	rs.LogToStdout = logStdout

	alwaysRun, ok := castTo[bool](rs.settingMap["run_after_close"])
	if !ok {
		return RuntimeSettings{}, fmt.Errorf("%s", makeErrorString("run_after_close"))
	} // if
	rs.RunAfterClose = alwaysRun

	jumpBurn, ok := castTo[bool](rs.settingMap["override_burn_in"])
	if !ok {
		return RuntimeSettings{}, fmt.Errorf("%s", makeErrorString(("override_burn_in")))
	} // if
	rs.OverrideBurnIn = jumpBurn

	retryCount, ok := castTo[float64](rs.settingMap["mlapi_retry_count"])
	if !ok {
		return RuntimeSettings{}, fmt.Errorf("%s", makeErrorString("mlapi_retry_count"))
	} // if 
	rs.MLAPIRetryCount = int(retryCount)

	return rs, nil
} // initializeSettingsMap

func makeErrorString(key string) string {
	return fmt.Sprintf("could not parse %s from runtime settings in config file", key)
}
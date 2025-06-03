package engine

// this file contains methods to compute technicals in real time


// To Connect To C library in future (use import C)
/*
#cgo CFLAGS: -I../../c-src/include
#cgo LDFLAGS: -L../../c-src -llive_data -lm
*/


// get user JSON 
// load onto a LiveIndicator struct (
// burn in data for 30 minutes on a float32 array 

// ====================== ^^^ DONE ||| vvv WIP =============================

// put that same array onto each technical indicator in LiveIndicator (goroutines!)
// since the burn in data is just read only for everything we can have every go routine
// just read from it at once without needing mutex (MAKE METHOD)

// cycle that needs to take in a new price every 60 seconds and call update
// on each element of the LiveIndicator Array (MAKE METHOD)

// send back as JSON to ML model somehow and get prediction back from the 
// Flask local server (MAKE METHODS)

// if we get buy sig utilze broker commands (just 1 share for now)

// need to do something to export this info for the bot API to expose to the frontend
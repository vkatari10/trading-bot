package technicals

// Dispatch table to call TA-Lib C based methods

type talibMethod func(args ...any) any

var (
	talibDispatcher map[string]talibMethod
)

func init() { // Create dispatch table for TA-lib wrapper methods	
	talibDispatcher = map[string]talibMethod{
	
	}
}
package technicals

// Dispatch table to call TA-Lib C based methods

type talibMethod func(tw *TALIBWrapper, technicalType string) (float64, error)

var (
	talibDispatchTable map[string]talibMethod
)

func init() {
	// talibDispatchTable := map[string]talibMethod {
	// 	"SMA": dummy,
	// }
}


/*

Overlap Studies 
- every technical goes to a new method that groups all these together
- within this method is something like a overlapStudies dispatch again 
- that calls another table to put a new method again from TA_XXXX and handles
- certain technicals that have more infromation like bollinger bands
- overlapStudies dispatch and move from there

Whatever is above we just do for every study 
like momentum and bolumne volatiltiy and maybe the pattern recognitions if we can 
make it applicable





*/
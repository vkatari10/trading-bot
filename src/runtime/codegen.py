'''
Codegen methods for TA-Lib wrappers for the Go runtime
module

Author: Vikas Katari
Date: 07/21/2025
'''

# Check for extra ' " '  or spaces
# to run this file just make an args array with function names
# and run a loop and then execute the file
# suggested: redirect stdout to a txt file

def get_moving_avg_timeperiod_only(func: str) -> str:
    print(f""" 
// TA_{func} is a wrapper for the {func} function in TA-Lib
// single return value only inside index 0
func TA_{func}(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {{

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib failed to initialize")
    }}
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {{
        return nil, fmt.Errorf("TA_{func}: Provided objects is not a technicals.FeatureTechnical")
    }}

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_{func}(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Computation failed")
    }}

    if int(outNBElement) == 0 {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Wrote no results")
    }} else {{
        return []float64{{output[outNBElement - 1]}}, nil
    }} 
}} // TA_{func}""")
    
def get_moving_avg_timeperiod_only_high_low_close(func: str) -> str:
    print(f""" 
// TA_{func} is a wrapper for the {func} function in TA-Lib
// single return value only inside index 0
func TA_{func}(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {{ 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib failed to initialize")
    }}
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {{
        return nil, fmt.Errorf("TA_{func}: Provided objects is not a technicals.FeatureTechnical")
    }}

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_{func}(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Computation failed")
    }}

    if int(outNBElement) == 0 {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Wrote no results")
    }} else {{
        return []float64{{output[outNBElement - 1]}}, nil
    }} 
}} // TA_{func}""")

def get_moving_avg_timeperiod_only_high_low(func: str) -> str:
    print(f""" 
// TA_{func} is a wrapper for the {func} function in TA-Lib
// single return value only inside index 0
func TA_{func}(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {{ 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib failed to initialize")
    }}
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {{
        return nil, fmt.Errorf("TA_{func}: Provided objects is not a technicals.FeatureTechnical")
    }}

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_{func}(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        tw.VolumePtr
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Computation failed")
    }}

    if int(outNBElement) == 0 {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Wrote no results")
    }} else {{
        return []float64{{output[outNBElement - 1]}}, nil
    }} 
}} // TA_{func}""")
    
def no_args(func: str) -> str:
    print(f""" 
// TA_{func} is a wrapper for the {func} function in TA-Lib
// single return value only inside index 0
func TA_{func}(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {{ 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib failed to initialize")
    }}
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_{func}(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.OpenPtr,
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        tw.VolumePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Computation failed")
    }}

    if int(outNBElement) == 0 {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Wrote no results")
    }} else {{
        return []float64{{output[outNBElement - 1]}}, nil
    }} 
}} // TA_{func}""")
    

def cycle_indicators(func: str) -> str:
    print(f""" 
// TA_{func} is a wrapper for the {func} function in TA-Lib
// single return value only inside index 0
func TA_{func}(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {{ 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib failed to initialize")
    }}
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    output2 := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))
    out2Ptr := (*C.double)(unsafe.Pointer(&output2[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_{func}(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
        out2Ptr,
    )
    if retCode != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Computation failed")
    }}

    if int(outNBElement) == 0 {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Wrote no results")
    }} else {{
        return []float64{{
            output[outNBElement - 1],
            output2[outNBElement - 1],
            }}, nil
    }} 
}} // TA_{func}""")
    
def pattern_no_args(func: str) -> str:
    print(f""" 
// TA_{func} is a wrapper for the {func} function in TA-Lib
// single return value only inside index 0
func TA_{func}(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {{ 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib failed to initialize")
    }}
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_{func}(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.OpenPtr,
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Computation failed")
    }}

    if int(outNBElement) == 0 {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Wrote no results")
    }} else {{
        return []float64{{float64(output[outNBElement - 1])}}, nil
    }} 
}} // TA_{func}""")
    
def pattern_w_penetration(func: str) -> str:
    print(f""" 
// TA_{func} is a wrapper for the {func} function in TA-Lib
// single return value only inside index 0
func TA_{func}(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {{ 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib failed to initialize")
    }}
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {{
        return nil, fmt.Errorf("TA_{func}: Provided objects is not a technicals.FeatureTechnical")
    }}

    penetration := f.Args["penetration"]

    penetrationArg := C.double(penetration)

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_{func}(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.OpenPtr,
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        penetrationArg,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Computation failed")
    }}

    if int(outNBElement) == 0 {{
        return nil, fmt.Errorf("TA_{func}: TA-Lib Wrote no results")
    }} else {{
        return []float64{{float64(output[outNBElement - 1])}}, nil
    }} 
}} // TA_{func}""")

args = [
    
    "CORREL"    
]


for arg in args:
    get_moving_avg_timeperiod_only(arg)
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
    
def test_func(func: str, i: int) -> str:
    print(f"""
func TestTalib{func}(t *testing.T) {{
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {{
		t.Errorf("failed to create runtimeData Object")
	}}

	eventloop.OverrideBurnIn(runtimeData)

	res, err := technicals.TA_{func}( 
		runtimeData.Objects[{i}],
		&runtimeData.OHLCV,
	)

	if err != nil {{
		t.Errorf("TA {func} failed to run: %v", err)
	}}

	if res == nil {{
		t.Errorf("TA_{func} invalid result came back")
	}}
}} // TestTalib{func}""")
    




ta_functions = [
    "BBANDS", "DEMA", "EMA", "HT_TRENDLINE", "KAMA", "MA", "MAMA", "MIDPOINT",
    "MIDPRICE", "SAR", "SAREXT", "SMA", "T3", "TEMA", "TRIMA", "WMA",

    "ADX", "ADXR", "APO", "AROON", "AROONOSC", "BOP", "CCI", "CMO", "DX", "MACD", "MACDEXT",
    "MACDFIX", "MFI", "MINUS_DI", "MINUS_DM", "MOM", "PLUS_DI", "PLUS_DM", "PPO", "ROC",
    "ROCP", "ROCR", "ROCR100", "RSI", "STOCH", "STOCHF", "STOCHRSI", "TRIX", "ULTOSC", "WILLR",

    "AD", "ADOSC", "OBV",

    "HT_DCPERIOD", "HT_DCPHASE", "HT_PHASOR", "HT_SINE", "HT_TRENDMODE",

    "AVGPRICE", "MEDPRICE", "TYPPRICE", "WCLPRICE",

    "ATR", "NATR", "TRANGE",

    "CDL2CROWS", "CDL3BLACKCROWS", "CDL3INSIDE", "CDL3LINESTRIKE", "CDL3OUTSIDE",
    "CDL3STARSINSOUTH", "CDL3WHITESOLDIERS", "CDLABANDONEDBABY", "CDLADVANCEBLOCK",
    "CDLBELTHOLD", "CDLBREAKAWAY", "CDLCLOSINGMARUBOZU", "CDLCONCEALBABYSWALL",
    "CDLCOUNTERATTACK", "CDLDARKCLOUDCOVER", "CDLDOJI", "CDLDOJISTAR", "CDLDRAGONFLYDOJI",
    "CDLENGULFING", "CDLEVENINGDOJISTAR", "CDLEVENINGSTAR", "CDLGAPSIDESIDEWHITE",
    "CDLGRAVESTONEDOJI", "CDLHAMMER", "CDLHANGINGMAN", "CDLHARAMI", "CDLHARAMICROSS",
    "CDLHIGHWAVE", "CDLHIKKAKE", "CDLHIKKAKEMOD", "CDLHOMINGPIGEON", "CDLIDENTICAL3CROWS",
    "CDLINNECK", "CDLINVERTEDHAMMER", "CDLKICKING", "CDLKICKINGBYLENGTH", "CDLLADDERBOTTOM",
    "CDLLONGLEGGEDDOJI", "CDLLONGLINE", "CDLMARUBOZU", "CDLMATCHINGLOW", "CDLMATHOLD",
    "CDLMORNINGDOJISTAR", "CDLMORNINGSTAR", "CDLONNECK", "CDLPIERCING", "CDLRICKSHAWMAN",
    "CDLRISEFALL3METHODS", "CDLSEPARATINGLINES", "CDLSHOOTINGSTAR", "CDLSHORTLINE",
    "CDLSPINNINGTOP", "CDLSTALLEDPATTERN", "CDLSTICKSANDWICH", "CDLTAKURI", "CDLTASUKIGAP",
    "CDLTHRUSTING", "CDLTRISTAR", "CDLUNIQUE3RIVER", "CDLUPSIDEGAP2CROWS",
    "CDLXSIDEGAP3METHODS",

    "BETA", "CORREL", "LINEARREG", "LINEARREG_ANGLE", "LINEARREG_INTERCEPT",
    "LINEARREG_SLOPE", "STDDEV", "TSF", "VAR"
]



for i in range(131):
    test_func(ta_functions[i], i)
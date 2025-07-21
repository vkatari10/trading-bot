package technicals

// This file contains TA-Lib wrapper methods 

// MODIFY FLAGS BELOW AND PATHS IF YOU ENCOUNTER 
// LINKER ERRORS BELOW

// DO NOT MODIFY THE INCLUDE STATEMENTS

// DO NOT REMOVE THE -I or -L BEFORE THE PATHS

// DO NOT MODIFY THE -lta-lib, -lm ARGUMENTS

// #cgo CFLAGS: -I/usr/include/ta-lib/
// #cgo LDFLAGS: -L/usr/lib -lta-lib -lm
// #include <ta_libc.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
	"fmt"
)

// UNSUPPORTED METHODS
// MAVP

//==============OVERLAP STUDIES================

// TA_SMA is a wrapper for the SMA function in TA-Lib
// single return value only inside index 0
func TA_SMA(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_SMA: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_SMA: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_SMA(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_SMA: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_SMA: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_SMA
 
// TA_EMA is a wrapper for the EMA function in TA-Lib
// single return value only inside index 0
func TA_EMA(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_EMA: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_EMA: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_EMA(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_EMA: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_EMA: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_EMA
 
// TA_DEMA is a wrapper for the DEMA function in TA-Lib
// single return value only inside index 0
func TA_DEMA(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_DEMA: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_DEMA: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_DEMA(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_DEMA: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_DEMA: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_DEMA
 
// TA_KAMA is a wrapper for the KAMA function in TA-Lib
// single return value only inside index 0
func TA_KAMA(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_KAMA: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_KAMA: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_KAMA(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_KAMA: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_KAMA: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_KAMA
 
// TA_MIDPOINT is a wrapper for the MIDPOINT function in TA-Lib
// single return value only inside index 0
func TA_MIDPOINT(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MIDPOINT: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_MIDPOINT: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_MIDPOINT(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MIDPOINT: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MIDPOINT: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_MIDPOINT
 
// TA_TEMA is a wrapper for the TEMA function in TA-Lib
// single return value only inside index 0
func TA_TEMA(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TEMA: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_TEMA: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_TEMA(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TEMA: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_TEMA: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_TEMA
 
// TA_TRIMA is a wrapper for the TRIMA function in TA-Lib
// single return value only inside index 0
func TA_TRIMA(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TRIMA: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_TRIMA: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_TRIMA(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TRIMA: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_TRIMA: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_TRIMA
 
// TA_WMA is a wrapper for the WMA function in TA-Lib
// single return value only inside index 0
func TA_WMA(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_WMA: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_WMA: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_WMA(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_WMA: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_WMA: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_WMA

// TA_HT_TRENDLINE is a wrapper for the HT_TRENDLINE function in TA-Lib
// single return value only inside index 0
func TA_HT_TRENDLINE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_TRENDLINE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_HT_TRENDLINE(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_TRENDLINE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_HT_TRENDLINE: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_HT_TRENDLINE

// TA_MA is a wrapper for the MA function in TA-Lib
// single return value only inside index 0
func TA_MA(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MA: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_MA: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]
    maType := f.Args["matype"]

    timePeriod := C.int(period)
    maTypeArg := C.TA_MAType(maType) // enum conversion

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_MA(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        maTypeArg,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MA: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MA: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_MA

// TA_MIDPRICE is a wrapper for the MIDPRICE function in TA-Lib
// single return value only inside index 0
func TA_MIDPRICE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MIDPRICE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_MIDPRICE: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_MIDPRICE(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MIDPRICE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MIDPRICE: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_MIDPRICE

// TA_SAR is a wrapper for the SAR function in TA-Lib
// single return value only inside index 0
func TA_SAR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_SAR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_SAR: Provided objects is not a technicals.FeatureTechnical")
    }

    acceleration := f.Args["acceleration"]
    max := f.Args["maximum"]

    accelerationArg := C.double(acceleration)
    maxArg := C.double(max)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_SAR(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        accelerationArg,
        maxArg,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_SAR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_SAR: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_SAR

// TA_SAREXT is a wrapper for the SAREXT function in TA-Lib
// single return value only inside index 0
func TA_SAREXT(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_SAREXT: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_SAREXT: Provided objects is not a technicals.FeatureTechnical")
    }

    start := f.Args["startvalue"]
    offsetOnReverse := f.Args["offsetonreverse"]
    accelerationInitLong := f.Args["accelerationinitlong"]
    accelerationLong := f.Args["accelerationlong"]
    accelerationMaxLong := f.Args["accelerationmaxlong"]
    accelrationInitShort := f.Args["accelerationinitshort"]
    accelerationShort := f.Args["accelerationshort"]
    accelerationMaxShort := f.Args["accelerationmaxshort"]

    startArg := C.double(start)
    offsetOnReverseArg := C.double(offsetOnReverse)
    accelerationInitLongArg := C.double(accelerationInitLong)
    accelerationLongArg := C.double(accelerationLong)
    accelerationMaxLongArg := C.double(accelerationMaxLong)
    accelrationInitShortArg := C.double(accelrationInitShort)
    accelerationShortArg := C.double(accelerationShort)
    accelerationMaxShortArg := C.double(accelerationMaxShort)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_SAREXT(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        startArg,
        offsetOnReverseArg,
        accelerationInitLongArg,
        accelerationLongArg,
        accelerationMaxLongArg,
        accelrationInitShortArg,
        accelerationShortArg,
        accelerationMaxShortArg,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_SAREXT: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_SAREXT: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_SAREXT


 
// TA_SMA is a wrapper for the SMA function in TA-Lib
// single return value only inside index 0
func TA_T3(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_T3: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_T3: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    vfactor := f.Args["vfactor"]

    // real = T3(real, timeperiod=5, vfactor=0)

    timePeriod := C.int(period)
    vArg := C.double(vfactor)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_T3(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        vArg,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_T3: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_T3: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_T3

// TA_MAMA is a wrapper for the MAMA function in TA-Lib
// returns [mama, fama]
func TA_MAMA(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MAMA: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_MAMA: Provided objects is not a technicals.FeatureTechnical")
    }

    fast := f.Args["fastlimit"]
    slow := f.Args["slowlimit"]

    fastArg := C.double(fast)
    slowArg := C.double(slow)

    mama := make([]float64, len(tw.Close))
    fama := make([]float64, len(tw.Close))

    mamaPtr := (*C.double)(unsafe.Pointer(&mama[0]))
    famaPtr := (*C.double)(unsafe.Pointer(&fama[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_MAMA(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        fastArg,
        slowArg,
        &outBegIdx,
        &outNBElement,
        mamaPtr,
        famaPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MAMA: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MAMA: TA-Lib Wrote no results")
    } else {
        return []float64{mama[outNBElement - 1], fama[outNBElement - 1]}, nil
    } 
} // TA_MAMA

 
// TA_BBANDS is a wrapper for the BBANDS function in TA-Lib
// returns [upperband, middleband, lowerband]
func TA_BBANDS(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_BBANDS: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_BBANDS: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]
    updevs := f.Args["nbdevup"]
    downdevs := f.Args["nbdevdn"]
    matype := f.Args["matype"]

    timePeriod := C.int(period)
    updevsArg := C.double(updevs)
    downdevsArg := C.double(downdevs)
    matypeArg := C.TA_MAType(matype)

    upper := make([]float64, len(tw.Close))
    middle := make([]float64, len(tw.Close))
    lower := make([]float64, len(tw.Close))

    upperPtr := (*C.double)(unsafe.Pointer(&upper[0]))
    middlePtr := (*C.double)(unsafe.Pointer(&middle[0]))
    lowerPtr := (*C.double)(unsafe.Pointer(&lower[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_BBANDS(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        updevsArg,
        downdevsArg,
        matypeArg,
        &outBegIdx,
        &outNBElement,
        upperPtr,
        middlePtr,
        lowerPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_BBANDS: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_BBANDS: TA-Lib Wrote no results")
    } else {
        return []float64{
            upper[outNBElement - 1],
            middle[outNBElement - 1],
            lower[outNBElement - 1],
            }, nil
    } 
} // TA_BBANDS

//==============MOMENTUM INDICATORS================
 
// TA_CMO is a wrapper for the CMO function in TA-Lib
// single return value only inside index 0
func TA_CMO(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CMO: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_CMO: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CMO(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CMO: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CMO: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_CMO
 
// TA_MOM is a wrapper for the MOM function in TA-Lib
// single return value only inside index 0
func TA_MOM(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MOM: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_MOM: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_MOM(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MOM: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MOM: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_MOM
 
// TA_ROC is a wrapper for the ROC function in TA-Lib
// single return value only inside index 0
func TA_ROC(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ROC: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_ROC: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_ROC(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ROC: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_ROC: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_ROC
 
// TA_ROCP is a wrapper for the ROCP function in TA-Lib
// single return value only inside index 0
func TA_ROCP(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ROCP: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_ROCP: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_ROCP(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ROCP: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_ROCP: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_ROCP
 
// TA_ROCR is a wrapper for the ROCR function in TA-Lib
// single return value only inside index 0
func TA_ROCR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ROCR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_ROCR: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_ROCR(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ROCR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_ROCR: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_ROCR
 
// TA_ROCR100 is a wrapper for the ROCR100 function in TA-Lib
// single return value only inside index 0
func TA_ROCR100(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ROCR100: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_ROCR100: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_ROCR100(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ROCR100: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_ROCR100: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_ROCR100
 
// TA_RSI is a wrapper for the RSI function in TA-Lib
// single return value only inside index 0
func TA_RSI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_RSI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_RSI: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_RSI(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_RSI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_RSI: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_RSI
 
// TA_TRIX is a wrapper for the TRIX function in TA-Lib
// single return value only inside index 0
func TA_TRIX(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TRIX: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_TRIX: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_TRIX(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TRIX: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_TRIX: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_TRIX

// TA_ADX is a wrapper for the ADX function in TA-Lib
// single return value only inside index 0
func TA_ADX(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ADX: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_ADX: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_ADX(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ADX: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_ADX: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_ADX
 
// TA_ADXR is a wrapper for the ADXR function in TA-Lib
// single return value only inside index 0
func TA_ADXR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ADXR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_ADXR: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_ADXR(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ADXR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_ADXR: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_ADXR
 
// TA_CCI is a wrapper for the CCI function in TA-Lib
// single return value only inside index 0
func TA_CCI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CCI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_CCI: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CCI(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CCI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CCI: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_CCI
 
// TA_DX is a wrapper for the DX function in TA-Lib
// single return value only inside index 0
func TA_DX(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_DX: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_DX: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_DX(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_DX: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_DX: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_DX
 
// TA_MINUS_DI is a wrapper for the MINUS_DI function in TA-Lib
// single return value only inside index 0
func TA_MINUS_DI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MINUS_DI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_MINUS_DI: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_MINUS_DI(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MINUS_DI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MINUS_DI: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_MINUS_DI
 
// TA_PLUS_DI is a wrapper for the PLUS_DI function in TA-Lib
// single return value only inside index 0
func TA_PLUS_DI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_PLUS_DI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_PLUS_DI: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_PLUS_DI(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_PLUS_DI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_PLUS_DI: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_PLUS_DI
 
// TA_WILLR is a wrapper for the WILLR function in TA-Lib
// single return value only inside index 0
func TA_WILLR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_WILLR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_WILLR: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_WILLR(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_WILLR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_WILLR: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_WILLR

 
// TA_AROON is a wrapper for the AROON function in TA-Lib
// return [aroonup, aroondown]
func TA_AROON(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_AROON: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_AROON: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    aroonup := make([]float64, len(tw.Close))
    aroondown := make([]float64, len(tw.Close))

    upPtr := (*C.double)(unsafe.Pointer(&aroonup[0]))
    downPtr := (*C.double)(unsafe.Pointer(&aroondown[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_AROON(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        upPtr,
        downPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_AROON: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_AROON: TA-Lib Wrote no results")
    } else {
        return []float64{
            aroonup[outNBElement - 1],
            aroondown[outNBElement - 1],
            }, nil
    } 
} // TA_AROON
 
// TA_AROONOSC is a wrapper for the AROONOSC function in TA-Lib
// single return value only inside index 0
func TA_AROONOSC(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_AROONOSC: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_AROONOSC: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_AROONOSC(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_AROONOSC: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_AROONOSC: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_AROONOSC
 
// TA_MINUS_DM is a wrapper for the MINUS_DM function in TA-Lib
// single return value only inside index 0
func TA_MINUS_DM(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MINUS_DM: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_MINUS_DM: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_MINUS_DM(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MINUS_DM: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MINUS_DM: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_MINUS_DM
 
// TA_PLUS_DM is a wrapper for the PLUS_DM function in TA-Lib
// single return value only inside index 0
func TA_PLUS_DM(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_PLUS_DM: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_PLUS_DM: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_PLUS_DM(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_PLUS_DM: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_PLUS_DM: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_PLUS_DM

 
// TA_MFI is a wrapper for the MFI function in TA-Lib
// single return value only inside index 0
func TA_MFI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MFI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_MFI: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_MFI(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        tw.VolumePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MFI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MFI: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_MFI

// TA_ULTOSC is a wrapper for the ULTOSC function in TA-Lib
// single return value only inside index 0
func TA_ULTOSC(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ULTOSC: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_ULTOSC: Provided objects is not a technicals.FeatureTechnical")
    }

    period1 := f.Args["timeperiod1"]
    period2 := f.Args["timeperiod2"]
    period3 := f.Args["timeperiod3"]

    timePeriod1 := C.int(period1)
    timePeriod2 := C.int(period2)
    timePeriod3 := C.int(period3)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_ULTOSC(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        timePeriod1,
        timePeriod2,
        timePeriod3,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ULTOSC: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_ULTOSC: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_ULTOSC

// TA_APO is a wrapper for the APO function in TA-Lib
// single return value only inside index 0
func TA_APO(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_APO: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_APO: Provided objects is not a technicals.FeatureTechnical")
    }

    fastPeriod := f.Args["fastperiod"]
    slowPeriod := f.Args["slowperiod"]
    maType := f.Args["matype"]

    fastPeriodArg := C.int(fastPeriod)
    slowPeriodArg := C.int(slowPeriod)
    maTypeArg := C.TA_MAType(maType)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_APO(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        fastPeriodArg,
        slowPeriodArg,
        maTypeArg,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_APO: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_APO: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_APO

// TA_BOP is a wrapper for the BOP function in TA-Lib
// single return value only inside index 0
func TA_BOP(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_BOP: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_BOP(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_BOP: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_BOP: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_BOP

// TA_MACDFIX is a wrapper for the MACDFIX function in TA-Lib
// return [macd, macd signal, macd hist]
func TA_MACDFIX(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MACDFIX: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_MACDFIX: Provided objects is not a technicals.FeatureTechnical")
    }

    signalPeriod := f.Args["signalperiod"]
    signalPeriodArg := C.int(signalPeriod)

    macd := make([]float64, len(tw.Close))
    macdSignal := make([]float64, len(tw.Close))
    macdHist := make([]float64, len(tw.Close))

    macdPtr := (*C.double)(unsafe.Pointer(&macd[0]))
    macdSigPtr := (*C.double)(unsafe.Pointer(&macdSignal[0]))
    macdHistPtr := (*C.double)(unsafe.Pointer(&macdHist[0]))

    var outBegIdx, outNBElement C.int
    //macd, macdsignal, macdhist = MACD(real, fastperiod=12, slowperiod=26, signalperiod=9)
    retCode := C.TA_MACDFIX(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        signalPeriodArg,
        &outBegIdx,
        &outNBElement,
        macdPtr,
        macdSigPtr,
        macdHistPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MACDFIX: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MACDFIX: TA-Lib Wrote no results")
    } else {
        return []float64{
            macd[outNBElement - 1],
            macdSignal[outNBElement - 1],
            macdHist[outNBElement - 1],
            }, nil
    } 
} // TA_MACDFIX

// TA_MACD is a wrapper for the MACD function in TA-Lib
// return [macd, macd signal, macd hist]
func TA_MACD(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MACD: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_MACD: Provided objects is not a technicals.FeatureTechnical")
    }

    fastPeriod := f.Args["fastperiod"]
    slowPeriod := f.Args["slowperiod"]
    signalPeriod := f.Args["signalperiod"]

    fastPeriodArg := C.int(fastPeriod)
    slowPeriodArg := C.int(slowPeriod)
    signalPeriodArg := C.int(signalPeriod)

    macd := make([]float64, len(tw.Close))
    macdSignal := make([]float64, len(tw.Close))
    macdHist := make([]float64, len(tw.Close))

    macdPtr := (*C.double)(unsafe.Pointer(&macd[0]))
    macdSigPtr := (*C.double)(unsafe.Pointer(&macdSignal[0]))
    macdHistPtr := (*C.double)(unsafe.Pointer(&macdHist[0]))

    var outBegIdx, outNBElement C.int
    //macd, macdsignal, macdhist = MACD(real, fastperiod=12, slowperiod=26, signalperiod=9)
    retCode := C.TA_MACD(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        fastPeriodArg,
        slowPeriodArg,
        signalPeriodArg,
        &outBegIdx,
        &outNBElement,
        macdPtr,
        macdSigPtr,
        macdHistPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MACD: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MACD: TA-Lib Wrote no results")
    } else {
        return []float64{
            macd[outNBElement - 1],
            macdSignal[outNBElement - 1],
            macdHist[outNBElement - 1],
            }, nil
    } 
} // TA_MACD

// TA_MACDEXT is a wrapper for the MACDEXT function in TA-Lib
// return [macd, macd signal, macd hist]
func TA_MACDEXT(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MACDEXT: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_MACDEXT: Provided objects is not a technicals.FeatureTechnical")
    }

    fastPeriod := f.Args["fastperiod"]
    slowPeriod := f.Args["slowperiod"]
    signalPeriod := f.Args["signalperiod"]

    fastMA := f.Args["fastmatype"]
    slowMA := f.Args["slowmatype"]
    signalMA := f.Args["signalmatype"]

    fastMAArg := C.TA_MAType(fastMA)
    slowMAArg := C.TA_MAType(slowMA)
    signalMAArg := C.TA_MAType(signalMA)

    fastPeriodArg := C.int(fastPeriod)
    slowPeriodArg := C.int(slowPeriod)
    signalPeriodArg := C.int(signalPeriod)

    macd := make([]float64, len(tw.Close))
    macdSignal := make([]float64, len(tw.Close))
    macdHist := make([]float64, len(tw.Close))

    macdPtr := (*C.double)(unsafe.Pointer(&macd[0]))
    macdSigPtr := (*C.double)(unsafe.Pointer(&macdSignal[0]))
    macdHistPtr := (*C.double)(unsafe.Pointer(&macdHist[0]))

    var outBegIdx, outNBElement C.int
    //macd, macdsignal, macdhist = MACD(real, fastperiod=12, slowperiod=26, signalperiod=9)
    retCode := C.TA_MACDEXT(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        fastPeriodArg,
        fastMAArg,
        slowPeriodArg,
        slowMAArg,
        signalPeriodArg,
        signalMAArg,
        &outBegIdx,
        &outNBElement,
        macdPtr,
        macdSigPtr,
        macdHistPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MACDEXT: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MACDEXT: TA-Lib Wrote no results")
    } else {
        return []float64{
            macd[outNBElement - 1],
            macdSignal[outNBElement - 1],
            macdHist[outNBElement - 1],
            }, nil
    } 
} // TA_MACDEXT

// TA_PPO is a wrapper for the PPO function in TA-Lib
// return single value in index 0
func TA_PPO(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_PPO: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_PPO: Provided objects is not a technicals.FeatureTechnical")
    }

    fastPeriod := f.Args["fastperiod"]
    slowPeriod := f.Args["slowperiod"]
    MAType := f.Args["matype"]

    fastPeriodArg := C.int(fastPeriod)
    slowPeriodArg := C.int(slowPeriod)
    MATypeArg := C.TA_MAType(MAType)

    output := make([]float64, len(tw.Open))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int
    retCode := C.TA_PPO(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        fastPeriodArg,
        slowPeriodArg,
        MATypeArg,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_PPO: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_PPO: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_PPO

// TA_STOCH is a wrapper for the STOCH function in TA-Lib
// return [slowk, slowd]
func TA_STOCH(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_STOCH: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_STOCH: Provided objects is not a technicals.FeatureTechnical")
    }

    fastkPeriod := f.Args["fastk_period"]
    slowkPeriod := f.Args["slowk_period"]
    slowkMA := f.Args["slowk_matype"]
    slowdPeriod := f.Args["slowd_period"]
    slowdMA := f.Args["slowd_matype"]

    fastkPeriodArg := C.int(fastkPeriod)
    slowkPeriodArg := C.int(slowkPeriod)
    slowkMAArg := C.TA_MAType(slowkMA)
    slowdPeriodArg := C.int(slowdPeriod)
    slowdMAArg := C.TA_MAType(slowdMA)


    slowk := make([]float64, len(tw.Close))
    slowd := make([]float64, len(tw.Close))
    slowkPtr := (*C.double)(unsafe.Pointer(&slowk[0]))
    slowdPtr := (*C.double)(unsafe.Pointer(&slowd[0]))

    var outBegIdx, outNBElement C.int
    //macd, macdsignal, macdhist = MACD(real, fastperiod=12, slowperiod=26, signalperiod=9)
    retCode := C.TA_STOCH(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        fastkPeriodArg,
        slowkPeriodArg,
        slowkMAArg,
        slowdPeriodArg,
        slowdMAArg,
        &outBegIdx,
        &outNBElement,
        slowkPtr,
        slowdPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_STOCH: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_STOCH: TA-Lib Wrote no results")
    } else {
        return []float64{
            slowk[outNBElement - 1],
            slowd[outNBElement - 1],
            }, nil
    } 
} // TA_STOCH

// TA_STOCHF is a wrapper for the STOCHF function in TA-Lib
// return [fastk, fastd]
func TA_STOCHF(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_STOCHF: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_STOCHF: Provided objects is not a technicals.FeatureTechnical")
    }

    fastkPeriod := f.Args["fastk_period"]
    fastdPeriod := f.Args["fastd_period"]
    fastdMA := f.Args["fastd_matype"]

    fastkPeriodArg := C.int(fastkPeriod)
    fastdPeriodArg := C.int(fastdPeriod)
    fastdMAArg := C.TA_MAType(fastdMA)


    slowk := make([]float64, len(tw.Close))
    slowd := make([]float64, len(tw.Close))
    slowkPtr := (*C.double)(unsafe.Pointer(&slowk[0]))
    slowdPtr := (*C.double)(unsafe.Pointer(&slowd[0]))

    var outBegIdx, outNBElement C.int
    retCode := C.TA_STOCHF(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        fastkPeriodArg,
        fastdPeriodArg,
        fastdMAArg,
        &outBegIdx,
        &outNBElement,
        slowkPtr,
        slowdPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_STOCHF: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_STOCHF: TA-Lib Wrote no results")
    } else {
        return []float64{
            slowk[outNBElement - 1],
            slowd[outNBElement - 1],
            }, nil
    } 
} // TA_STOCHF

// TA_STOCHRSI is a wrapper for the STOCHRSI function in TA-Lib
// return [fastk, fastd]
func TA_STOCRSI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_STOCHRSI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_STOCHRSI: Provided objects is not a technicals.FeatureTechnical")
    }

    timeperiod := f.Args["timeperiod"]
    fastkPeriod := f.Args["fastk_period"]
    fastdPeriod := f.Args["fastd_period"]
    fastdMA := f.Args["fastd_matype"]

    fastkPeriodArg := C.int(fastkPeriod)
    fastdPeriodArg := C.int(fastdPeriod)
    fastdMAArg := C.TA_MAType(fastdMA)
    timeperiodArg := C.int(timeperiod)

    //fastk, fastd = STOCHRSI(real, timeperiod=14, fastk_period=5, fastd_period=3, fastd_matype=0)

    slowk := make([]float64, len(tw.Close))
    slowd := make([]float64, len(tw.Close))
    slowkPtr := (*C.double)(unsafe.Pointer(&slowk[0]))
    slowdPtr := (*C.double)(unsafe.Pointer(&slowd[0]))

    var outBegIdx, outNBElement C.int
    retCode := C.TA_STOCHRSI(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timeperiodArg,
        fastkPeriodArg,
        fastdPeriodArg,
        fastdMAArg, 
        &outBegIdx,
        &outNBElement,
        slowkPtr,
        slowdPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_STOCHRSI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_STOCHRSI: TA-Lib Wrote no results")
    } else {
        return []float64{
            slowk[outNBElement - 1],
            slowd[outNBElement - 1],
            }, nil
    } 
} // TA_STOCHRSI

//===============VOLUME INDICATORS==============

// TA_AD is a wrapper for the AD function in TA-Lib
// single return value only inside index 0
func TA_AD(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_AD: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_AD(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        tw.VolumePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_AD: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_AD: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_AD

// TA_OBV is a wrapper for the OBV function in TA-Lib
// single return value only inside index 0
func TA_OBV(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_OBV: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_OBV(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        tw.VolumePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_OBV: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_OBV: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_OBV

// TA_ADOSC is a wrapper for the ADOSC function in TA-Lib
// single return value only inside index 0
func TA_ADOSC(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ADOSC: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_STOCHRSI: Provided objects is not a technicals.FeatureTechnical")
    }

    fastPeriod := f.Args["fastperiod"]
    slowPeriod := f.Args["slowperiod"]

    fastPeriodArgs := C.int(fastPeriod)
    slowPeriodArgs := C.int(slowPeriod)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_ADOSC(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        tw.VolumePtr,
        fastPeriodArgs,
        slowPeriodArgs,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ADOSC: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_ADOSC: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_ADOSC

 
// TA_HT_DCPERIOD is a wrapper for the HT_DCPERIOD function in TA-Lib
// single return value only inside index 0
func TA_HT_DCPERIOD(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_DCPERIOD: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_HT_DCPERIOD(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_DCPERIOD: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_HT_DCPERIOD: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_HT_DCPERIOD
 
// TA_HT_DCPHASE is a wrapper for the HT_DCPHASE function in TA-Lib
// single return value only inside index 0
func TA_HT_DCPHASE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_DCPHASE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_HT_DCPHASE(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_DCPHASE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_HT_DCPHASE: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_HT_DCPHASE
 
// TA_HT_TRENDMODE is a wrapper for the HT_TRENDMODE function in TA-Lib
// single return value only inside index 0
func TA_HT_TRENDMODE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_TRENDMODE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_HT_TRENDMODE(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_TRENDMODE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_HT_TRENDMODE: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_HT_TRENDMODE

 
// TA_HT_PHASOR is a wrapper for the HT_PHASOR function in TA-Lib
// return [inphase, quadrature]
func TA_HT_PHASOR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_PHASOR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    output2 := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))
    out2Ptr := (*C.double)(unsafe.Pointer(&output2[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_HT_PHASOR(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
        out2Ptr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_PHASOR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_HT_PHASOR: TA-Lib Wrote no results")
    } else {
        return []float64{
            output[outNBElement - 1],
            output2[outNBElement - 1],
            }, nil
    } 
} // TA_HT_PHASOR
 
// TA_HT_SINE is a wrapper for the HT_SINE function in TA-Lib
// return [sine, leadsine]
func TA_HT_SINE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_SINE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    output2 := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))
    out2Ptr := (*C.double)(unsafe.Pointer(&output2[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_HT_SINE(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
        out2Ptr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_HT_SINE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_HT_SINE: TA-Lib Wrote no results")
    } else {
        return []float64{
            output[outNBElement - 1],
            output2[outNBElement - 1],
            }, nil
    } 
} // TA_HT_SINE

//============PRICE TRANSFORM STUDIES=============
 
// TA_AVGPRICE is a wrapper for the AVGPRICE function in TA-Lib
// single return value only inside index 0
func TA_AVGPRICE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_AVGPRICE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_AVGPRICE(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_AVGPRICE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_AVGPRICE: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_AVGPRICE
 
// TA_MEDPRICE is a wrapper for the MEDPRICE function in TA-Lib
// single return value only inside index 0
func TA_MEDPRICE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MEDPRICE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_MEDPRICE(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_MEDPRICE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_MEDPRICE: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_MEDPRICE
 
// TA_TYPPRICE is a wrapper for the TYPPRICE function in TA-Lib
// single return value only inside index 0
func TA_TYPPRICE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TYPPRICE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_TYPPRICE(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TYPPRICE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_TYPPRICE: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_TYPPRICE
 
// TA_WCLPRICE is a wrapper for the WCLPRICE function in TA-Lib
// single return value only inside index 0
func TA_WCLPRICE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_WCLPRICE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_WCLPRICE(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_WCLPRICE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_WCLPRICE: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_WCLPRICE

//================VOLATILITY INDICATORS================

// TA_TRANGE is a wrapper for the TRANGE function in TA-Lib
// single return value only inside index 0
func TA_TRANGE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TRANGE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_TRANGE(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.HighPtr,
        tw.LowPtr,
        tw.ClosePtr,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TRANGE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_TRANGE: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_TRANGE

// TA_ATR is a wrapper for the ATR function in TA-Lib
// single return value only inside index 0
func TA_ATR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ATR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_ATR: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_ATR(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_ATR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_ATR: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_ATR
 
// TA_NATR is a wrapper for the NATR function in TA-Lib
// single return value only inside index 0
func TA_NATR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_NATR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_NATR: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_NATR(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_NATR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_NATR: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_NATR

//==============PATTERN RECOGNITION===============

// TA_CDL2CROWS is a wrapper for the CDL2CROWS function in TA-Lib
// single return value only inside index 0
func TA_CDL2CROWS(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL2CROWS: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDL2CROWS(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL2CROWS: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDL2CROWS: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDL2CROWS
 
// TA_CDL3BLACKCROWS is a wrapper for the CDL3BLACKCROWS function in TA-Lib
// single return value only inside index 0
func TA_CDL3BLACKCROWS(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3BLACKCROWS: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDL3BLACKCROWS(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3BLACKCROWS: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDL3BLACKCROWS: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDL3BLACKCROWS
 
// TA_CDL3INSIDE is a wrapper for the CDL3INSIDE function in TA-Lib
// single return value only inside index 0
func TA_CDL3INSIDE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3INSIDE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDL3INSIDE(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3INSIDE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDL3INSIDE: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDL3INSIDE
 
// TA_CDL3LINESTRIKE is a wrapper for the CDL3LINESTRIKE function in TA-Lib
// single return value only inside index 0
func TA_CDL3LINESTRIKE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3LINESTRIKE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDL3LINESTRIKE(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3LINESTRIKE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDL3LINESTRIKE: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDL3LINESTRIKE
 
// TA_CDL3OUTSIDE is a wrapper for the CDL3OUTSIDE function in TA-Lib
// single return value only inside index 0
func TA_CDL3OUTSIDE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3OUTSIDE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDL3OUTSIDE(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3OUTSIDE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDL3OUTSIDE: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDL3OUTSIDE
 
// TA_CDL3STARSINSOUTH is a wrapper for the CDL3STARSINSOUTH function in TA-Lib
// single return value only inside index 0
func TA_CDL3STARSINSOUTH(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3STARSINSOUTH: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDL3STARSINSOUTH(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3STARSINSOUTH: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDL3STARSINSOUTH: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDL3STARSINSOUTH
 
// TA_CDL3WHITESOLDIERS is a wrapper for the CDL3WHITESOLDIERS function in TA-Lib
// single return value only inside index 0
func TA_CDL3WHITESOLDIERS(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3WHITESOLDIERS: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDL3WHITESOLDIERS(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDL3WHITESOLDIERS: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDL3WHITESOLDIERS: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDL3WHITESOLDIERS
 
// TA_CDLADVANCEBLOCK is a wrapper for the CDLADVANCEBLOCK function in TA-Lib
// single return value only inside index 0
func TA_CDLADVANCEBLOCK(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLADVANCEBLOCK: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLADVANCEBLOCK(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLADVANCEBLOCK: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLADVANCEBLOCK: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLADVANCEBLOCK
 
// TA_CDLBELTHOLD is a wrapper for the CDLBELTHOLD function in TA-Lib
// single return value only inside index 0
func TA_CDLBELTHOLD(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLBELTHOLD: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLBELTHOLD(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLBELTHOLD: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLBELTHOLD: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLBELTHOLD
 
// TA_CDLBREAKAWAY is a wrapper for the CDLBREAKAWAY function in TA-Lib
// single return value only inside index 0
func TA_CDLBREAKAWAY(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLBREAKAWAY: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLBREAKAWAY(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLBREAKAWAY: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLBREAKAWAY: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLBREAKAWAY
 
// TA_CDLCLOSINGMARUBOZU is a wrapper for the CDLCLOSINGMARUBOZU function in TA-Lib
// single return value only inside index 0
func TA_CDLCLOSINGMARUBOZU(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLCLOSINGMARUBOZU: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLCLOSINGMARUBOZU(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLCLOSINGMARUBOZU: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLCLOSINGMARUBOZU: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLCLOSINGMARUBOZU
 
// TA_CDLCONCEALBABYSWALL is a wrapper for the CDLCONCEALBABYSWALL function in TA-Lib
// single return value only inside index 0
func TA_CDLCONCEALBABYSWALL(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLCONCEALBABYSWALL: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLCONCEALBABYSWALL(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLCONCEALBABYSWALL: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLCONCEALBABYSWALL: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLCONCEALBABYSWALL
 
// TA_CDLCOUNTERATTACK is a wrapper for the CDLCOUNTERATTACK function in TA-Lib
// single return value only inside index 0
func TA_CDLCOUNTERATTACK(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLCOUNTERATTACK: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLCOUNTERATTACK(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLCOUNTERATTACK: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLCOUNTERATTACK: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLCOUNTERATTACK
 
// TA_CDLDOJI is a wrapper for the CDLDOJI function in TA-Lib
// single return value only inside index 0
func TA_CDLDOJI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLDOJI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLDOJI(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLDOJI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLDOJI: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLDOJI
 
// TA_CDLDOJISTAR is a wrapper for the CDLDOJISTAR function in TA-Lib
// single return value only inside index 0
func TA_CDLDOJISTAR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLDOJISTAR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLDOJISTAR(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLDOJISTAR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLDOJISTAR: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLDOJISTAR
 
// TA_CDLDRAGONFLYDOJI is a wrapper for the CDLDRAGONFLYDOJI function in TA-Lib
// single return value only inside index 0
func TA_CDLDRAGONFLYDOJI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLDRAGONFLYDOJI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLDRAGONFLYDOJI(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLDRAGONFLYDOJI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLDRAGONFLYDOJI: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLDRAGONFLYDOJI
 
// TA_CDLENGULFING is a wrapper for the CDLENGULFING function in TA-Lib
// single return value only inside index 0
func TA_CDLENGULFING(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLENGULFING: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLENGULFING(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLENGULFING: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLENGULFING: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLENGULFING
 
// TA_CDLGAPSIDESIDEWHITE is a wrapper for the CDLGAPSIDESIDEWHITE function in TA-Lib
// single return value only inside index 0
func TA_CDLGAPSIDESIDEWHITE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLGAPSIDESIDEWHITE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLGAPSIDESIDEWHITE(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLGAPSIDESIDEWHITE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLGAPSIDESIDEWHITE: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLGAPSIDESIDEWHITE
 
// TA_CDLGRAVESTONEDOJI is a wrapper for the CDLGRAVESTONEDOJI function in TA-Lib
// single return value only inside index 0
func TA_CDLGRAVESTONEDOJI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLGRAVESTONEDOJI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLGRAVESTONEDOJI(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLGRAVESTONEDOJI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLGRAVESTONEDOJI: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLGRAVESTONEDOJI
 
// TA_CDLHAMMER is a wrapper for the CDLHAMMER function in TA-Lib
// single return value only inside index 0
func TA_CDLHAMMER(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHAMMER: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLHAMMER(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHAMMER: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLHAMMER: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLHAMMER
 
// TA_CDLHANGINGMAN is a wrapper for the CDLHANGINGMAN function in TA-Lib
// single return value only inside index 0
func TA_CDLHANGINGMAN(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHANGINGMAN: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLHANGINGMAN(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHANGINGMAN: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLHANGINGMAN: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLHANGINGMAN
 
// TA_CDLHARAMI is a wrapper for the CDLHARAMI function in TA-Lib
// single return value only inside index 0
func TA_CDLHARAMI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHARAMI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLHARAMI(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHARAMI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLHARAMI: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLHARAMI
 
// TA_CDLHARAMICROSS is a wrapper for the CDLHARAMICROSS function in TA-Lib
// single return value only inside index 0
func TA_CDLHARAMICROSS(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHARAMICROSS: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLHARAMICROSS(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHARAMICROSS: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLHARAMICROSS: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLHARAMICROSS
 
// TA_CDLHIGHWAVE is a wrapper for the CDLHIGHWAVE function in TA-Lib
// single return value only inside index 0
func TA_CDLHIGHWAVE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHIGHWAVE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLHIGHWAVE(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHIGHWAVE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLHIGHWAVE: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLHIGHWAVE
 
// TA_CDLHIKKAKE is a wrapper for the CDLHIKKAKE function in TA-Lib
// single return value only inside index 0
func TA_CDLHIKKAKE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHIKKAKE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLHIKKAKE(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHIKKAKE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLHIKKAKE: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLHIKKAKE
 
// TA_CDLHIKKAKEMOD is a wrapper for the CDLHIKKAKEMOD function in TA-Lib
// single return value only inside index 0
func TA_CDLHIKKAKEMOD(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHIKKAKEMOD: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLHIKKAKEMOD(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHIKKAKEMOD: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLHIKKAKEMOD: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLHIKKAKEMOD
 
// TA_CDLHOMINGPIGEON is a wrapper for the CDLHOMINGPIGEON function in TA-Lib
// single return value only inside index 0
func TA_CDLHOMINGPIGEON(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHOMINGPIGEON: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLHOMINGPIGEON(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLHOMINGPIGEON: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLHOMINGPIGEON: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLHOMINGPIGEON
 
// TA_CDLIDENTICAL3CROWS is a wrapper for the CDLIDENTICAL3CROWS function in TA-Lib
// single return value only inside index 0
func TA_CDLIDENTICAL3CROWS(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLIDENTICAL3CROWS: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLIDENTICAL3CROWS(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLIDENTICAL3CROWS: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLIDENTICAL3CROWS: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLIDENTICAL3CROWS
 
// TA_CDLINNECK is a wrapper for the CDLINNECK function in TA-Lib
// single return value only inside index 0
func TA_CDLINNECK(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLINNECK: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLINNECK(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLINNECK: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLINNECK: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLINNECK
 
// TA_CDLINVERTEDHAMMER is a wrapper for the CDLINVERTEDHAMMER function in TA-Lib
// single return value only inside index 0
func TA_CDLINVERTEDHAMMER(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLINVERTEDHAMMER: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLINVERTEDHAMMER(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLINVERTEDHAMMER: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLINVERTEDHAMMER: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLINVERTEDHAMMER
 
// TA_CDLKICKING is a wrapper for the CDLKICKING function in TA-Lib
// single return value only inside index 0
func TA_CDLKICKING(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLKICKING: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLKICKING(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLKICKING: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLKICKING: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLKICKING
 
// TA_CDLKICKINGBYLENGTH is a wrapper for the CDLKICKINGBYLENGTH function in TA-Lib
// single return value only inside index 0
func TA_CDLKICKINGBYLENGTH(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLKICKINGBYLENGTH: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLKICKINGBYLENGTH(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLKICKINGBYLENGTH: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLKICKINGBYLENGTH: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLKICKINGBYLENGTH
 
// TA_CDLLADDERBOTTOM is a wrapper for the CDLLADDERBOTTOM function in TA-Lib
// single return value only inside index 0
func TA_CDLLADDERBOTTOM(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLLADDERBOTTOM: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLLADDERBOTTOM(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLLADDERBOTTOM: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLLADDERBOTTOM: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLLADDERBOTTOM
 
// TA_CDLLONGLEGGEDDOJI is a wrapper for the CDLLONGLEGGEDDOJI function in TA-Lib
// single return value only inside index 0
func TA_CDLLONGLEGGEDDOJI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLLONGLEGGEDDOJI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLLONGLEGGEDDOJI(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLLONGLEGGEDDOJI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLLONGLEGGEDDOJI: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLLONGLEGGEDDOJI
 
// TA_CDLLONGLINE is a wrapper for the CDLLONGLINE function in TA-Lib
// single return value only inside index 0
func TA_CDLLONGLINE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLLONGLINE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLLONGLINE(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLLONGLINE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLLONGLINE: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLLONGLINE
 
// TA_CDLMARUBOZU is a wrapper for the CDLMARUBOZU function in TA-Lib
// single return value only inside index 0
func TA_CDLMARUBOZU(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLMARUBOZU: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLMARUBOZU(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLMARUBOZU: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLMARUBOZU: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLMARUBOZU
 
// TA_CDLMATCHINGLOW is a wrapper for the CDLMATCHINGLOW function in TA-Lib
// single return value only inside index 0
func TA_CDLMATCHINGLOW(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLMATCHINGLOW: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLMATCHINGLOW(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLMATCHINGLOW: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLMATCHINGLOW: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLMATCHINGLOW
 
// TA_CDLONNECK is a wrapper for the CDLONNECK function in TA-Lib
// single return value only inside index 0
func TA_CDLONNECK(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLONNECK: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLONNECK(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLONNECK: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLONNECK: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLONNECK
 
// TA_CDLPIERCING is a wrapper for the CDLPIERCING function in TA-Lib
// single return value only inside index 0
func TA_CDLPIERCING(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLPIERCING: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLPIERCING(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLPIERCING: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLPIERCING: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLPIERCING
 
// TA_CDLRICKSHAWMAN is a wrapper for the CDLRICKSHAWMAN function in TA-Lib
// single return value only inside index 0
func TA_CDLRICKSHAWMAN(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLRICKSHAWMAN: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLRICKSHAWMAN(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLRICKSHAWMAN: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLRICKSHAWMAN: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLRICKSHAWMAN
 
// TA_CDLRISEFALL3METHODS is a wrapper for the CDLRISEFALL3METHODS function in TA-Lib
// single return value only inside index 0
func TA_CDLRISEFALL3METHODS(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLRISEFALL3METHODS: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLRISEFALL3METHODS(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLRISEFALL3METHODS: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLRISEFALL3METHODS: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLRISEFALL3METHODS
 
// TA_CDLSEPARATINGLINES is a wrapper for the CDLSEPARATINGLINES function in TA-Lib
// single return value only inside index 0
func TA_CDLSEPARATINGLINES(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSEPARATINGLINES: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLSEPARATINGLINES(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSEPARATINGLINES: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLSEPARATINGLINES: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLSEPARATINGLINES
 
// TA_CDLSHOOTINGSTAR is a wrapper for the CDLSHOOTINGSTAR function in TA-Lib
// single return value only inside index 0
func TA_CDLSHOOTINGSTAR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSHOOTINGSTAR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLSHOOTINGSTAR(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSHOOTINGSTAR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLSHOOTINGSTAR: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLSHOOTINGSTAR
 
// TA_CDLSHORTLINE is a wrapper for the CDLSHORTLINE function in TA-Lib
// single return value only inside index 0
func TA_CDLSHORTLINE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSHORTLINE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLSHORTLINE(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSHORTLINE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLSHORTLINE: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLSHORTLINE
 
// TA_CDLSPINNINGTOP is a wrapper for the CDLSPINNINGTOP function in TA-Lib
// single return value only inside index 0
func TA_CDLSPINNINGTOP(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSPINNINGTOP: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLSPINNINGTOP(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSPINNINGTOP: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLSPINNINGTOP: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLSPINNINGTOP
 
// TA_CDLSTALLEDPATTERN is a wrapper for the CDLSTALLEDPATTERN function in TA-Lib
// single return value only inside index 0
func TA_CDLSTALLEDPATTERN(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSTALLEDPATTERN: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLSTALLEDPATTERN(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSTALLEDPATTERN: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLSTALLEDPATTERN: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLSTALLEDPATTERN
 
// TA_CDLSTICKSANDWICH is a wrapper for the CDLSTICKSANDWICH function in TA-Lib
// single return value only inside index 0
func TA_CDLSTICKSANDWICH(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSTICKSANDWICH: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLSTICKSANDWICH(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLSTICKSANDWICH: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLSTICKSANDWICH: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLSTICKSANDWICH
 
// TA_CDLTAKURI is a wrapper for the CDLTAKURI function in TA-Lib
// single return value only inside index 0
func TA_CDLTAKURI(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLTAKURI: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLTAKURI(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLTAKURI: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLTAKURI: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLTAKURI
 
// TA_CDLTASUKIGAP is a wrapper for the CDLTASUKIGAP function in TA-Lib
// single return value only inside index 0
func TA_CDLTASUKIGAP(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLTASUKIGAP: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLTASUKIGAP(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLTASUKIGAP: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLTASUKIGAP: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLTASUKIGAP
 
// TA_CDLTHRUSTING is a wrapper for the CDLTHRUSTING function in TA-Lib
// single return value only inside index 0
func TA_CDLTHRUSTING(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLTHRUSTING: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLTHRUSTING(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLTHRUSTING: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLTHRUSTING: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLTHRUSTING
 
// TA_CDLTRISTAR is a wrapper for the CDLTRISTAR function in TA-Lib
// single return value only inside index 0
func TA_CDLTRISTAR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLTRISTAR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLTRISTAR(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLTRISTAR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLTRISTAR: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLTRISTAR
 
// TA_CDLUNIQUE3RIVER is a wrapper for the CDLUNIQUE3RIVER function in TA-Lib
// single return value only inside index 0
func TA_CDLUNIQUE3RIVER(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLUNIQUE3RIVER: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLUNIQUE3RIVER(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLUNIQUE3RIVER: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLUNIQUE3RIVER: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLUNIQUE3RIVER
 
// TA_CDLUPSIDEGAP2CROWS is a wrapper for the CDLUPSIDEGAP2CROWS function in TA-Lib
// single return value only inside index 0
func TA_CDLUPSIDEGAP2CROWS(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLUPSIDEGAP2CROWS: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLUPSIDEGAP2CROWS(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLUPSIDEGAP2CROWS: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLUPSIDEGAP2CROWS: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLUPSIDEGAP2CROWS
 
// TA_CDLXSIDEGAP3METHODS is a wrapper for the CDLXSIDEGAP3METHODS function in TA-Lib
// single return value only inside index 0
func TA_CDLXSIDEGAP3METHODS(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLXSIDEGAP3METHODS: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLXSIDEGAP3METHODS(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLXSIDEGAP3METHODS: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLXSIDEGAP3METHODS: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLXSIDEGAP3METHODS

 
// TA_CDLABANDONEDBABY is a wrapper for the CDLABANDONEDBABY function in TA-Lib
// single return value only inside index 0
func TA_CDLABANDONEDBABY(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLABANDONEDBABY: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_CDLABANDONEDBABY: Provided objects is not a technicals.FeatureTechnical")
    }

    penetration := f.Args["penetration"]

    penetrationArg := C.double(penetration)

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLABANDONEDBABY(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLABANDONEDBABY: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLABANDONEDBABY: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLABANDONEDBABY
 
// TA_CDLDARKCLOUDCOVER is a wrapper for the CDLDARKCLOUDCOVER function in TA-Lib
// single return value only inside index 0
func TA_CDLDARKCLOUDCOVER(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLDARKCLOUDCOVER: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_CDLDARKCLOUDCOVER: Provided objects is not a technicals.FeatureTechnical")
    }

    penetration := f.Args["penetration"]

    penetrationArg := C.double(penetration)

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLDARKCLOUDCOVER(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLDARKCLOUDCOVER: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLDARKCLOUDCOVER: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLDARKCLOUDCOVER
 
// TA_CDLEVENINGDOJISTAR is a wrapper for the CDLEVENINGDOJISTAR function in TA-Lib
// single return value only inside index 0
func TA_CDLEVENINGDOJISTAR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLEVENINGDOJISTAR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_CDLEVENINGDOJISTAR: Provided objects is not a technicals.FeatureTechnical")
    }

    penetration := f.Args["penetration"]

    penetrationArg := C.double(penetration)

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLEVENINGDOJISTAR(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLEVENINGDOJISTAR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLEVENINGDOJISTAR: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLEVENINGDOJISTAR
 
// TA_CDLEVENINGSTAR is a wrapper for the CDLEVENINGSTAR function in TA-Lib
// single return value only inside index 0
func TA_CDLEVENINGSTAR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLEVENINGSTAR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_CDLEVENINGSTAR: Provided objects is not a technicals.FeatureTechnical")
    }

    penetration := f.Args["penetration"]

    penetrationArg := C.double(penetration)

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLEVENINGSTAR(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLEVENINGSTAR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLEVENINGSTAR: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLEVENINGSTAR
 
// TA_CDLMATHOLD is a wrapper for the CDLMATHOLD function in TA-Lib
// single return value only inside index 0
func TA_CDLMATHOLD(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLMATHOLD: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_CDLMATHOLD: Provided objects is not a technicals.FeatureTechnical")
    }

    penetration := f.Args["penetration"]

    penetrationArg := C.double(penetration)

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLMATHOLD(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLMATHOLD: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLMATHOLD: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLMATHOLD
 
// TA_CDLMORNINGDOJISTAR is a wrapper for the CDLMORNINGDOJISTAR function in TA-Lib
// single return value only inside index 0
func TA_CDLMORNINGDOJISTAR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLMORNINGDOJISTAR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_CDLMORNINGDOJISTAR: Provided objects is not a technicals.FeatureTechnical")
    }

    penetration := f.Args["penetration"]

    penetrationArg := C.double(penetration)

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLMORNINGDOJISTAR(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLMORNINGDOJISTAR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLMORNINGDOJISTAR: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLMORNINGDOJISTAR
 
// TA_CDLMORNINGSTAR is a wrapper for the CDLMORNINGSTAR function in TA-Lib
// single return value only inside index 0
func TA_CDLMORNINGSTAR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) { 

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLMORNINGSTAR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_CDLMORNINGSTAR: Provided objects is not a technicals.FeatureTechnical")
    }

    penetration := f.Args["penetration"]

    penetrationArg := C.double(penetration)

    output := make([]int, len(tw.Close))
    outPtr := (*C.int)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CDLMORNINGSTAR(
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
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CDLMORNINGSTAR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CDLMORNINGSTAR: TA-Lib Wrote no results")
    } else {
        return []float64{float64(output[outNBElement - 1])}, nil
    } 
} // TA_CDLMORNINGSTAR

//============STATISTIC FUNCTIONS=============

 
// TA_LINEARREG is a wrapper for the LINEARREG function in TA-Lib
// single return value only inside index 0
func TA_LINEARREG(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_LINEARREG: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_LINEARREG: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_LINEARREG(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_LINEARREG: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_LINEARREG: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_LINEARREG
 
// TA_LINEARREG_ANGLE is a wrapper for the LINEARREG_ANGLE function in TA-Lib
// single return value only inside index 0
func TA_LINEARREG_ANGLE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_LINEARREG_ANGLE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_LINEARREG_ANGLE: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_LINEARREG_ANGLE(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_LINEARREG_ANGLE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_LINEARREG_ANGLE: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_LINEARREG_ANGLE
 
// TA_LINEARREG_INTERCEPT is a wrapper for the LINEARREG_INTERCEPT function in TA-Lib
// single return value only inside index 0
func TA_LINEARREG_INTERCEPT(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_LINEARREG_INTERCEPT: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_LINEARREG_INTERCEPT: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_LINEARREG_INTERCEPT(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_LINEARREG_INTERCEPT: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_LINEARREG_INTERCEPT: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_LINEARREG_INTERCEPT
 
// TA_LINEARREG_SLOPE is a wrapper for the LINEARREG_SLOPE function in TA-Lib
// single return value only inside index 0
func TA_LINEARREG_SLOPE(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_LINEARREG_SLOPE: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_LINEARREG_SLOPE: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_LINEARREG_SLOPE(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_LINEARREG_SLOPE: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_LINEARREG_SLOPE: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_LINEARREG_SLOPE
 
// TA_TSF is a wrapper for the TSF function in TA-Lib
// single return value only inside index 0
func TA_TSF(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TSF: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_TSF: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_TSF(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_TSF: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_TSF: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_TSF
 
// TA_VAR is a wrapper for the VAR function in TA-Lib
// single return value only inside index 0
func TA_VAR(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_VAR: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_VAR: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    nbdev := f.Args["nbdev"]

    timePeriod := C.int(period)
    nbdevArg := C.double(nbdev)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_VAR(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        nbdevArg,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_VAR: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_VAR: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_VAR
 
// TA_STDDEV is a wrapper for the STDDEV function in TA-Lib
// single return value only inside index 0
func TA_STDDEV(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_STDDEV: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_STDDEV: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]
    nbdev := f.Args["nbdev"]

    timePeriod := C.int(period)
    nbdevArg := C.double(nbdev)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_STDDEV(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        timePeriod,
        nbdevArg,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_STDDEV: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_STDDEV: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_STDDEV

 
// TA_BETA is a wrapper for the BETA function in TA-Lib
// single return value only inside index 0
func TA_BETA(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_BETA: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_BETA: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_BETA(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        tw.OpenPtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_BETA: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_BETA: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_BETA
 

// TA_CORREL is a wrapper for the CORREL function in TA-Lib
// single return value only inside index 0
func TA_CORREL(
    ft Feature,
    tw *TALIBWrapper,
) ([]float64, error) {

    err := C.TA_Initialize()
    if err != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CORREL: TA-Lib failed to initialize")
    }
    defer C.TA_Shutdown()

    f, ok := ft.(FeatureTechnical)
    if !ok {
        return nil, fmt.Errorf("TA_CORREL: Provided objects is not a technicals.FeatureTechnical")
    }

    period := f.Args["timeperiod"]

    timePeriod := C.int(period)

    output := make([]float64, len(tw.Close))
    outPtr := (*C.double)(unsafe.Pointer(&output[0]))

    var outBegIdx, outNBElement C.int

    retCode := C.TA_CORREL(
        C.int(0), 
        C.int(len(tw.Close) - 1),
        tw.ClosePtr,
        tw.OpenPtr,
        timePeriod,
        &outBegIdx,
        &outNBElement,
        outPtr,
    )
    if retCode != C.TA_SUCCESS {
        return nil, fmt.Errorf("TA_CORREL: TA-Lib Computation failed")
    }

    if int(outNBElement) == 0 {
        return nil, fmt.Errorf("TA_CORREL: TA-Lib Wrote no results")
    } else {
        return []float64{output[outNBElement - 1]}, nil
    } 
} // TA_CORREL
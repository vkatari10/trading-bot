package tests 

import (
	"testing"
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
	json "github.com/vkatari10/trading-bot/src/runtime/go-src/json"
	eventloop "github.com/vkatari10/trading-bot/src/runtime/go-src/eventloop"
	"log"
	"time"
)

// func TestTalibSMA(t *testing.T) {
// 	runtimeData, err := json.NewRuntimeData("../test.json")
// 	if err != nil {
// 		t.Errorf("failed to create runtimeData Object")
// 	}

// 	eventloop.OverrideBurnIn(runtimeData)

// 	res, err := technicals.TA_SMA( // ~10-20 microseconds w/ len 31 window
// 		runtimeData.Objects[0],
// 		&runtimeData.OHLCV,
// 	)

// 	if err != nil {
// 		t.Errorf("TA SMA failed")
// 	} 

// 	if res == nil {
// 		t.Errorf("no valid result came back")
// 	}
// } // TestTalibSMA



func TestTalibBBANDS(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_BBANDS( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[0]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_BBANDS execution time: %v", end)

	if err != nil {
		t.Errorf("TA BBANDS failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_BBANDS invalid result came back")
	}
} // TestTalibBBANDS

func TestTalibDEMA(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_DEMA( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[1]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_DEMA execution time: %v", end)

	if err != nil {
		t.Errorf("TA DEMA failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_DEMA invalid result came back")
	}
} // TestTalibDEMA

func TestTalibEMA(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_EMA( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[2]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_EMA execution time: %v", end)

	if err != nil {
		t.Errorf("TA EMA failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_EMA invalid result came back")
	}
} // TestTalibEMA

func TestTalibHT_TRENDLINE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_HT_TRENDLINE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[3]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_HT_TRENDLINE execution time: %v", end)

	if err != nil {
		t.Errorf("TA HT_TRENDLINE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_HT_TRENDLINE invalid result came back")
	}
} // TestTalibHT_TRENDLINE

func TestTalibKAMA(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_KAMA( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[4]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_KAMA execution time: %v", end)

	if err != nil {
		t.Errorf("TA KAMA failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_KAMA invalid result came back")
	}
} // TestTalibKAMA

func TestTalibMA(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_MA( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[5]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_MA execution time: %v", end)

	if err != nil {
		t.Errorf("TA MA failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_MA invalid result came back")
	}
} // TestTalibMA

func TestTalibMIDPOINT(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_MIDPOINT( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[7]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_MIDPOINT execution time: %v", end)

	if err != nil {
		t.Errorf("TA MIDPOINT failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_MIDPOINT invalid result came back")
	}
} // TestTalibMIDPOINT

func TestTalibMIDPRICE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_MIDPRICE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[8]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_MIDPRICE execution time: %v", end)

	if err != nil {
		t.Errorf("TA MIDPRICE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_MIDPRICE invalid result came back")
	}
} // TestTalibMIDPRICE

func TestTalibSAR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_SAR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[9]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_SAR execution time: %v", end)

	if err != nil {
		t.Errorf("TA SAR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_SAR invalid result came back")
	}
} // TestTalibSAR

func TestTalibSAREXT(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_SAREXT( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[10]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_SAREXT execution time: %v", end)

	if err != nil {
		t.Errorf("TA SAREXT failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_SAREXT invalid result came back")
	}
} // TestTalibSAREXT

func TestTalibSMA(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_SMA( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[11]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_SMA execution time: %v", end)

	if err != nil {
		t.Errorf("TA SMA failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_SMA invalid result came back")
	}
} // TestTalibSMA

func TestTalibT3(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_T3( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[12]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_T3 execution time: %v", end)

	if err != nil {
		t.Errorf("TA T3 failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_T3 invalid result came back")
	}
} // TestTalibT3

func TestTalibTEMA(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_TEMA( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[13]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_TEMA execution time: %v", end)

	if err != nil {
		t.Errorf("TA TEMA failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_TEMA invalid result came back")
	}
} // TestTalibTEMA

func TestTalibTRIMA(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_TRIMA( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[14]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_TRIMA execution time: %v", end)

	if err != nil {
		t.Errorf("TA TRIMA failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_TRIMA invalid result came back")
	}
} // TestTalibTRIMA

func TestTalibWMA(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_WMA( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[15]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_WMA execution time: %v", end)

	if err != nil {
		t.Errorf("TA WMA failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_WMA invalid result came back")
	}
} // TestTalibWMA

func TestTalibADX(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_ADX( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[16]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_ADX execution time: %v", end)

	if err != nil {
		t.Errorf("TA ADX failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_ADX invalid result came back")
	}
} // TestTalibADX

func TestTalibADXR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_ADXR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[17]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_ADXR execution time: %v", end)

	if err != nil {
		t.Errorf("TA ADXR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_ADXR invalid result came back")
	}
} // TestTalibADXR

func TestTalibAPO(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_APO( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[18]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_APO execution time: %v", end)

	if err != nil {
		t.Errorf("TA APO failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_APO invalid result came back")
	}
} // TestTalibAPO

func TestTalibAROON(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_AROON( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[19]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_AROON execution time: %v", end)

	if err != nil {
		t.Errorf("TA AROON failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_AROON invalid result came back")
	}
} // TestTalibAROON

func TestTalibAROONOSC(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_AROONOSC( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[20]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_AROONOSC execution time: %v", end)

	if err != nil {
		t.Errorf("TA AROONOSC failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_AROONOSC invalid result came back")
	}
} // TestTalibAROONOSC

func TestTalibBOP(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_BOP( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[21]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_BOP execution time: %v", end)

	if err != nil {
		t.Errorf("TA BOP failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_BOP invalid result came back")
	}
} // TestTalibBOP

func TestTalibCCI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CCI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[22]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CCI execution time: %v", end)

	if err != nil {
		t.Errorf("TA CCI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CCI invalid result came back")
	}
} // TestTalibCCI

func TestTalibCMO(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CMO( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[23]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CMO execution time: %v", end)

	if err != nil {
		t.Errorf("TA CMO failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CMO invalid result came back")
	}
} // TestTalibCMO

func TestTalibDX(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_DX( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[24]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_DX execution time: %v", end)

	if err != nil {
		t.Errorf("TA DX failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_DX invalid result came back")
	}
} // TestTalibDX

func TestTalibMACD(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_MACD( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[25]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_MACD execution time: %v", end)

	if err != nil {
		t.Errorf("TA MACD failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_MACD invalid result came back")
	}
} // TestTalibMACD

func TestTalibMACDEXT(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_MACDEXT( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[26]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_MACDEXT execution time: %v", end)

	if err != nil {
		t.Errorf("TA MACDEXT failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_MACDEXT invalid result came back")
	}
} // TestTalibMACDEXT

func TestTalibMACDFIX(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_MACDFIX( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[27]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_MACDFIX execution time: %v", end)

	if err != nil {
		t.Errorf("TA MACDFIX failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_MACDFIX invalid result came back")
	}
} // TestTalibMACDFIX

func TestTalibMFI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_MFI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[28]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_MFI execution time: %v", end)

	if err != nil {
		t.Errorf("TA MFI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_MFI invalid result came back")
	}
} // TestTalibMFI

func TestTalibMINUS_DI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_MINUS_DI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[29]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_MINUS_DI execution time: %v", end)

	if err != nil {
		t.Errorf("TA MINUS_DI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_MINUS_DI invalid result came back")
	}
} // TestTalibMINUS_DI

func TestTalibMINUS_DM(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_MINUS_DM( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[30]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_MINUS_DM execution time: %v", end)

	if err != nil {
		t.Errorf("TA MINUS_DM failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_MINUS_DM invalid result came back")
	}
} // TestTalibMINUS_DM

func TestTalibMOM(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_MOM( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[31]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_MOM execution time: %v", end)

	if err != nil {
		t.Errorf("TA MOM failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_MOM invalid result came back")
	}
} // TestTalibMOM

func TestTalibPLUS_DI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_PLUS_DI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[32]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_PLUS_DI execution time: %v", end)

	if err != nil {
		t.Errorf("TA PLUS_DI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_PLUS_DI invalid result came back")
	}
} // TestTalibPLUS_DI

func TestTalibPLUS_DM(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_PLUS_DM( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[33]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_PLUS_DM execution time: %v", end)

	if err != nil {
		t.Errorf("TA PLUS_DM failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_PLUS_DM invalid result came back")
	}
} // TestTalibPLUS_DM

func TestTalibPPO(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_PPO( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[34]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_PPO execution time: %v", end)

	if err != nil {
		t.Errorf("TA PPO failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_PPO invalid result came back")
	}
} // TestTalibPPO

func TestTalibROC(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_ROC( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[35]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_ROC execution time: %v", end)

	if err != nil {
		t.Errorf("TA ROC failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_ROC invalid result came back")
	}
} // TestTalibROC

func TestTalibROCP(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_ROCP( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[36]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_ROCP execution time: %v", end)

	if err != nil {
		t.Errorf("TA ROCP failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_ROCP invalid result came back")
	}
} // TestTalibROCP

func TestTalibROCR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_ROCR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[37]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_ROCR execution time: %v", end)

	if err != nil {
		t.Errorf("TA ROCR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_ROCR invalid result came back")
	}
} // TestTalibROCR

func TestTalibROCR100(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_ROCR100( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[38]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_ROCR100 execution time: %v", end)

	if err != nil {
		t.Errorf("TA ROCR100 failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_ROCR100 invalid result came back")
	}
} // TestTalibROCR100

func TestTalibRSI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_RSI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[39]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_RSI execution time: %v", end)

	if err != nil {
		t.Errorf("TA RSI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_RSI invalid result came back")
	}
} // TestTalibRSI

func TestTalibSTOCH(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_STOCH( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[40]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_STOCH execution time: %v", end)

	if err != nil {
		t.Errorf("TA STOCH failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_STOCH invalid result came back")
	}
} // TestTalibSTOCH

func TestTalibSTOCHF(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_STOCHF( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[41]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_STOCHF execution time: %v", end)

	if err != nil {
		t.Errorf("TA STOCHF failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_STOCHF invalid result came back")
	}
} // TestTalibSTOCHF

func TestTalibSTOCHRSI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_STOCHRSI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[42]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_STOCHRSI execution time: %v", end)

	if err != nil {
		t.Errorf("TA STOCHRSI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_STOCHRSI invalid result came back")
	}
} // TestTalibSTOCHRSI

func TestTalibTRIX(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_TRIX( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[43]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_TRIX execution time: %v", end)

	if err != nil {
		t.Errorf("TA TRIX failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_TRIX invalid result came back")
	}
} // TestTalibTRIX

func TestTalibULTOSC(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_ULTOSC( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[44]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_ULTOSC execution time: %v", end)

	if err != nil {
		t.Errorf("TA ULTOSC failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_ULTOSC invalid result came back")
	}
} // TestTalibULTOSC

func TestTalibWILLR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_WILLR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[45]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_WILLR execution time: %v", end)

	if err != nil {
		t.Errorf("TA WILLR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_WILLR invalid result came back")
	}
} // TestTalibWILLR

func TestTalibAD(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_AD( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[46]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_AD execution time: %v", end)

	if err != nil {
		t.Errorf("TA AD failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_AD invalid result came back")
	}
} // TestTalibAD

func TestTalibADOSC(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_ADOSC( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[47]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_ADOSC execution time: %v", end)

	if err != nil {
		t.Errorf("TA ADOSC failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_ADOSC invalid result came back")
	}
} // TestTalibADOSC

func TestTalibOBV(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_OBV( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[48]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_OBV execution time: %v", end)

	if err != nil {
		t.Errorf("TA OBV failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_OBV invalid result came back")
	}
} // TestTalibOBV

func TestTalibHT_DCPERIOD(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_HT_DCPERIOD( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[49]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_HT_DCPERIOD execution time: %v", end)

	if err != nil {
		t.Errorf("TA HT_DCPERIOD failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_HT_DCPERIOD invalid result came back")
	}
} // TestTalibHT_DCPERIOD

func TestTalibHT_DCPHASE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_HT_DCPHASE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[50]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_HT_DCPHASE execution time: %v", end)

	if err != nil {
		t.Errorf("TA HT_DCPHASE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_HT_DCPHASE invalid result came back")
	}
} // TestTalibHT_DCPHASE

func TestTalibHT_PHASOR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_HT_PHASOR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[51]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_HT_PHASOR execution time: %v", end)

	if err != nil {
		t.Errorf("TA HT_PHASOR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_HT_PHASOR invalid result came back")
	}
} // TestTalibHT_PHASOR

func TestTalibHT_SINE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_HT_SINE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[52]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_HT_SINE execution time: %v", end)

	if err != nil {
		t.Errorf("TA HT_SINE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_HT_SINE invalid result came back")
	}
} // TestTalibHT_SINE

func TestTalibHT_TRENDMODE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_HT_TRENDMODE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[53]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_HT_TRENDMODE execution time: %v", end)

	if err != nil {
		t.Errorf("TA HT_TRENDMODE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_HT_TRENDMODE invalid result came back")
	}
} // TestTalibHT_TRENDMODE

func TestTalibAVGPRICE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_AVGPRICE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[54]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_AVGPRICE execution time: %v", end)

	if err != nil {
		t.Errorf("TA AVGPRICE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_AVGPRICE invalid result came back")
	}
} // TestTalibAVGPRICE

func TestTalibMEDPRICE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_MEDPRICE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[55]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_MEDPRICE execution time: %v", end)

	if err != nil {
		t.Errorf("TA MEDPRICE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_MEDPRICE invalid result came back")
	}
} // TestTalibMEDPRICE

func TestTalibTYPPRICE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_TYPPRICE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[56]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_TYPPRICE execution time: %v", end)

	if err != nil {
		t.Errorf("TA TYPPRICE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_TYPPRICE invalid result came back")
	}
} // TestTalibTYPPRICE

func TestTalibWCLPRICE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_WCLPRICE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[57]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_WCLPRICE execution time: %v", end)

	if err != nil {
		t.Errorf("TA WCLPRICE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_WCLPRICE invalid result came back")
	}
} // TestTalibWCLPRICE

func TestTalibATR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_ATR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[58]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_ATR execution time: %v", end)

	if err != nil {
		t.Errorf("TA ATR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_ATR invalid result came back")
	}
} // TestTalibATR

func TestTalibNATR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_NATR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[59]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_NATR execution time: %v", end)

	if err != nil {
		t.Errorf("TA NATR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_NATR invalid result came back")
	}
} // TestTalibNATR

func TestTalibTRANGE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_TRANGE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[60]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_TRANGE execution time: %v", end)

	if err != nil {
		t.Errorf("TA TRANGE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_TRANGE invalid result came back")
	}
} // TestTalibTRANGE

func TestTalibCDL2CROWS(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDL2CROWS( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[61]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDL2CROWS execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDL2CROWS failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDL2CROWS invalid result came back")
	}
} // TestTalibCDL2CROWS

func TestTalibCDL3BLACKCROWS(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDL3BLACKCROWS( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[62]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDL3BLACKCROWS execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDL3BLACKCROWS failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDL3BLACKCROWS invalid result came back")
	}
} // TestTalibCDL3BLACKCROWS

func TestTalibCDL3INSIDE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDL3INSIDE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[63]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDL3INSIDE execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDL3INSIDE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDL3INSIDE invalid result came back")
	}
} // TestTalibCDL3INSIDE

func TestTalibCDL3LINESTRIKE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDL3LINESTRIKE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[64]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDL3LINESTRIKE execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDL3LINESTRIKE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDL3LINESTRIKE invalid result came back")
	}
} // TestTalibCDL3LINESTRIKE

func TestTalibCDL3OUTSIDE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDL3OUTSIDE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[65]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDL3OUTSIDE execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDL3OUTSIDE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDL3OUTSIDE invalid result came back")
	}
} // TestTalibCDL3OUTSIDE

func TestTalibCDL3STARSINSOUTH(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDL3STARSINSOUTH( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[66]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDL3STARSINSOUTH execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDL3STARSINSOUTH failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDL3STARSINSOUTH invalid result came back")
	}
} // TestTalibCDL3STARSINSOUTH

func TestTalibCDL3WHITESOLDIERS(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDL3WHITESOLDIERS( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[67]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDL3WHITESOLDIERS execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDL3WHITESOLDIERS failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDL3WHITESOLDIERS invalid result came back")
	}
} // TestTalibCDL3WHITESOLDIERS

func TestTalibCDLABANDONEDBABY(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLABANDONEDBABY( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[68]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLABANDONEDBABY execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLABANDONEDBABY failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLABANDONEDBABY invalid result came back")
	}
} // TestTalibCDLABANDONEDBABY

func TestTalibCDLADVANCEBLOCK(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLADVANCEBLOCK( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[69]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLADVANCEBLOCK execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLADVANCEBLOCK failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLADVANCEBLOCK invalid result came back")
	}
} // TestTalibCDLADVANCEBLOCK

func TestTalibCDLBELTHOLD(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLBELTHOLD( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[70]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLBELTHOLD execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLBELTHOLD failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLBELTHOLD invalid result came back")
	}
} // TestTalibCDLBELTHOLD

func TestTalibCDLBREAKAWAY(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLBREAKAWAY( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[71]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLBREAKAWAY execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLBREAKAWAY failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLBREAKAWAY invalid result came back")
	}
} // TestTalibCDLBREAKAWAY

func TestTalibCDLCLOSINGMARUBOZU(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLCLOSINGMARUBOZU( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[72]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLCLOSINGMARUBOZU execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLCLOSINGMARUBOZU failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLCLOSINGMARUBOZU invalid result came back")
	}
} // TestTalibCDLCLOSINGMARUBOZU

func TestTalibCDLCONCEALBABYSWALL(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLCONCEALBABYSWALL( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[73]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLCONCEALBABYSWALL execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLCONCEALBABYSWALL failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLCONCEALBABYSWALL invalid result came back")
	}
} // TestTalibCDLCONCEALBABYSWALL

func TestTalibCDLCOUNTERATTACK(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLCOUNTERATTACK( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[74]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLCOUNTERATTACK execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLCOUNTERATTACK failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLCOUNTERATTACK invalid result came back")
	}
} // TestTalibCDLCOUNTERATTACK

func TestTalibCDLDARKCLOUDCOVER(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLDARKCLOUDCOVER( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[75]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLDARKCLOUDCOVER execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLDARKCLOUDCOVER failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLDARKCLOUDCOVER invalid result came back")
	}
} // TestTalibCDLDARKCLOUDCOVER

func TestTalibCDLDOJI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLDOJI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[76]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLDOJI execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLDOJI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLDOJI invalid result came back")
	}
} // TestTalibCDLDOJI

func TestTalibCDLDOJISTAR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLDOJISTAR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[77]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLDOJISTAR execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLDOJISTAR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLDOJISTAR invalid result came back")
	}
} // TestTalibCDLDOJISTAR

func TestTalibCDLDRAGONFLYDOJI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLDRAGONFLYDOJI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[78]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLDRAGONFLYDOJI execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLDRAGONFLYDOJI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLDRAGONFLYDOJI invalid result came back")
	}
} // TestTalibCDLDRAGONFLYDOJI

func TestTalibCDLENGULFING(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLENGULFING( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[79]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLENGULFING execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLENGULFING failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLENGULFING invalid result came back")
	}
} // TestTalibCDLENGULFING

func TestTalibCDLEVENINGDOJISTAR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLEVENINGDOJISTAR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[80]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLEVENINGDOJISTAR execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLEVENINGDOJISTAR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLEVENINGDOJISTAR invalid result came back")
	}
} // TestTalibCDLEVENINGDOJISTAR

func TestTalibCDLEVENINGSTAR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLEVENINGSTAR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[81]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLEVENINGSTAR execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLEVENINGSTAR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLEVENINGSTAR invalid result came back")
	}
} // TestTalibCDLEVENINGSTAR

func TestTalibCDLGAPSIDESIDEWHITE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLGAPSIDESIDEWHITE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[82]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLGAPSIDESIDEWHITE execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLGAPSIDESIDEWHITE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLGAPSIDESIDEWHITE invalid result came back")
	}
} // TestTalibCDLGAPSIDESIDEWHITE

func TestTalibCDLGRAVESTONEDOJI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLGRAVESTONEDOJI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[83]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLGRAVESTONEDOJI execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLGRAVESTONEDOJI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLGRAVESTONEDOJI invalid result came back")
	}
} // TestTalibCDLGRAVESTONEDOJI

func TestTalibCDLHAMMER(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLHAMMER( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[84]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLHAMMER execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLHAMMER failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLHAMMER invalid result came back")
	}
} // TestTalibCDLHAMMER

func TestTalibCDLHANGINGMAN(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLHANGINGMAN( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[85]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLHANGINGMAN execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLHANGINGMAN failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLHANGINGMAN invalid result came back")
	}
} // TestTalibCDLHANGINGMAN

func TestTalibCDLHARAMI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLHARAMI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[86]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLHARAMI execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLHARAMI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLHARAMI invalid result came back")
	}
} // TestTalibCDLHARAMI

func TestTalibCDLHARAMICROSS(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLHARAMICROSS( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[87]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLHARAMICROSS execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLHARAMICROSS failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLHARAMICROSS invalid result came back")
	}
} // TestTalibCDLHARAMICROSS

func TestTalibCDLHIGHWAVE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLHIGHWAVE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[88]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLHIGHWAVE execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLHIGHWAVE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLHIGHWAVE invalid result came back")
	}
} // TestTalibCDLHIGHWAVE

func TestTalibCDLHIKKAKE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLHIKKAKE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[89]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLHIKKAKE execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLHIKKAKE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLHIKKAKE invalid result came back")
	}
} // TestTalibCDLHIKKAKE

func TestTalibCDLHIKKAKEMOD(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLHIKKAKEMOD( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[90]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLHIKKAKEMOD execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLHIKKAKEMOD failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLHIKKAKEMOD invalid result came back")
	}
} // TestTalibCDLHIKKAKEMOD

func TestTalibCDLHOMINGPIGEON(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLHOMINGPIGEON( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[91]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLHOMINGPIGEON execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLHOMINGPIGEON failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLHOMINGPIGEON invalid result came back")
	}
} // TestTalibCDLHOMINGPIGEON

func TestTalibCDLIDENTICAL3CROWS(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLIDENTICAL3CROWS( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[92]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLIDENTICAL3CROWS execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLIDENTICAL3CROWS failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLIDENTICAL3CROWS invalid result came back")
	}
} // TestTalibCDLIDENTICAL3CROWS

func TestTalibCDLINNECK(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLINNECK( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[93]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLINNECK execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLINNECK failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLINNECK invalid result came back")
	}
} // TestTalibCDLINNECK

func TestTalibCDLINVERTEDHAMMER(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLINVERTEDHAMMER( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[94]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLINVERTEDHAMMER execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLINVERTEDHAMMER failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLINVERTEDHAMMER invalid result came back")
	}
} // TestTalibCDLINVERTEDHAMMER

func TestTalibCDLKICKING(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLKICKING( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[95]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLKICKING execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLKICKING failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLKICKING invalid result came back")
	}
} // TestTalibCDLKICKING

func TestTalibCDLKICKINGBYLENGTH(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLKICKINGBYLENGTH( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[96]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLKICKINGBYLENGTH execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLKICKINGBYLENGTH failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLKICKINGBYLENGTH invalid result came back")
	}
} // TestTalibCDLKICKINGBYLENGTH

func TestTalibCDLLADDERBOTTOM(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLLADDERBOTTOM( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[97]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLLADDERBOTTOM execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLLADDERBOTTOM failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLLADDERBOTTOM invalid result came back")
	}
} // TestTalibCDLLADDERBOTTOM

func TestTalibCDLLONGLEGGEDDOJI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLLONGLEGGEDDOJI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[98]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLLONGLEGGEDDOJI execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLLONGLEGGEDDOJI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLLONGLEGGEDDOJI invalid result came back")
	}
} // TestTalibCDLLONGLEGGEDDOJI

func TestTalibCDLLONGLINE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLLONGLINE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[99]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLLONGLINE execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLLONGLINE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLLONGLINE invalid result came back")
	}
} // TestTalibCDLLONGLINE

func TestTalibCDLMARUBOZU(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLMARUBOZU( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[100]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLMARUBOZU execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLMARUBOZU failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLMARUBOZU invalid result came back")
	}
} // TestTalibCDLMARUBOZU

func TestTalibCDLMATCHINGLOW(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLMATCHINGLOW( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[101]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLMATCHINGLOW execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLMATCHINGLOW failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLMATCHINGLOW invalid result came back")
	}
} // TestTalibCDLMATCHINGLOW

func TestTalibCDLMATHOLD(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLMATHOLD( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[102]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLMATHOLD execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLMATHOLD failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLMATHOLD invalid result came back")
	}
} // TestTalibCDLMATHOLD

func TestTalibCDLMORNINGDOJISTAR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLMORNINGDOJISTAR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[103]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLMORNINGDOJISTAR execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLMORNINGDOJISTAR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLMORNINGDOJISTAR invalid result came back")
	}
} // TestTalibCDLMORNINGDOJISTAR

func TestTalibCDLMORNINGSTAR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLMORNINGSTAR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[104]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLMORNINGSTAR execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLMORNINGSTAR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLMORNINGSTAR invalid result came back")
	}
} // TestTalibCDLMORNINGSTAR

func TestTalibCDLONNECK(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLONNECK( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[105]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLONNECK execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLONNECK failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLONNECK invalid result came back")
	}
} // TestTalibCDLONNECK

func TestTalibCDLPIERCING(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLPIERCING( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[106]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLPIERCING execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLPIERCING failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLPIERCING invalid result came back")
	}
} // TestTalibCDLPIERCING

func TestTalibCDLRICKSHAWMAN(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLRICKSHAWMAN( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[107]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLRICKSHAWMAN execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLRICKSHAWMAN failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLRICKSHAWMAN invalid result came back")
	}
} // TestTalibCDLRICKSHAWMAN

func TestTalibCDLRISEFALL3METHODS(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLRISEFALL3METHODS( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[108]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLRISEFALL3METHODS execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLRISEFALL3METHODS failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLRISEFALL3METHODS invalid result came back")
	}
} // TestTalibCDLRISEFALL3METHODS

func TestTalibCDLSEPARATINGLINES(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLSEPARATINGLINES( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[109]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLSEPARATINGLINES execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLSEPARATINGLINES failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLSEPARATINGLINES invalid result came back")
	}
} // TestTalibCDLSEPARATINGLINES

func TestTalibCDLSHOOTINGSTAR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLSHOOTINGSTAR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[110]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLSHOOTINGSTAR execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLSHOOTINGSTAR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLSHOOTINGSTAR invalid result came back")
	}
} // TestTalibCDLSHOOTINGSTAR

func TestTalibCDLSHORTLINE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLSHORTLINE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[111]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLSHORTLINE execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLSHORTLINE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLSHORTLINE invalid result came back")
	}
} // TestTalibCDLSHORTLINE

func TestTalibCDLSPINNINGTOP(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLSPINNINGTOP( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[112]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLSPINNINGTOP execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLSPINNINGTOP failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLSPINNINGTOP invalid result came back")
	}
} // TestTalibCDLSPINNINGTOP

func TestTalibCDLSTALLEDPATTERN(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLSTALLEDPATTERN( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[113]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLSTALLEDPATTERN execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLSTALLEDPATTERN failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLSTALLEDPATTERN invalid result came back")
	}
} // TestTalibCDLSTALLEDPATTERN

func TestTalibCDLSTICKSANDWICH(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLSTICKSANDWICH( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[114]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLSTICKSANDWICH execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLSTICKSANDWICH failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLSTICKSANDWICH invalid result came back")
	}
} // TestTalibCDLSTICKSANDWICH

func TestTalibCDLTAKURI(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLTAKURI( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[115]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLTAKURI execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLTAKURI failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLTAKURI invalid result came back")
	}
} // TestTalibCDLTAKURI

func TestTalibCDLTASUKIGAP(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLTASUKIGAP( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[116]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLTASUKIGAP execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLTASUKIGAP failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLTASUKIGAP invalid result came back")
	}
} // TestTalibCDLTASUKIGAP

func TestTalibCDLTHRUSTING(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLTHRUSTING( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[117]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLTHRUSTING execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLTHRUSTING failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLTHRUSTING invalid result came back")
	}
} // TestTalibCDLTHRUSTING

func TestTalibCDLTRISTAR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLTRISTAR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[118]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLTRISTAR execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLTRISTAR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLTRISTAR invalid result came back")
	}
} // TestTalibCDLTRISTAR

func TestTalibCDLUNIQUE3RIVER(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLUNIQUE3RIVER( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[119]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLUNIQUE3RIVER execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLUNIQUE3RIVER failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLUNIQUE3RIVER invalid result came back")
	}
} // TestTalibCDLUNIQUE3RIVER

func TestTalibCDLUPSIDEGAP2CROWS(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLUPSIDEGAP2CROWS( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[120]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLUPSIDEGAP2CROWS execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLUPSIDEGAP2CROWS failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLUPSIDEGAP2CROWS invalid result came back")
	}
} // TestTalibCDLUPSIDEGAP2CROWS

func TestTalibCDLXSIDEGAP3METHODS(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CDLXSIDEGAP3METHODS( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[121]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CDLXSIDEGAP3METHODS execution time: %v", end)

	if err != nil {
		t.Errorf("TA CDLXSIDEGAP3METHODS failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CDLXSIDEGAP3METHODS invalid result came back")
	}
} // TestTalibCDLXSIDEGAP3METHODS

func TestTalibBETA(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_BETA( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[122]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_BETA execution time: %v", end)

	if err != nil {
		t.Errorf("TA BETA failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_BETA invalid result came back")
	}
} // TestTalibBETA

func TestTalibCORREL(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_CORREL( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[123]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_CORREL execution time: %v", end)

	if err != nil {
		t.Errorf("TA CORREL failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_CORREL invalid result came back")
	}
} // TestTalibCORREL

func TestTalibLINEARREG(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_LINEARREG( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[124]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_LINEARREG execution time: %v", end)

	if err != nil {
		t.Errorf("TA LINEARREG failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_LINEARREG invalid result came back")
	}
} // TestTalibLINEARREG

func TestTalibLINEARREG_ANGLE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_LINEARREG_ANGLE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[125]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_LINEARREG_ANGLE execution time: %v", end)

	if err != nil {
		t.Errorf("TA LINEARREG_ANGLE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_LINEARREG_ANGLE invalid result came back")
	}
} // TestTalibLINEARREG_ANGLE

func TestTalibLINEARREG_INTERCEPT(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_LINEARREG_INTERCEPT( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[126]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_LINEARREG_INTERCEPT execution time: %v", end)

	if err != nil {
		t.Errorf("TA LINEARREG_INTERCEPT failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_LINEARREG_INTERCEPT invalid result came back")
	}
} // TestTalibLINEARREG_INTERCEPT

func TestTalibLINEARREG_SLOPE(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_LINEARREG_SLOPE( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[127]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_LINEARREG_SLOPE execution time: %v", end)

	if err != nil {
		t.Errorf("TA LINEARREG_SLOPE failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_LINEARREG_SLOPE invalid result came back")
	}
} // TestTalibLINEARREG_SLOPE

func TestTalibSTDDEV(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_STDDEV( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[128]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_STDDEV execution time: %v", end)

	if err != nil {
		t.Errorf("TA STDDEV failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_STDDEV invalid result came back")
	}
} // TestTalibSTDDEV

func TestTalibTSF(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_TSF( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[129]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_TSF execution time: %v", end)

	if err != nil {
		t.Errorf("TA TSF failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_TSF invalid result came back")
	}
} // TestTalibTSF

func TestTalibVAR(t *testing.T) {
	runtimeData, err := json.NewRuntimeData("../talib_test.json")
	if err != nil {
		t.Errorf("failed to create runtimeData Object")
	}

	eventloop.OverrideBurnIn(runtimeData)

    start := time.Now()
	res, err := technicals.TA_VAR( 
		technicals.Feature(runtimeData.TALIBFeatureTechnicals[130]),
		&runtimeData.OHLCV,
	)
    end := time.Since(start)
    log.Printf("TA_VAR execution time: %v", end)

	if err != nil {
		t.Errorf("TA VAR failed to run: %v", err)
	}

	if res == nil {
		t.Errorf("TA_VAR invalid result came back")
	}
} // TestTalibVAR

#ifndef LIVE_TECHNICALS_H
#define LIVE_TECHNICALS_H

/**
 * @brief A simple wrapper for double arrays that contains
 * both the data and the length of the array.
 */
typedef struct {
  /**
     @brief pointer to data
   */
  double * data;

  /**
     @brief length of the array
   */
  size_t len;

} DoubleArray;


/**
 * @brief Calculates the simple moving average of a given array of
 * doubles.
 *
 * @param arr the array containing doubles representing prices
 * @param len the length of the input array
 * @param window the window of the SMA calculation
 * @return the SMA of the given input array with a final size of
 * len - window
 */
DoubleArray sma(const double * arr, size_t len, unsigned int window);


/**
 * @brief Calculates the exponential moving average of a given array
 * of doubles.
 *
 * @param arr the array containing doubles representing prices
 * @param len the length of the input array
 * @param window the window of the EMA calculation
 * @param smoothing smoothing value to apply to the EMA calculation
 * @return the EMA of the given input array with a final size of
 * len - window
 */
DoubleArray ema(const double * arr, size_t len, unsigned int window,
              double smoothing);


/**
 * @brief calculates the upper Bollinger bands of a given
 * array of doubles
 *
 * @param arr the array containing doubles representing prices
 * @param len the legnth of the input array
 * @param sma the SMA to use in the calculation
 * @return an array containing the upper Bollinger bands
 * values
 */
double * bbands_upper(const double * arr, size_t len,
                      unsigned int sma);


/**
 * @brief calculates the lower Bollinger bands of a given
 * array of doubles
 *
 * @param arr the array containing doubles representing prices
 * @param len the legnth of the input array
 * @param sma the SMA to use in the calculation
 * @return an array containing the lowerBollinger bands
 * values
 */
double * bbands_lower(const double * arr, size_t len,
                      unsigned int sma);

/**
 * @brief Calculates Moving Average Convergence/Divergence
 * based on a given array of doubles by computing the
 * difference of two EMAs (ema1 - ema2)
 *
 * @param arr an array of doubles representing prices
 * @param ema1 the first EMA
 * @param ema32 the second EMA
 * @param smoothing the smoothing to apply to both ema calculations
 * @return the array containing the MCAD line
 */
double * macd(const double * arr, size_t len, unsigned int ema1,
              unsigned int ema2, double smoothing);

/**
 * @brief calculates the signal line of an array containing
 * the MACD values
 *
 * @param macd the array containing the MACD values
 * @param ema the EMA value to use to calculate the signal line
 * @param smoothing smoothing value to apply to EMA calculations
 * @return the signal line of the given mcad values
 */
double * macd_sig(const double * macd, size_t len, unsigned int ema,
                  unsigned int smoothing);



#endif // LIVE_TECHNICALS_H

#ifndef LIVE_UPDATES_H
#define LIVE_UDPATES_H

/**
 * @brief Gets the new ema value from a given a new price
 * 
 * @param old_ema the old ema value from the last interval
 * @param new_value the new price at the current interval
 * @param window the window of this EMA computation
 * @param smoothing the smoothing value to use in the EMA calculation
 */
double get_ema(double old_ema, double new_value, unsigned int window,
               double smoothing);

/**
 * @brief Gets the new sma value from a given the oldest
 * price and newest price in a given window
 * 
 * @param old_price the oldest price in the given window
 * @param new_price the current price in the given window
 * @param window the window of this SMA calculation
 */
double get_sma(double old_price, double new_price, 
  unsigned int window);

#endif // LIVE_UPDATES_H
/*
 * File containing methods to compute technical indicators
 * at runtime.
 *
 * Author: Vikas Katari
 * Date: 05/24/2025
 */

#include <stdio.h>
#include <stdlib.h>


double * sma (const double * arr, size_t len, unsigned int window) {
  return arr;
} // sma


double * ema (const double * arr, size_t len, unsigned int window,
              double smoothing) {
  return arr;
} // ema


double * bbands_upper(const double * arr, size_t len, unsigned int sma) {
 return arr;
} // bbands_upper


double * bbands_lower(const double * arr, size_t len, unsigned int sma) {
 return arr;
} // bbands_lower


double * macd(const double * arr, size_t len, unsigned int ema1, unsigned int ema2, double smoothing) {
  return arr;
} // macd


double * macd_sig(const double * arr, size_t len, unsigned int ema, unsigned int smoothing) {
  return arr;
} // macd_sig

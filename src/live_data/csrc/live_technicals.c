/*
 * File containing methods to compute technical indicators
 * at runtime.
 *
 * Author: Vikas Katari
 * Date: 05/24/2025
 */


#include <stdio.h>
#include <stdlib.h>
#include "live_stats.h" // for mean / standard deviation methods
#include "live_technicals.h"


DoubleArray sma(const double * arr, size_t len, unsigned int window) {

  DoubleArray result = {NULL, 0};

  if (window > len || len == 0) {
    return result;
  } // if

  double sum = 0.0;

  for (size_t i = 0; i < window; i++) {
    sum += arr[i];
  } // for

  size_t final_len = len - (size_t)window + 1;

  result.len = final_len;

  double * smas = (double *) malloc(sizeof(double) * final_len);

  for (size_t i = window; i <= len; i++) {

    smas[i - window] = sum / (double)window;

    sum -= arr[i - window];
    sum += arr[i];

  } // for

  result.data = smas;
  return result;
} // sma


DoubleArray ema(const double * arr, size_t len, unsigned int window,
              double smoothing) {

  DoubleArray result = {NULL, 0};

  if (len < window || len == 0) {
    return result;
  } // if

  double sum = 0.0;

  for (size_t i = 0; i < window; i++) {
    sum += arr[i];
  } // for

  size_t final_len = len - (size_t)window + 1;

  result.len = final_len;

  double alpha = (double)smoothing / (1 + window);

  double * emas = (double *) malloc(sizeof(double) * final_len);

  emas[0] = sum / (double)window;
  double old_ema = sum / (double)window;

  for (size_t i = window; i <= len; i++) {
    double new_ema = (arr[i] * alpha) + ((1 - alpha) * old_ema);
    emas[i - window + 1] = new_ema;
    old_ema = new_ema;
  } // for

  result.data = emas;
  return result;
} // ema


// update SMA

double get_ema(double old_value, double new_value, unsigned int window,
               double smoothing) {

  double alpha = smoothing / (double)window;
  return (new_value * alpha) + ((1- alpha) * old_value);

} // new_ema




double * bbands_upper(const double * arr, size_t len,
                      unsigned int sma) {

  if (len < sma) {
    return NULL;
  } // if


  static double dummy = 0.0;
  return &dummy;

} // bbands_upper


double * bbands_lower(const double * arr, size_t len,
                      unsigned int sma) {

  static double dummy = 0.0;
  return &dummy;
} // bbands_lower


double * macd(const double * arr, size_t len,
              unsigned int ema1, unsigned int ema2, double smoothing) {

  static double dummy = 0.0;
  return &dummy;
} // macd


double * macd_sig(const double * arr, size_t len,
                  unsigned int ema, unsigned int smoothing) {

static double dummy = 0.0;
  return &dummy;

} // macd_sig

void dummy_test() {
  printf("Testing from C library into Go\n");
  fflush(stdout);
} // dummy

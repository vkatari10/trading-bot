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


double * sma (const double * arr, size_t len, unsigned int window) {

  if (window > len || len == 0) {
    return NULL;
  } // if

  double sum = 0.0;

  for (size_t i = 0; i < window; i++) {
    sum += arr[i];
  } // for

  size_t final_len = len - (size_t)window + 1;

  double * smas = (double *) malloc(sizeof(double) * final_len);

  smas[0] = sum / (double)window;

  size_t front = 0;
  size_t end = window;
  size_t index = 1;

  for (size_t i = 0; i < final_len; i++) {

    sum -= arr[front];
    front++;

    sum += arr[end];
    end++;

    smas[index] = sum / (double)window;
    index++;

  } // for

  return smas;
} // sma


double * ema (const double * arr, size_t len, unsigned int window,
              double smoothing) {

  if (len < window) {
    return NULL;
  } // if

  double old_ema = arr[0];

  double * emas = (double *) malloc(sizeof(double) * len);
  emas[0] = old_ema;

  double alpha = (double)smoothing / (1 + window);

  for (size_t i = 1; i < len; i++) {
    double new_ema = (arr[i] * alpha) + ((1 - alpha) * old_ema);
    emas[i] = new_ema;
    old_ema = new_ema;
  } // for

  return emas;
} // ema


double * bbands_upper(const double * arr, size_t len,
                      unsigned int sma) {

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

int main(void) {

  double * data = (double *) malloc (sizeof(double) * 10);

  for (int i = 0; i < 10; i++) {
    data[i] = i + 1;
  } //for

  double * result = ema(data, 10, 3, 2);

  unsigned int final_len = 10 - 3 + 1;

  for (int i = 0; i < final_len; i++) {
    printf("result[i] = %lf\n", result[i]);
  } // for

  free(data);
  free(result);


  return 0;

} // main

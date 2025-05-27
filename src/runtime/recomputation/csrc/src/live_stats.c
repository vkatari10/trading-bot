/*
 * File containg statisical methods to be computed at runtime,
 * a dependency for the "live_technicals.c" file contained
 * in the same folder.
 *
 * Author: Vikas Katari
 * Date: 05/24/2025
 */


#include <stdio.h>
#include <stdlib.h>
#include <math.h>


double mean (const double * arr, size_t len) {

  double sum = 0.0;

  for (size_t i = 0; i < len; i++) {
    sum += arr[i];
  } // for

  return sum / len;
} // mean


double std_dev (const double * arr, size_t len) {

  double arr_mean = mean(arr, len);

  double sum = 0.0;

  for (size_t i = 0; i < len; i++) {
    double diff = arr[i] - arr_mean;
    sum += (diff * diff);
  } // for

  sum /= len;
  return sqrt(sum);
} // std_dev

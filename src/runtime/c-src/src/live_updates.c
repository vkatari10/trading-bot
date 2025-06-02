/*
Contains methods to update technicals in real time once 
the array of the given technical has already been computed
from "live_technicals.c"

Author: Vikas Katari
Date: 5/31/2025
*/

double new_ema(double old_ema, double new_value, unsigned int window,
               double smoothing) {

  double alpha = smoothing / (double)window;
  return (new_value * alpha) + ((1- alpha) * old_ema);

} // get_ema

double new_sma(double old_price, double new_price, 
  double sum,   unsigned int window) {

    sum -= old_price;
    sum += new_price;

    return sum / (double)window;

} // get_sma

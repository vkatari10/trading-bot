/**
 * This flie contains methods to recompute financial technical
 * indicators, to be used at runtime to quickly get new
 * technical values to advise the machine learning model.
 *
 * Author: Vikas Katari
 * Date: 05/19/2025
 */


#include <pybind11/pybind11.h>
#include <pybind11/numpy.h>
#include <cmath>
#include <iostream>
#include <stdexcept>
#include <deque>

namespace py = pybind11;


// =====================MOVING AVERAGES========================


/*
 * @brief Calculates the simple moving average of a given array of
 * prices. Let the window = n, then this method assumes that we have
 * n prices before averaged which is prices[0].
 *
 * @param prices the NumPy array containing doubles representing prices
 * @param window the window of the SMA
 * @return the SMA of the given prices array with the given windows with
 * a new array size of prices.size - window
 * @throws std::logic_error if the length of the array is less
 * than the window
 */
py::array_t<double> sma(py::array_t<double> prices,
                        unsigned int window) {

  py::buffer_info buf = prices.request();

  double *ptr = static_cast<double *>(buf.ptr);

  unsigned int size = buf.size;

  if (size < window) {
    throw std::logic_error(
       "Array size must be larger than the window"
    );
  } // if

  double sum = 0.0;
  std::deque<double> dq;

  for (unsigned int i = 0; i < window; i++) {
    dq.push_back(ptr[i]);
    sum += ptr[i];
  } // for

  std::vector<double> smas;
  smas.push_back(sum / window);

  for (unsigned int i = window; i < size; i++) {
    double old_value = dq.front();
    dq.pop_front();

    sum -= old_value;

    sum += ptr[i];
    dq.push_back(ptr[i]);

    smas.push_back(sum / window);
  } // for

  py::array_t<double> arr(smas.size(), smas.data());
  return arr;
} // sma


/*
 * @brief Calculates the exponential moving average of a given array
 * of prices. Let the window = n, then thsi methods asusmes that we
 * havenn prices before average which is prices[0].
 *
 * @param prices the NumPy arary containg double representing prices
 * @param window the window of the EMA
 * @param smoothing the smoothing value to apply to the EMA formula
 * defaults to
 */
py::array_t<double> ema(py::array_t<double> prices,
                        unsigned int window,
                        double smoothing = 2.0) {

  py::buffer_info buf = prices.request();

  double * ptr = static_cast<double *>(buf.ptr);

  unsigned int size = buf.size;

  if (size < window) {
   throw std::logic_error(
      "Array size must be larger than window"
    );
  }

  double sum = 0.0;

  for (unsigned int i = 0; i < window; i++) {
    sum += ptr[i];
  } // for

  std::vector<double> emas;

  double old_ema = sum / window;
  emas.push_back(old_ema);

  double alpha = 0.0;
  alpha = smoothing / (window + 1);

  for (unsigned int i = window; i < size; i++) {
    double new_ema = (ptr[i] * alpha) + ((1 - alpha) * old_ema);
    emas.push_back(new_ema);
    old_ema = new_ema;
  } // for

  py::array_t<double> arr(emas.size(), emas.data());
  return arr;
} // ema

/*
 * @brief Calculates the mean of a NumPy dataset
 * of doubles
 *
 * @param arr the NumPy array containing doubles
 * @return the mean value of the param arr
 */
double mean(py::array_t<double> arr) {

  py::buffer_info buf = arr.request();

  double * ptr = static_cast<double *>(buf.ptr);
  unsigned int size = buf.size;

  if (size == 0) {
    throw std::logic_error("Empty array");
  }

  double sum = 0.0;

  for (unsigned int i = 0; i < size; i++) {
    sum += ptr[i];
  } // for

  return sum / size;
} // mean

/*
 * @brief Calculates the standard deviation of a dataset
 * of doubles contained in an NumPy Array
 *
 * @param arr the NumPy array containing doubles
 * @return the standard deviation of the set
 */
double std_dev(py::array_t<double> arr) {

  py::buffer_info buf = arr.request();

  double arr_mean = mean(arr);
  double sum = 0.0;

  double * ptr = static_cast<double *>(buf.ptr);

  unsigned int size = buf.size;
  double diff;

  // sum items - mean in entire array
  for (unsigned int i = 0; i < size; i++) {
    diff = ptr[i] - arr_mean;
    sum += diff * diff;
  } // for

  sum /= size;
  return std::sqrt(sum);
} // std_dev

py::array_t<double> bbands_upper(py::array_t<double> arr) {
  throw std::logic_error("No implementation");
} // bbands_upper

py::array_t<double> bbands_lower(py::array_t<double> arr) {
  throw std::logic_error("No implementation");
} // bbands_lower

py::array_t<double> rsi(py::array_t<double> arr) {
  throw std::logic_error("No implementation");
} // rsi

py::array_t<double> macd(py::array_t<double> arr) {
  throw std::logic_error("No implementation");
} // macd

PYBIND11_MODULE(live_recompute, m) {
  m.def("sma", &sma);
  m.def("ema", &ema);
  m.def("mean", &mean);
  m.def("std_dev", &std_dev);
  m.def("bbands_upper", &bbands_upper);
  m.def("bbands_lower", &bbands_lower);
  m.def("rsi", &rsi);
  m.def("macd", &macd);
} // PYBIND11

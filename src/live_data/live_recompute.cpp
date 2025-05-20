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
#include <string>
#include <iostream>
#include <stdexcept>


namespace py = pybind11;


/*
 * @brief Calculates the simple moving average of a given array of
 * prices. Let the window = n, then this method assumes that we have
 * n prices before averaged which is prices[0].
 *
 * @param prices the NumPy array containing doubles representing prices
 * @param window the window of the SMA
 * @return the SMA of the given prices array with the given windows
 */
py::array_t<double> sma(py::array_t<double> prices, int window) {

  py::buffer_info buf = prices.request();

  double *ptr = static_cast<double *>(buf.ptr);
  int size = buf.size;

  int terms = 0;
  double total = 0.0;

  // new array
  std::vector<double> smas;

  for (int i = 0; i < size; i++) {
    total += ptr[i];

    double sma = total / window;
    smas.push_back(sma);
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
 * defaults to 2
 */
py::array_t<double> ema(py::array_t<double> prices, int window,
                        int smoothing = 2) {

  py::buffer_info buf = prices.request();

  double * ptr = static_cast<double *>(buf.ptr);
  int size = buf.size;

  std::vector<double> emas;
  emas.push_back(ptr[0]);

  for (int i = 1; i < size; i++) {
    double alpha = smoothing / (i + 1);
    double new_ema = alpha * ptr[i] + (1-alpha) * ptr[i-1];
    emas.push_back(new_ema);
  } // for

  py::array_t<double> arr(emas.size(), emas.data());
  return arr;

} // ema

py::array_t<double> bbands(py::array_t<double> arr) {
  throw std::logic_error("No implementation");
} // bbands

py::array_t<double> rsi(py::array_t<double> arr) {
  throw std::logic_error("No implementation");
} // rsi

py::array_t<double> macd(py::array_t<double> arr) {
  throw std::logic_error("No implementation");
} // macd

PYBIND11_MODULE(live_recomp, m) {
  m.def("sma", &sma);
  m.def("ema", &ema);
  m.def("bbands", &bbands);
  m.def("rsi", &rsi);
  m.def("macd", &macd);
} // PYBIND11

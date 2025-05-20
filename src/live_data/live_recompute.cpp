#include <pybind11/pybind11.h>
#include <pybind11/numpy.h>
#include <string>
#include <iostream>
#include <stdexcept>

namespace py = pybind11;

/**
 * This flie contains methods to recompute financial technical
 * indicators, to be used at runtime to quickly get new
 * technical values to advise the machine learning model.
 *
 * Author: Vikas Katari
 * Date: 05/19/2025
 */

py::array_t<double> sma(py::array_t<double> prices) {

  py::buffer_info buf = prices.request();

  double *ptr = static_cast<double *>(buf.ptr);
  int size = buf.size;

  int terms = 0;
  double total = 0.0;

  // new array
  std::vector<double> smas;

  for (int i = 0; i < size; i++) {
    total += ptr[i];
    terms += 1;

    double sma = total / terms;
    smas.push_back(sma);
  } // for

  py::array_t<double> arr(smas.size(), smas.data());
  return arr;
} // sma

py::array_t<double> ema(py::array_t<double> prices, int smoothing = 2) {

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

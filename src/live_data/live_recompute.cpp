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

/*
 * TODO
 * - sma
 * - ema
 * - bbands
 * - rsi
 * - macd (derived from ema)
 *
 */

py::array_t<double> sma(py::array_t<double> arr) {
  throw std::logic_error("No implementation");
} // sma

py::array_t<double> ema(py::array_t<double> arr) {
  throw std::logic_error("No implementation");
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

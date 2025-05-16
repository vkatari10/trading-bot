#include <pybind11/pybind11.h>
#include <pybind11/numpy.h>
#include <string>
#include <iostream>

// test file

void print() {
  std::cout << "poop" << std::endl;
}

PYBIND11_MODULE(example, m) {
  m.def("print", &print);
}

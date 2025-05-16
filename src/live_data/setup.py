from setuptools import setup, Extension
import pybind11
import numpy

ext_modules = [
    Extension(
        'example',
        ['live_recompute.cpp'],  # all your source files here
        include_dirs=[pybind11.get_include(), numpy.get_include()],
        language='c++',
    ),
]

setup(
    name='live_data',
    version='0.1',
    ext_modules=ext_modules,
    zip_safe=False,
)

"""
File containing Python CFFI logic to call 
C source risk code from other Python files

Author: Vikas Katari
Date: 08/19/2025
"""
from typing import List, Dict, Any

def new_ffi(dll_path="../librisk.so"):
    """Generates a new FFI object that allows Python to call methods from c-src"""
    from cffi import FFI
    ffi = FFI()

    # header file 
    ffi.cdef("""
        double max_drawdown(const double * equity, int n);
        double volatility(const double * equity, int n);
        double sharpe_ratio(const double * equity, int n, double risk_free_rate);
    """)

    C = ffi.dlopen("../librisk.so") # call functions above with C. prefix
    return ffi, C


def new_cffi_dispatch_table(ffi, C) -> Dict[str, Any]:
    """Generate a dispatch table for C risk methods"""
    return { # match with header file in c-include/
        "max_drawdown": C.max_drawdown,
        "sharpe": C.sharpe_ratio,
        "volatility": C.volatility
    }


def call_cffi_method(ffi, method: str, values: List[float], dispatch: Dict[str, Any], *args, **kwargs) -> float:
    """Uses the CFFI dispatch table to call methods"""
    values_c = ffi.new("double[]", values) # C rep of python list of floats
    res = dispatch[method](values_c, len(values), *args, **kwargs)
    return res


if __name__ == "__main__": # test

    ffi, C = new_ffi()
    c_dispatch = new_cffi_dispatch_table(ffi, C)

    equity = [100.0, 105.0, 103.0, 110.0, 90.0, 95.0]

    # Convert Python list to double*
    equity_c = ffi.new("double[]", equity)

    test = C.sharpe_ratio(equity_c, len(equity_c), 0.03)
    
    args = {"risk_free_rate": 0.03} # test user JSON

    res = call_cffi_method(ffi, "sharpe", equity, c_dispatch, *args.values())
    


"""
File containing Python CFFI logic to call 
C source risk code from other Python files

Author: Vikas Katari
Date: 08/19/2025
"""
def new_ffi(dll_path="../librisk.so"):

    from cffi import FFI

    ffi = FFI()

    ffi.cdef("""
        double max_drawdown(const double * equity, int n);
        double volatility(const double * equity, int n);
        double sharpe_ratio(const double * equity, int n, double risk_free_rate);
    """)

    C = ffi.dlopen("../librisk.so") # call functions above with C. prefix
    return ffi, C

def call_method(method: str) -> float:
    pass
    # TODO: implement wrapper methods here and then create a dispatch table but we just use kwargs ig
    # to force everything onto the same table   

if __name__ == "__main__": # test

    ffi, C = new_ffi()

    # Example Python data
    equity = [100.0, 105.0, 103.0, 110.0, 90.0, 95.0]

    # Convert Python list → C array (double[])
    equity_c = ffi.new("double[]", equity)

    print("Max Drawdown:", C.max_drawdown(equity_c, len(equity)))
    print("Volatility:", C.volatility(equity_c, len(equity)))
    print("Sharpe Ratio:", C.sharpe_ratio(equity_c, len(equity), 0.02))



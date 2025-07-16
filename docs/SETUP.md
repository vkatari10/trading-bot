# Setup

## Local Use

1. Clone this repo 

2. Install `TA-Lib` (The hardest part)
    - `TA-Lib` is a C library 
    - Installing it is far easier to do on Linux system but can also be done on Windows and MacOS
    - If you use Windows I suggest using `WSL` (Windows subsystem for Linux) which gives you a `Linux` machine on Windows 
    - I would follow instructions from the Python `TA-Lib` wrapper repo README which can be found [here](https://github.com/TA-Lib/ta-lib-python?tab=readme-ov-file), just to download and get the library onto your system
    - From the root run `./contrade_cli.py build`, this will attempt to compile the `Go` source code
    - If there is an error along the lines of could not find `-lta-lib` this just means that the one `Go` file could not find the `TA-Lib` install to link against 
    - There are two things you need to do 
        1. Find the location of the install and header files of `TA-Lib` on your system 
        2. Open this file `src/runtime/go-src/technicals/talib.go`
            - At the top is a comment starting with `#cgo`, this is the header linker path (CFLAGS), adjust path as needed
            - The second argument is the location of the actual compiled static or dynamic library (LDFLAGS), adjust path as needed (keep the `-lta-lib` and `-lm`)
        3. Keep calling `./contrade_cli.py build` until the Go module compiles without error.

3. Register for an `Alpaca` Brokerage Account
    - [Registration](https://app.alpaca.markets/account/login) is free!
    - Once signed in you can find your own API keys
    - Find and copy these values into the `.env.example` in the root of the project and rename this file to `.env` (important)
    - The other Alpaca links in the `.env` file are configured by default to use paper money, which involves no financial risk since these are simulated funds. 

4. Download python dependencies (The easiest part)
    - Run these commands:
    - `python3 -m venv venv` (Creates the virtual environment)
    - `source venv/bin/activate` (Start the virtual environment)
    - `pip install -r requirements.txt` (Download all python dependencies)


## Docker 

A docker-compose solution is coming soon to avoid complex `TA-Lib` installs
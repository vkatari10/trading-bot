'''
This file contains the method needed to interact with viewing and
managing positions

Modules used

- requests
- os
- dotenv
- alpaca_trade_api.rest

Author: Vikas Katari
Date: 04/24/2025
'''
# Alpaca imports
from alpaca_trade_api.rest import REST, TimeFrame

# Other imports
from dotenv import load_dotenv  # to retrieve API keys
import requests  # for API calls
import os
from typing import List, Dict, Any

# Exception imports
from src.exceptions.custom_exceptions import Forbidden, InvalidRequest, Liquidation

# Load environment variables from .env file
load_dotenv()

# API Keys
API_KEY = os.getenv("ALPACA_API_KEY")
SECRET_KEY = os.getenv("ALPACA_SECRET_KEY")

# Initialize the Alpaca client
trading_client = REST(API_KEY, SECRET_KEY,
                      base_url="https://paper-api.alpaca.markets")

# URLs for API calls
url = "https://paper-api.alpaca.markets/v2/positions"

headers = {
    "accept": "application/json",
    "APCA-API-KEY-ID": API_KEY,
    "APCA-API-SECRET-KEY": SECRET_KEY
}  # Headers for API calls


def get_all_positions() -> List[Dict[str, Any]]:
    '''
    Gets all open positions in an account.

    Args:
    None

    Returns:

    An array of JSON objects that provide information about each position

    Raises:

    None
    '''
    response = requests.get(url, headers=headers)
    return response.json()


def close_all_positions(cancel_orders: bool) -> List[Dict[str, Any]]:
    '''
    Closes all open positions in an account

    Args:

    cancel_orders: true if you want to cancel all open orders, else false

    Raises:

    Liquidation errors if a position fails to liquidate

    Returns:

    An array of closed positions objects

    '''
    if cancel_orders:
        close_url = url + "?cancel_orders=true"
    else:
        close_url = url + "?cancel_orders=false"

    response = requests.delete(close_url, headers=headers)

    if response.status_code == 500:
        raise Liquidation("Failed to liquidate")

    return response.json()

def get_position(ticker: str) -> Dict[str, Any]:
    '''
    Gets an open position based on a specific ticker

    Args:

    ticker: ticker for a stock

    Returns:

    An JSON objects representing the position
    '''
    symbol_url = url + "ticker"

    response = requests.get(symbol_url, headers=headers)

    return response.json()


def close_position(ticker: str) -> Dict[str, Any]:
    '''
    This methods closes out an entire position at market value

    Args:

    ticker: ticker of asset

    Returns:

    An JSON representing the closed position
    '''
    ticker_url = url + "/ticker"

    response = requests.delete(ticker_url, headers=headers)

    return response.json()


def close_position_shares(ticker: str, shares: str) -> Dict[str, Any]:
    '''
    This method closes a specified number of shares at market value.

    Args:

    ticker: ticker of the asset
    shares: the number of shares as a string to close

    Returns:

    An JSON object representing the closed position
    '''
    total_url = url + "/ticker?qty=" + shares

    response = requests.delete(total_url, headers=headers)

    return response.json()

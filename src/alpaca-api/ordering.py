'''
This file contains the methods needed to interact with the

Modules used

- requests
- os
- string
- dotenv
- alpaca_trade_api.rest

Author: Vikas Katari
Date: 04/21/2025

'''
# Alpaca imports
from alpaca_trade_api.rest import REST, TimeFrame

# Other imports
from dotenv import load_dotenv # to retrieve API keys
from typing import List # for typing methods
import requests # for API calls
import os
import string

# Load environment variables from .env file
load_dotenv()

# API Keys
API_KEY = os.getenv("ALPACA_API_KEY")
SECRET_KEY = os.getenv("ALPACA_SECRET_KEY")

# Initialize the Alpaca client
trading_client = REST(API_KEY, SECRET_KEY, base_url="https://paper-api.alpaca.markets")

# URLs for API calls
orders_url = "https://paper-api.alpaca.markets/v2/orders"

def place_market_order(time: string, symbol: string, qty: string, side: string, ext: bool) -> None:
    '''
    Places a order at market price through the Alpaca API.

    Args:

    time (str): time in force of the order
    symbol (str): symbol of ticker
    qty (str): amount of shares
    side (str): buy or sell
    ext (bool): enable day ext for time in force

    Returns:
    None.
    '''
    payload = {
        "type": "market",
        "time_in_force": time,
        "symbol": symbol,
        "qty": qty,
        "side": side,
        "extended_hours": ext
    } # Payload

    headers = {
        "accept": "application/json",
        "content-type": "application/json"
    } # Headers

    requests.post(orders_url, json=payload, headers=headers)

def place_limit_order(time: string, symbol: string, qty: string, side: string, ext: bool, limit_price: string) -> None:
    '''
    Places a order at a limit price through the Alpaca API.

    Args:

    time (str): time in force of the order
    symbol (str): symbol of ticker
    qty (str): amount of shares
    side (str): buy or sell
    ext (bool): enable day ext for time in force
    limit_price (str): the specified limit price

    Returns:
    None.
    '''
    payload = {
        "type": "limit",
        "time_in_force": time,
        "symbol": symbol,
        "qty": qty,
        "side": side,
        "limit_price": limit_price,
        "extended_hours": ext
    } # Payload

    headers = {
        "accept": "application/json",
        "content-type": "application/json"
    } # Headers

    respose = requests.post(orders_url, json=payload, headers=headers)

    print(respose.text)

def cancel_all_orders() -> None:
    '''Cancels all pending orders'''
    headers = {
        "accept": "application/json",
        "APCA-API-KEY-ID": API_KEY,
        "APCA-API-SECRET-KEY": SECRET_KEY
    } # Headers

    response = requests.delete(orders_url, headers=headers)

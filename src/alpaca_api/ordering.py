'''
This file contains the methods needed to interact with the ordering
and canceling of orders

Modules used

- requests
- os
- dotenv
- alpaca_trade_api.rest

Author: Vikas Katari
Date: 04/21/2025

'''
# Alpaca imports
from alpaca_trade_api.rest import REST, TimeFrame

# Other imports
from dotenv import load_dotenv  # to retrieve API keys
import requests  # for API calls
import os
from typing import List, Dict

# Exception imports
from src.exceptions.custom_exceptions import Forbidden, InvalidRequest

# Load environment variables from .env file
load_dotenv()

# API Keys
API_KEY = os.getenv("ALPACA_API_KEY")
SECRET_KEY = os.getenv("ALPACA_SECRET_KEY")

# Initialize the Alpaca client
trading_client = REST(API_KEY, SECRET_KEY,
                      base_url="https://paper-api.alpaca.markets")

# URLs for API calls
orders_url = "https://paper-api.alpaca.markets/v2/orders"


def place_market_order(symbol: str, qty: str, side: str) -> List[str, any]:
    '''
    Places a order at market price through the Alpaca API.

    Args:

    time (str): time in force of the order
    symbol (str): symbol of ticker
    qty (str): amount of shares
    side (str): buy or sell
    ext (bool): enable day ext for time in force

    Raises:

    Forbidden: if the Alpaca API returns "forbidden".
    InvalidRequest: if the request was considered invalid by Alpaca API.

    Returns:

    The str representing the ID of the order
    '''
    payload = {
        "type": "market",
        "time_in_force": "day",
        "symbol": symbol,
        "qty": qty,
        "side": side,
    }  # Payload

    headers = {
        "accept": "application/json",
        "content-type": "application/json",
        "APCA-API-KEY-ID": API_KEY,
        "APCA-API-SECRET-KEY": SECRET_KEY
    }  # Headers

    request = requests.post(orders_url, json=payload, headers=headers)

    # for bad requests
    if request.status_code == 403:
        raise Forbidden("Forbidden")
    elif request.status_code == 422:
        raise InvalidRequest("Request was invalid")

    return request.json()


def place_limit_order(time: str, symbol: str, qty: str, side: str,
                      ext: bool, limit_price: str) -> Dict[str, any]:
    '''
    Places a order at a limit price through the Alpaca API.

    Args:

    time (str): time in force of the order
    symbol (str): symbol of ticker
    qty (str): amount of shares
    side (str): buy or sell
    ext (bool): enable day ext for time in force
    limit_price (str): the specified limit price

    Raises:

    Forbidden: if the Alpaca API returns "forbidden".
    InvalidRequest: if the request was considered invalid by Alpaca API.

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
    }  # Payload

    headers = {
        "accept": "application/json",
        "content-type": "application/json",
        "APCA-API-KEY-ID": API_KEY,
        "APCA-API-SECRET-KEY": SECRET_KEY
    }  # Headers

    response = requests.post(orders_url, json=payload, headers=headers)

    if response.status_code == 403:
        raise Forbidden("Forbidden")
    elif response.status_code == 422:
        raise InvalidRequest("Invalid request")
    else:
        return response.json()


def cancel_all_orders() -> None:
    '''
    Cancels all pending orders

    Raises:

    Forbidden: if the Alpaca API returns "forbidden".
    InvalidRequest: if the request was considered invalid by Alpaca API.

    Returns:

    None
    '''
    headers = {
        "accept": "application/json",
        "APCA-API-KEY-ID": API_KEY,
        "APCA-API-SECRET-KEY": SECRET_KEY
    }  # Headers

    response = requests.delete(orders_url, headers=headers)

    if response.status_code == 403:
        raise Forbidden("Forbidden")
    elif response.status_code == 422:
        raise InvalidRequest("Invalid request")


def cancel_order(id: str) -> None:
    '''
    Cancels a pending order give the id of the order

    Raises:

    Forbidden: if the Alpaca API returns "forbidden".
    InvalidRequest: if the request was considered invalid by Alpaca API.

    Returns:

    None
    '''
    headers = {
        "APCA-API-KEY-ID": API_KEY,
        "APCA-API-SECRET-KEY": SECRET_KEY
    }  # Headers

    id_url = orders_url + "/" + id

    response = requests.delete(id_url, headers=headers)

    if response.status_code == 422:
        raise InvalidRequest("Order uncancelable")

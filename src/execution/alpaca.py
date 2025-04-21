# Alpaca imports
from alpaca_trade_api.rest import REST, TimeFrame  # Or use your specific Alpaca version import

# Other imports
from dotenv import load_dotenv
import os
import requests
import string

# Load environment variables from .env file
load_dotenv()

# API keys
API_KEY = os.getenv("ALPACA_API_KEY")
SECRET_KEY = os.getenv("ALPACA_SECRET_KEY")

# Initialize the Alpaca client
trading_client = REST(API_KEY, SECRET_KEY, base_url="https://paper-api.alpaca.markets")

# URLs for API calls
orders_url = "https://paper-api.alpaca.markets/v2/orders"

def cancel_all_orders() -> None:
    '''Cancels all open orders when called'''
    headers = {
        "accept": "application/json",
        "APCA-API-KEY-ID": API_KEY,
        "APCA-API-SECRET-KEY": SECRET_KEY
    }

    response = requests.delete(orders_url, headers=headers)


def place_order_limit(type: string, time: string, symbol: string, qty: int, side: string, limit: int) -> None:

    payload = {
        "type": type,
        "time_in_force": time,
        "symbol" : symbol,
        "qty": qty,
        "side" : side,
        "limit_price" : limit
    }

    headers = {
        "accept": "application/json",
        "content-type": "application/json"
    }

    response = requests.post(orders_url, json=payload, headers=headers)

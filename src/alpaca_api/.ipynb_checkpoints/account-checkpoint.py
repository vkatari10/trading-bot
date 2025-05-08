'''
This file contains the methods needed to interact with getting account information

Modules used
- request
- os
- dotenv
- alpaca_trade_api.rest

Author Vikas Katari
Date: 04/22/2025
'''
# Alpaca imports
from alpaca_trade_api.rest import REST, TimeFrame

# Other imports
from dotenv import load_dotenv  # to retrieve API keys
import requests  # for API calls
import os

# Exception imports
from src.exceptions.custom_exceptions import Forbidden, InvalidRequest

# Object imports
from src.alpaca_api.alpaca_classes import Asset

# Typing import
from typing import Dict, Any

# Load environment variables from .env file
load_dotenv()

# API Keys
API_KEY = os.getenv("ALPACA_API_KEY")
SECRET_KEY = os.getenv("ALPACA_SECRET_KEY")

# Initialize the Alpaca client
trading_client = REST(API_KEY, SECRET_KEY,
                      base_url="https://paper-api.alpaca.markets")

url = "https://paper-api.alpaca.markets/v2/account"

headers = {
    "accept": "application/json",
    "APCA-API-KEY-ID": API_KEY,
    "APCA-API-SECRET-KEY": SECRET_KEY
}


def get_account() -> Dict[str, Any]:
    '''
    Get account data

    Returns:

    a JSON object representing the account
    '''
    response = requests.get(url, headers=headers)
    return response.json()

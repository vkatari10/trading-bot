'''
Contains wrapper methods for the Finnhub API

Modules used
- finhub
- os
- requests
- dotenv

Author: Vikas Katari
Date: 04/25/2025
'''

import finnhub
import os
from dotenv import load_dotenv
import requests
from typing import Dict, Any

load_dotenv()

API_KEY = os.getenv("FINNHUB_API_KEY")
SECRET_KEY = os.getenv("FINNHUB_SECRET_API_KEY")

finnhub_client = finnhub.Client(api_key=API_KEY)

def get_quote(ticker: str) -> Dict[str, any]:
    '''

    '''
    return finnhub_client.quote(ticker)

print(get_quote("AAPL"))

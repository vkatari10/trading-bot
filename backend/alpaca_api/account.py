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

# Load environment variables from .env file
load_dotenv()

# API Keys
API_KEY = os.getenv("ALPACA_API_KEY")
SECRET_KEY = os.getenv("ALPACA_SECRET_KEY")

# Initialize the Alpaca client
trading_client = REST(API_KEY, SECRET_KEY,
                      base_url="https://paper-api.alpaca.markets")

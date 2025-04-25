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


url = "

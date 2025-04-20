# Alpaca imports
from alpaca_trade_api.rest import REST, TimeFrame  # Or use your specific Alpaca version import

# Other imports
from dotenv import load_dotenv
import os

# Load environment variables from .env file
load_dotenv()

API_KEY = os.getenv("ALPACA_API_KEY")
SECRET_KEY = os.getenv("ALPACA_SECRET_KEY")

# Initialize the Alpaca client
trading_client = REST(API_KEY, SECRET_KEY, base_url="https://paper-api.alpaca.markets")

# Get our account information
account = trading_client.get_account()

# Check if our account is restricted from trading
if account.trading_blocked:
    print('Account is currently restricted from trading.')

# Check how much money we can use to open new positions
print('${} is available as buying power.'.format(account.buying_power))

from dotenv import load_dotenv
import os

import time

from src.alpaca_api.ordering import place_market_order, cancel_all_orders, place_limit_order, cancel_order


API_KEY = os.getenv("ALPACA_API_KEY")
SECRET_KEY = os.getenv("ALPACA_SECRET_KEY")

id = place_limit_order("day", "AAPL", "2", "buy", True, "199")
print(id)

time.sleep(5)
cancel_order(id['id'])

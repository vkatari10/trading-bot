from dotenv import load_dotenv
import os

from src.alpaca_api.ordering import place_market_order, cancel_all_orders, place_limit_order, cancel_order


API_KEY = os.getenv("ALPACA_API_KEY")
SECRET_KEY = os.getenv("ALPACA_SECRET_KEY")

# place_limit_order("gtc", "AAPL", "2", "buy", False, "150")

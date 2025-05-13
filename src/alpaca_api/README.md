# alpaca_api

This file contains API wrapper methods to interact with the Alpaca Broker API to place orders and execute trades.

# Dependencies
- requests
- dotenv
- os

# Files

## account.py

This file includes a method to view the account metrics

## ordering.py

This file contains the methods to place market, limit orders, as well as canceling orders.

## positions.py

This file contains the methods to view and close current positions.

# Notes

These methods depend on the API keys for the Alpaca API to be stored in a `.env` file in the top level directory.
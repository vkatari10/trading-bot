# alpaca_api

This folder contains several classes to interact with the Alpaca API.<br>
More specifically, these files act as a way to place orders and manage positions by utilizing the Alpaca API. All classes here are just wrapper methods to interact with the API.

## account.py

This file includes a method to view the account metrics

## ordering.py

This file contains the methods to place market, limit orders, as well as canceling orders.

## positions.py

This file contains the methods to view and close current positions.

# Notes

These methods depend on the API keys for the Alpaca API to be stored in a `.env` file in the top level directory.
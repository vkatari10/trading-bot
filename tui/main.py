import contextlib
import time
from rich.console import Console
from rich.live import Live
from rich.table import Table

console = Console()

def make_table(coin_list):
    """Generate a Rich table from a list of coins"""
    table = Table(
        title=f"ConTrade - {time.asctime()}",
        style="black on grey66",
        header_style="white on dark_blue",
    )
    table.add_column("Symbol")
    table.add_column("Name", width=30)
    table.add_column("Price (USD)", justify="right")
    table.add_column("Volume (24h)", justify="right", width=16)
    table.add_column("Percent Change (7d)", justify="right", width=8)
    for coin in coin_list:
        symbol, name, price, volume, pct_change = (
            coin["symbol"],
            coin["name"],
            coin["price_usd"],
            f"{coin['volume24']:.2f}",
            float(coin["percent_change_7d"]),
        )
        pct_change_str = f"{pct_change:2.1f}%"
        if pct_change > 5.0:
            pct_change_str = f"[white on dark_green]{pct_change_str:>8}[/]"
        elif pct_change < -5.0:
            pct_change_str = f"[white on red]{pct_change_str:>8}[/]"
        table.add_row(symbol, name, price, volume, pct_change_str)
    return table

# Dummy data
raw_data = [
    {
        "symbol": "BTC",
        "name": "Bitcoin",
        "price_usd": "58250.00",
        "volume24": 35700000000,
        "percent_change_7d": 3.5,
    },
    {
        "symbol": "ETH",
        "name": "Ethereum",
        "price_usd": "3080.25",
        "volume24": 18200000000,
        "percent_change_7d": -6.3,
    },
    {
        "symbol": "ADA",
        "name": "Cardano",
        "price_usd": "0.375",
        "volume24": 950000000,
        "percent_change_7d": 7.1,
    },
    {
        "symbol": "SOL",
        "name": "Solana",
        "price_usd": "132.44",
        "volume24": 2200000000,
        "percent_change_7d": -2.4,
    },
    {
        "symbol": "XRP",
        "name": "Ripple",
        "price_usd": "0.486",
        "volume24": 1500000000,
        "percent_change_7d": 1.2,
    },
]

# Expand the list to simulate more entries
coins = raw_data * 5
num_coins = len(coins)
num_lines = 10

with Live(make_table(coins[:num_lines]), screen=True) as live:
    index = 0
    with contextlib.suppress(KeyboardInterrupt):
        while True:
            live.update(make_table(coins[index : index + num_lines]))
            time.sleep(0.5)
            index = (index + 1) % (num_coins - num_lines + 1)

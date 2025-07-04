from rich.live import Live
from rich.table import Table
import time

table = Table()
table.add_column("Number")

with Live(table, refresh_per_second=4) as live:
    for i in range(440):
        table.add_row(str(i))
        live.update(table)
  

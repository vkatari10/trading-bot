from textual.app import App, ComposeResult
from textual.containers import Horizontal, Vertical
from textual.widgets import Static, Button, Header, Footer, RichLog, DataTable, Select

from datetime import datetime


import asyncio
import websockets

from dotenv import load_dotenv
import os
import sys

from typing import List

import src.ml.json.json_parser as jp

import json

load_dotenv()

data_link = os.getenv("DATA_API_LINK")
log_link = os.getenv("LOG_API_LINK")

TICKERS = []

# temp data
SERVICE_STATUS = [  
    ("Service", "Status"),
    ("ENGINE", "[blink red]FAIL[/blink red]"),
    ("ML SERVER", "[blink red]FAIL[/blink red]"),
]   

ACCOUNT_DATA = [
    ("Item", "Value"),
    ("Account Value", 100000),
    ("Avail. Cash", 73594.43),
]

# def on_button_pressed(self, event: Button.Pressed) -> None:
#     button_id = event.button.id
#     info_panel = self.app.query_one(InfoPanel)
#     table = info_panel.query_one(DataTable)
#     table.clear()  # Clear old data

#     if button_id == "start_engine":
#         SERVICE_STATUS[1] = ("ENGINE", "[blink yellow]BURN IN[/blink yellow]")
#     elif button_id == "start_server":
#         SERVICE_STATUS[2] = ("ML SERVER", "[bold green]OK[/bold green]")
#     elif button_id == "exit":
#         self.app.exit()

#     table.add_rows(SERVICE_STATUS[1:])

# --- Panels ---


class LogPanel(Vertical):
    def compose(self):
        yield Static("[b]LOGS[/b]")
        yield RichLog(auto_scroll=True, id="main_log")

    def on_mount(self):
        log = self.query_one(RichLog)
        log.write("New Logs will appear here...")
        log.write("STARTUP CHECKLIST")
        log.write("1. Start ML server")
        log.write("2. Start engine")
        log.write("3. Start this TUI (You already did)")

    

class AccountPanel(Static): 
    def compose(self):
        yield Static("[b]ACCOUNT DATA[/b]")
        yield DataTable(id="account_table")

    def on_mount(self):
        table = self.query_one(DataTable)
        table.add_columns(*ACCOUNT_DATA[0])
        table.add_rows(ACCOUNT_DATA[1:])

class PositionPanel(Static):
    def compose(self):
        yield Static("POSITIONS")
        yield DataTable(id="data_table")

    def on_mount(self):
        table = self.query_one(DataTable)

class InfoPanel(Static):
    def compose(self) -> ComposeResult:
        yield Static("STATUS")
        yield DataTable(id="status_table")

    def on_mount(self):
        table = self.query_one(DataTable)
        table.add_columns(*SERVICE_STATUS[0])
        table.add_rows(SERVICE_STATUS[1:])

# --- Mid-section (Center) ---

class MidSection(Static):
    def compose(self):
        yield InfoPanel()
        # yield DataPanel()

    
class MidSection(Vertical):
    def compose(self):
        yield InfoPanel()
        yield AccountPanel()
        yield PositionPanel()

# --- Full Layout ---

class AllPanels(Horizontal):
    def compose(self):
        yield MidSection()
        yield LogPanel()

# --- Main App ---

class ConTrade(App):
    CSS_PATH = "contrade.tcss"

    def compose(self) -> ComposeResult:
        yield Header(show_clock=True)
        yield AllPanels()
        yield Footer()

    async def on_mount(self) -> None:
        log = self.query_one("#main_log", RichLog)
        data_table = self.query_one("#data_table", DataTable)
        log.write("mounting started")

        # init main data_table
        data_table.add_columns("Asset", "Close", "High", "Low", "Open", "Volume", *config.get_all_feature_label_names())

        # put place holder data
        tickers = config.get_live_stocks()
        placeholder = ["NULL"] * (5 + len(config.get_all_feature_label_names())) # account for OHLCV

        # put placeholders
        for i in range(len(tickers)):
            data_table.add_row(tickers[i], *placeholder, key=f"row_{i}")

        for i in range(len(tickers)):
            asyncio.create_task(listen_data_api(i, data_table, log))
            asyncio.create_task(listen_log_api(i, log))


        log.write(f"Added rows: {list(data_table.rows)}")

        data_table.update_cell_at((0, 2), 1)




async def listen_log_api(id: int, log=None):
   col_index = log_link.index(":")
   url = "ws" + log_link[col_index:] + f"/{id}"
   async with websockets.connect(url) as websocket:
       while True:
            raw_data = await websocket.recv()
            data = json.loads(raw_data)  # decode JSON string to dict
            log.write(data["msg"])

async def listen_data_api(id: int, table: DataTable, log: RichLog):
    col_index = data_link.index(":")
    url = "ws" + data_link[col_index:] + f"/{id}"
    log.write(f"[DATA {id}] Connecting to: {url}")

    async with websockets.connect(url) as websocket:
        while True:
            raw_data = await websocket.recv()
            data = json.loads(raw_data)

            for i in range(0, 5 + len(config.get_all_feature_label_names())):
                try:
                    # Safe to update
                    table.update_cell_at((id, i + 1), data[str(i)], update_width=True)
                except KeyError:
                    log.write(f"[MISSING] Row key {id} not found in table")
                except Exception as e:
                    log.write(f"[ERROR {id}] {str(e)}")

if __name__ == "__main__":

    args = sys.argv

    if len(args) == 0:
        raise ValueError("usage ./contrade_cli.py dash <CONFIG_FILE_PATH>")

    file_name = args[1]

    global config
    config = jp.UserConfig(file_name)
    

    ConTrade().run()

# TUI
from textual.app import App, ComposeResult
from textual.containers import Horizontal, Vertical
from textual.widgets import Static, Button, Header, Footer, RichLog, DataTable, Select
from rich.text import Text

# Other 
from datetime import datetime
import asyncio
import websockets
from dotenv import load_dotenv
import os
import sys
import json
import re
import subprocess
import signal

# Internal
import src.ml.json.json_parser as jp

load_dotenv()
data_link = os.getenv("DATA_API_LINK")
log_link = os.getenv("LOG_API_LINK")


# temp data 
SERVICE_STATUS = [
    ("Service", "Status"),
    ("ENGINE", "[blink red]OFF[/blink red]"),
    ("ML SERVER", "[blink red]OFF[/blink red]"),
]   



# --- Panels ---


class LogPanel(Static):
    def compose(self):
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
        _ = self.query_one(DataTable)
    
class PositionPanel(Static):
    def compose(self):
        yield Static("TECHNICALS")
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


class ControlPanel(Static):
    def compose(self) -> ComposeResult:
        yield Static("CONTROLS")
        yield Button("Start Engine", id="run_engine")
        yield Button("Start Server", id="mlapi")
        yield Button("STOP ALL", id="stop_engine")

    async def on_button_pressed(self, event: Button.Pressed) -> None:
        if event.button.id == "run_engine":
            await self.app.start_engine()
        elif event.button.id == "mlapi":
            await self.app.start_server()
        elif event.button.id == "stop_engine":
            self.app.stop_all()



# --- Full Layout ---

class AllPanels(Vertical):
    def compose(self):
        yield InfoPanel(id="info_panel")
        yield PositionPanel(id="position_panel")
        yield ControlPanel(id="control_panel")  
        yield LogPanel()
       

# --- Main App ---

class ConTrade(App):
    CSS_PATH = "contrade.tcss"

    async def start_engine(self):
        await asyncio.to_thread(self._start_engine_sync)

    def _start_engine_sync(self):
        self.engine_process = subprocess.Popen(
            ["go", "run", ".", file_name],
            cwd="./src/runtime/go-src",
            stdout=subprocess.DEVNULL,
            stderr=subprocess.DEVNULL,
            preexec_fn=os.setsid
        )

    async def start_server(self):
        await asyncio.to_thread(self._start_server_sync)

    def _start_server_sync(self) -> None:
        self.server_process = subprocess.Popen(
            ["python", "-m", "src.api.internal.model_api.fast_model_api", file_name],
            stdout=subprocess.DEVNULL,
            stderr=subprocess.DEVNULL
        )
    
    async def on_shutdown(self):
       await self.stop_all()


    def stop_all(self):
        log = self.query_one("#main_log", RichLog)
        log.write("[yellow]Attempting to stop all processes[/yellow]")

        if self.engine_process and self.engine_process.poll() is None:
            log.write("Stopping engine...")
            self.engine_process.kill()
            self.engine_process.wait()
        if self.server_process and self.server_process.poll() is None:
            log.write("[red]Stopping server...[/red]")
            self.server_process.terminate()
            self.server_process.wait()

        if self.engine_process and self.engine_process.poll() is None:
            os.killpg(os.getpgid(self.engine_process.pid), signal.SIGTERM)
            self.engine_process.wait()  



    def compose(self) -> ComposeResult:
        yield Header(show_clock=True)
        yield AllPanels()
        yield Footer()

    async def on_mount(self) -> None:
        log = self.query_one("#main_log", RichLog)
        data_table = self.query_one("#data_table", DataTable)
        log.write((Text("MOUNT START", style="red")))

        self.engine_process= None  
        self.server_process = None

        # init main data_table
        data_table.add_columns("Asset", "Close", "High", "Low", "Open", "Volume", *config.get_all_feature_label_names())

        # put place holder data
        tickers = config.get_live_stocks()
        placeholder = [0] * (5 + len(config.get_all_feature_label_names())) # account for OHLCV

        # put placeholders
        for i in range(len(tickers)):
            data_table.add_row(tickers[i], *placeholder, key=f"row_{i}")

        for i in range(len(tickers)):
            asyncio.create_task(listen_data_api(i, data_table, log))
            asyncio.create_task(listen_log_api(i, log))

        log.write(Text("MOUNT OK", style="bold green"))



async def listen_log_api(id: int, log=None):
   '''
   Connects to the logging end point to get logging information from 
   the runtime engine from all tickers
   '''
   col_index = log_link.index(":")
   url = "ws" + log_link[col_index:] + f"/{id}" 
   async with websockets.connect(url) as websocket:
       while True:
            raw_data = await websocket.recv()
            data = json.loads(raw_data)  # decode JSON string to dict
            timestamp = datetime.now().strftime("%H:%M:%S")
            log.write(Text(f"[{timestamp}] {data["msg"]}", overflow="fold"))


async def listen_data_api(id: int, table: DataTable, log: RichLog):
    '''
    Connects to the data end point to get data payloads from the runtime engine
    for all tickers
    '''
    col_index = data_link.index(":")
    url = "ws" + data_link[col_index:] + f"/{id}"
    async with websockets.connect(url) as websocket:
        while True:
            raw_data = await websocket.recv()
            data = json.loads(raw_data)

            for i in range(0, 5 + len(config.get_all_feature_label_names())):
                try:
                    old_val = str(table.get_cell_at((id, i + 1)))
                    plain = float(re.sub(r"\[/?[^\]]+\]", "", old_val))

                    new_val = float(data[str(i)])
            
                    if new_val > plain:
                        table.update_cell_at((id, i + 1), f"[green]{data[str(i)]:.2f}[/green]", update_width=True)
                    elif new_val < plain:
                        table.update_cell_at((id, i + 1), Text(f"[red]{data[str(i)]:.2f}[/red]", style="red").plain, update_width=True)
                    else:
                        table.update_cell_at((id, i + 1), Text(f"{data[str(i)]:.2f}").plain, update_width=True)
                except KeyError:
                    log.write(f"[MISSING] Row key {id} not found in table")
                except Exception as e:
                    log.write(f"[ERROR {id}] {str(e)}")


if __name__ == "__main__":

    args = sys.argv

    if len(args) == 0:
        raise ValueError("usage ./contrade_cli.py dash <CONFIG_FILE_PATH>")

    global file_name
    file_name = args[1]

    global config
    config = jp.UserConfig(file_name)
    
    ConTrade().run()

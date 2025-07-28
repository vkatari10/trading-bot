from textual.app import App, ComposeResult
from textual.containers import Horizontal, Vertical
from textual.widgets import Static, Button, Header, Footer, RichLog, DataTable, Select

from datetime import datetime



TICKERS = [
    ("AAPL", 0),
    ("GOOGL", 1),
    ("AMZN", 2)
]

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

POSITIONS = [
    ("Ticker", "Last", "Change", "Qty", "Avg. Cost", "Open P/L", "Day P/L", "Net Liq.", "Action"),
    ("AAPL", 202.50, 0.50, 1, 202.50, 3.20, -0.45, 0, "HOLD"),
    ("GOOGL", 143.75, -0.25, 2, 144.25, -1.00, -0.50, 287.50, "SELL"),
    ("AMZN", 126.50, 0.75, 5, 125.00, 7.50, 3.75, 632.50, "HOLD"),
    ("MSFT", 310.80, -1.20, 3, 312.00, -3.60, -3.60, 932.40, "BUY"),
    ("TSLA", 265.10, 1.10, 4, 260.00, 20.40, 4.40, 1060.40, "HOLD"),
    ("NVDA", 880.50, 2.50, 2, 875.00, 11.00, 5.00, 1761.00, "SELL"),
    ("META", 318.20, -0.80, 1, 319.00, -0.80, -0.80, 318.20, "BUY"),
    ("INTC", 34.95, 0.15, 10, 34.50, 4.50, 1.50, 349.50, "HOLD"),
    ("AMD", 123.60, -0.40, 6, 124.00, -2.40, -2.40, 741.60, "SELL"),
    ("NFLX", 428.70, 1.95, 1, 425.00, 3.70, 1.95, 428.70, "HOLD"),
    ("AAPL", 203.00, 0.30, 2, 202.50, 1.00, 0.60, 406.00, "BUY"),
    ("GOOGL", 144.20, -0.50, 3, 145.00, -2.40, -1.50, 432.60, "SELL"),
    ("AMZN", 127.25, 0.25, 4, 126.00, 5.00, 1.00, 509.00, "HOLD"),
    ("MSFT", 312.10, -1.00, 1, 313.00, -0.90, -1.00, 312.10, "BUY"),
    ("TSLA", 266.00, 0.90, 5, 265.00, 5.00, 4.50, 1330.00, "SELL"),
    ("NVDA", 882.25, -1.75, 2, 883.00, -1.50, -3.50, 1764.50, "HOLD"),
    ("META", 319.10, 0.40, 2, 318.00, 2.20, 0.80, 638.20, "BUY"),
    ("INTC", 35.05, -0.10, 8, 35.50, -3.60, -0.80, 280.40, "SELL"),
    ("AMD", 124.10, 0.50, 3, 123.00, 3.30, 1.50, 372.30, "HOLD"),
    ("NFLX", 429.80, -1.20, 1, 430.00, -0.20, -1.20, 429.80, "BUY"),
    ("AAPL", 204.10, 1.10, 2, 202.00, 4.20, 2.20, 408.20, "HOLD"),
    ("GOOGL", 144.50, 0.30, 2, 144.00, 1.00, 0.60, 289.00, "SELL"),
    ("AMZN", 128.00, 0.75, 5, 126.00, 10.00, 3.75, 640.00, "HOLD"),
    ("MSFT", 313.20, -0.80, 2, 314.00, -1.60, -1.60, 626.40, "BUY"),
    ("TSLA", 267.50, 1.50, 3, 265.00, 7.50, 4.50, 802.50, "SELL"),
    ("NVDA", 883.50, 1.00, 1, 882.00, 1.50, 1.00, 883.50, "HOLD"),
    ("META", 320.00, -1.00, 2, 321.00, -2.00, -2.00, 640.00, "BUY"),
    ("INTC", 35.10, 0.05, 6, 35.00, 0.60, 0.30, 210.60, "HOLD"),
    ("AMD", 124.60, -0.20, 4, 125.00, -1.60, -0.80, 498.40, "SELL"),
    ("NFLX", 431.20, 1.40, 1, 430.00, 1.20, 1.40, 431.20, "BUY"),
    ("AAPL", 205.00, 0.90, 1, 204.00, 1.00, 0.90, 205.00, "HOLD"),
    ("GOOGL", 145.00, 0.50, 3, 144.00, 3.00, 1.50, 435.00, "SELL"),
    ("AMZN", 129.50, 1.00, 2, 128.00, 3.00, 2.00, 259.00, "BUY"),
    ("MSFT", 314.00, 0.80, 2, 313.00, 2.00, 1.60, 628.00, "HOLD"),
    ("TSLA", 269.00, 1.50, 2, 267.00, 4.00, 3.00, 538.00, "SELL"),
    ("NVDA", 885.00, 1.50, 1, 884.00, 1.00, 1.50, 885.00, "BUY"),
    ("META", 321.50, 1.50, 2, 320.00, 3.00, 3.00, 643.00, "HOLD"),
    ("INTC", 35.20, 0.10, 5, 35.00, 1.00, 0.50, 176.00, "SELL"),
    ("AMD", 125.00, 0.40, 3, 124.00, 3.00, 1.20, 375.00, "BUY"),
    ("NFLX", 432.60, 1.40, 2, 431.00, 3.20, 2.80, 865.20, "HOLD"),
    ("AAPL", 206.20, 1.20, 1, 205.00, 1.20, 1.20, 206.20, "BUY"),
    ("GOOGL", 145.80, 0.80, 2, 145.00, 1.60, 1.60, 291.60, "HOLD"),
    ("AMZN", 130.20, 0.70, 1, 129.00, 1.20, 0.70, 130.20, "SELL"),
    ("MSFT", 315.50, 1.50, 3, 313.00, 7.50, 4.50, 946.50, "BUY"),
    ("TSLA", 270.50, 1.50, 2, 268.00, 5.00, 3.00, 541.00, "HOLD"),
    ("NVDA", 886.50, 1.50, 1, 885.00, 1.50, 1.50, 886.50, "SELL"),
    ("META", 323.00, 1.50, 1, 322.00, 1.00, 1.50, 323.00, "BUY"),
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
        yield Select(TICKERS)
        yield RichLog(auto_scroll=True)

    def on_mount(self):
        log = self.query_one(RichLog)
        log.write("New Logs will appear here...")
        log.write("STARTUP CHECKLIST")
        log.write("1. Start ML server")
        log.write("2. Start engine")
        log.write("3. ")

        filter = self.query_one(Select)
        filter.title = "Filter"

class AccountPanel(Static): 
    def compose(self):
        yield Static("[b]ACCOUNT DATA[/b]")
        yield DataTable()

    def on_mount(self):
        table = self.query_one(DataTable)
        table.add_columns(*ACCOUNT_DATA[0])
        table.add_rows(ACCOUNT_DATA[1:])

class PositionPanel(Static):
    def compose(self):
        yield Static("POSITIONS")
        yield DataTable(id="positions_table")

    def on_mount(self):
        table = self.query_one(DataTable)
        table.add_columns(*POSITIONS[0])
        table.add_rows(POSITIONS[1:])

class InfoPanel(Static):
    def compose(self) -> ComposeResult:
        yield Static("STATUS")
        yield DataTable()

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

if __name__ == "__main__":
    ConTrade().run()

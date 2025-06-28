from textual.app import App, ComposeResult
from textual.containers import Horizontal, Vertical
from textual.widgets import Button, Header, Footer, Static
from textual.reactive import reactive


class TopBar(Horizontal):
    def compose(self) -> ComposeResult:
        yield Button("Config", id="config", classes="menu-btn")
        yield Button("Env", id="env", classes="menu-btn")
        yield Button("Train", id="train", classes="menu-btn")
        yield Button("Backtest", id="backtest", classes="menu-btn")
        yield Button("Live Trade", id="trade", classes="menu-btn")


class MainPanel(Static):
    section = reactive("Welcome")

    def render(self) -> str:
        return f"# {self.section}\n\nThis is the `{self.section}` panel. Add your UI logic here."


class StratForgeTUI(App):
    CSS = """
    Screen {
        layout: vertical;
    }
    .menu-btn {
        border: none;
        content-align: center middle;
        height: 10;
        color: yellow;
    }
    .menu-btn:hover {
        background: #333333;
    }
    TopBar {
        height: 3;
        background: black;
        border: black;
        padding: 0 1;
    }
    MainPanel {
        padding: 2;
        border: black;
        height: 1fr;
    }
    """

    def compose(self) -> ComposeResult:
        yield Header()
        yield TopBar()
        self.panel = MainPanel()
        yield self.panel
        yield Footer()

    def on_button_pressed(self, event: Button.Pressed) -> None:
        self.panel.section = event.button.label


if __name__ == "__main__":
    app = StratForgeTUI()
    app.run()

from textual.app import App, ComposeResult
from textual.containers import Horizontal, Vertical
from textual.widgets import Button, Header, Footer, Static
from textual.reactive import reactive


class SideMenu(Vertical):
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
        layout: horizontal;
    }
    SideMenu {
        width: 20;
        background: black;
        border: heavy white;
    }
    .menu-btn {
        border: none;
        content-align: center middle;
        height: 3;
        background: black;
        color: white;
    }
    .menu-btn:hover {
        background: #333333;
    }
    MainPanel {
        padding: 2;
        border: heavy white;
        height: 1fr;
        width: 1fr;
    }
    """

    def compose(self) -> ComposeResult:
        yield SideMenu()
        yield MainPanel()
        yield Header()
        yield Footer()

    def on_button_pressed(self, event: Button.Pressed) -> None:
        main_panel = self.query_one(MainPanel)
        main_panel.section = event.button.label


if __name__ == "__main__":
    app = StratForgeTUI()
    app.run()

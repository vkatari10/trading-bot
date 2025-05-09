import backtrader as bt
import yfinance as yf

class SmaCross(bt.Strategy):
    params = dict(pfast=10, pslow=30)

    def __init__(self):
        sma1 = bt.ind.SMA(period=self.p.pfast)
        sma2 = bt.ind.SMA(period=self.p.pslow)
        self.crossover = bt.ind.CrossOver(sma1, sma2)

    def next(self):

        if len(self) < max(self.p.pfast, self.p.pslow):  # Important idk why yet
            return

        if not self.position and self.crossover > 0:
            self.buy()
        elif self.position and self.crossover < 0:
            self.close()

# Download data
df = yf.download("MSFT", start="2011-01-01", end="2012-12-31")
df.dropna(inplace=True)

# Because yfinance returns cols as a multiindex
df.columns = df.columns.get_level_values(0)

# Convert data to Backtrader feed
data = bt.feeds.PandasData(dataname=df)

cerebro = bt.Cerebro()
cerebro.addstrategy(SmaCross)
cerebro.adddata(data)

cerebro.run()
cerebro.plot()

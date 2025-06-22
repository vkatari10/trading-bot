'''
Backtesting file to test specific technical strategies

Modules Used:
- backtrader
- yfinance

Author: Vikas Katari
Date: 05/09/2025
'''
import backtrader as bt
import yfinance as yf

class CustomStrategy(bt.Strategy):
    '''
    SMAs -> slow cross fast
    EMAs -> 9, 12, 26 for MACD
    Bollinger Bands -> price cross above or below
    '''
    params = dict(sma_slow=30, sma_fast=10, macd1=12, macd2=26,
                  macdsig=9)


    def __init__(self):
        self.sma_slow = bt.ind.SMA(period=self.p.sma_slow)
        self.sma_fast = bt.ind.SMA(period=self.p.sma_fast)
        self.macd = bt.ind.MACD(self.data,
                                period_me1 = self.p.macd1,
                                period_me2 = self.p.macd2,
                                period_signal = self.p.macdsig)
        self.bands = bt.ind.BollingerBands()
        self.crossover = bt.ind.CrossOver(self.sma_slow, self.sma_fast)
        self.macd_cross = bt.ind.CrossOver(self.macd.macd, self.macd.signal)
        self.rsi = bt.ind.RSI()

    def next(self):


        if len(self) < max(self.p.sma_slow, 26): # 26 for MACD
            return

        signal = 0

        # Determine SMA Signals
        if self.crossover > 0:
            signal += 1
        elif self.crossover < 0:
            signal -= 1

        # Determine MACD Signal Line Cross
        if self.macd_cross > 0:
            signal += 1
        elif self.macd_cross < 0:
            signal -= 1

        # RSI
        if self.rsi < 45:
            signal += 1
        elif self.rsi > 50:
            signal -= 1

        # Bollinger Bands
        # Will not include right now in data/data_processing.py
        if self.data < self.bands.lines.bot:
            signal += 1
        elif self.data > self.bands.lines.top:
            signal -= 1

        if signal > 0:
            # if not self.position:
                self.buy(size=1)
        elif signal < 0:
            if self.position:
                self.close(size=1)


#                Adjust ticker          Adjust Dates
# Download data   vvvv            vvvvvvv          vvvvvvvvv
df = yf.download("CPNG", start="2021-01-01", end="2024-12-31")
df.dropna(inplace=True)

# Because yfinance returns cols as a multiindex
df.columns = df.columns.get_level_values(0)

# Convert data to Backtrader feed
data = bt.feeds.PandasData(dataname=df)

cerebro = bt.Cerebro()
cerebro.addstrategy(CustomStrategy)
cerebro.adddata(data)

cerebro.run()
cerebro.plot()

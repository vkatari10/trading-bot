'''
Class representing an account to be used during backtesting

Author: Vikas Katari
Date: 06/22/2025
'''
from collections import defaultdict

class Account():
    '''Models account'''

    def __init__(self, cash=10000.0, commission=0.00, batch_size=1):
        self.cash = cash
        self.commission = commission
        self.positions = []
        self.batch_size = batch_size

    def buy(self, price: float) -> None:

        cost = price * self.batch_size + self.commission

        if self.cash < cost: # cannot afford
            return

        self.positions.append(Position(start_price=price, shares=self.batch_size))
        self.cash -= cost
    
    def sell_all(self, end_price: float) -> None:
        if len(self.positions) == 0:
            return 
        
        proceeds = 0

        for pos in self.positions:
            proceeds += pos.findReturn(end_price)
        
        self.positions.clear()
        self.cash += proceeds

    def sell(self, end_price):
        if len(self.positions) == 0: # nothing to sell 
            return  
        
        pl = 0

        for i in range(len(self.positions)):
            pl += self.positions[i].findReturn(end_price)
        
        self.positions.clear()
        self.cash += pl


class Position():
    def __init__(self, start_price: float, shares: float):
        self.start_price = start_price
        self.shares = shares
    
    def setStart(self, price: float) -> None:
        self.start_price = price

    def findReturn(self, end_price: float) -> float:
        return self.shares * end_price


class RiskAccount():
    """
    The risk account needs to account for the risk of the portfolio when we are watching multiple assets

    In short we will be traversing multiple dataframes at the same time row by row to determine 
    if we need to buy or not 

    and then we need to evaluate the risk of the account (portfolio) at every step by calling the 
    C functions we have 

    for now we will just use account value itself rather than tracking 
    per stock stuff but we will need to consider that in the future since a specific
    stock is tied to the overall risk of the portfolio 

    We will want to track w/ each share not just proceeds
    """

    def __init__(self,
                 cash=10000,
                 commission=1,
                 batch_size=1
                 ):
        """Initialize a RiskAccount Object"""
        self.cash = cash
        self.commission = commission
        self.batch_size = batch_size
        self.positions = {} # contains positions
        self.risk_metrics = defaultdict(list())
        """
        Map positions from like TICKER: list of positions
        """

    def buy(self, ticker: str, entry_price: float) -> None:
        """Buy position of some stock?"""
        if self.cash < 0:
            return
        
        cost = entry_price * self.batch_size + self.commission

        if self.cash < cost: # cannot afford
            return

        self.positions[ticker].append(Position(start_price=entry_price, shares=self.batch_size))
        self.cash -= cost

    def sell_all(self, ticker: str, exit_price: float) -> None:
        """Sells alls position when signal model decides to sell"""
        ticker_positions = self.positions.get(ticker)

        if not ticker_positions or len(ticker_positions) == 0: # ticker DNE or no curr position on ticker
            return 
        
        pl = 0 

        for pos in ticker_positions:
            pl += pos.find_return(exit_price)

        self.positions.get(ticker).clear()
        self.cash += pl


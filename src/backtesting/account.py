'''
Class representing an account to be used during backtesting

Author: Vikas Katari
Date: 06/22/2025
'''

class Account():

    def __init__(self, cash=10000.0, comission=0.00):
        self.cash = cash
        self.comission = comission
        self.positions = list(Position)
    
    # thing for manging account anf stuff


class Position():
    def __init__(self, start_price: float, sell_price: float, 
                 shares: float):
        self.start_price = start_price
        self.sell_price = sell_price
        self.shares = shares
    
    def setStart(self, price: float) -> None:
        self.start_price = price

    def findReturn(self, end_price: float) -> float:
        return self.shares * (self.start_price - end_price)

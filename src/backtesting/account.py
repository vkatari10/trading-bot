'''
Class representing an account to be used during backtesting

Author: Vikas Katari
Date: 06/22/2025
'''

class Account():

    def __init__(self, cash=10000.0, comission=0.00, batch_size=1):
        self.cash = cash
        self.comission = comission
        self.positions = []
        self.batch_size = batch_size

    def buy(self, price):

        cost = price * self.batch_size + self.comission

        if self.cash < cost: # cannot afford
            return

        self.positions.append(Position(start_price=price, shares=self.batch_size))
        self.cash -= cost
        
    def sell(self, end_price):
        if len(self.positions) == 0: # nothing to sell 
            return  
        
        pl = 0

        for i in range(len(self.positions)):
            pl += self.positions[i].findReturn(end_price)
        
        self.positions.clear()
        self.cash += pl - self.comission


class Position():
    def __init__(self, start_price: float, shares: float):
        self.start_price = start_price
        self.shares = shares
    
    def setStart(self, price: float) -> None:
        self.start_price = price

    def findReturn(self, end_price: float) -> float:
        return self.shares * end_price

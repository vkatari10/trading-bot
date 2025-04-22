import string

class Asset:
    def __init__(self, id: string, as_class: string, exchange: string,
                 symbol: string, name: string, status: string, tradeable: bool,
                 marginable: bool, shortable: bool, easy_to_borrow: bool, fractionable: bool):
        self.id = id;
        self.as_class = as_class
        self.exchange = exchange
        self.symbol = symbol
        self.name = name
        self.status = status
        self.tradeable = tradeable
        self.marginable = marginable
        self.shortable = shortable
        self.easy_to_borrow = easy_to_borrow
        self.fractionable = fractionable

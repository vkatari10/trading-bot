from src.alpaca_api.ordering import place_market_order, place_limit_order, cancel_all_orders


user_input = ""

while (user_input != 'q'):
    user_input = input("Commands: x = place market order y = place limit  order z = cancel all orders: ")

    if (user_input.lower() == 'q'):
        break
    elif (user_input.lower() == 'x'):
        symbol = input("Ticker: ")
        qty = input("Shares: ")
        side = input("buy or sell: ")
        place_market_order(symbol, qty, side)
    elif (user_input.lower() == 'y'):
        time = input("Time in force? ")
        symbol = input("Ticker? ")
        qty = input("Shares? ")
        side = input("buy or sell? ")
        ext = input("After hours? True or False ")
        limit_price = input("Limit price? ")
        place_limit_order(time, symbol, qty, side, ext, limit_price)
    elif (user_input.lower() == 'z'):
        cancel_all_orders()
    else:
        print("invalid command")
        continue

print("done")

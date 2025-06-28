import json 
from typing import Dict

# features
# label log
# train stocks
# trade stocks
# model type
# model/ file name
# model training timeframe
# model training interval

def script() -> None:
    user_config = {}
    try: 
    
        print("Config Builder V1")
        print("""
    If you make ANY errors, you can either

    - Continue through this script
        - And then modify the written JSON file later
    - Or Quit the program 
        - And write any data that you have filled so far
            """)

        

    except KeyboardInterrupt:

        try:
            write_payload(user_config)
        except:
            print("Could not write file")
            return
    



def add_object(type: str, payload: Dict) -> None:
    """
    Add objects
    """
    pass

def add_labels(payload: Dict) -> None:
    pass

def add_train_stocks(payload: Dict) -> None:
    
    input = ""
    stocks = []
    count = 0

    print("Add a train stock, at least 1 required")

    while input != "done":

        input ("Add train stock ticker: ")
        

def add_trade_stocks(paylod: Dict) -> None:
    pass

def add_model_info(payload: Dict) -> None:
    pass

def write_payload(payload: Dict):
    filename = f"config/{payload["model_name"]}.json"

    with open(filename, 'w'):
        json.dumps(payload)
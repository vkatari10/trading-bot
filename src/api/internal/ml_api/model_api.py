'''
Contains API logic to expose the ML model trained in python to the Go
based runtime engine.

Modules used
- Flask
- Pickle
- JSON

Author: Vikas Katari
Date: 05/28/2025
'''

from flask import Flask
import pickle
import json

"""
Assume input jsons in src/logic look like the example

from go we get -> [SMA_10: 200, SMA_20: 205]

we need to fit this to model and let it predict something
and then push that back out for go to grab
"""

'''
The only responsbility of this is to recieve data and push
out predictions based on recieved data
'''

# Load in deciding ML model
with open('src/ml/models/decider/model.pkl', 'rb') as f:
    model = pickle.load(f)

# Load in user defined features
with open("src/logic/features.json") as f:
    features = json.load(f)

print(features)



app = Flask(__name__)
# Add find free socket method here


@app.route('/prediction')
def send_prediction():
    pass
    # import model
    # get data from the get_lastest data method
    # make prediction
    # send prediction back as JSON


@app.route('/data')
def get_lastest_data():
    pass
    # Go should send this back as a JSON
    # after getting the lastest market data and recomputing the
    # technicals we let GO return the data back as an array
    # that we can then feed to the ML models in the above
    # method to get a prediction and send back to Go


# Run API server
if __name__ == '__main__':
    app.run()

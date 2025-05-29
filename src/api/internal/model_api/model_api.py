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
import numpy as np

# Load in deciding ML model
with open('src/ml/models/decider/model.pkl', 'rb') as f:
    model = pickle.load(f)

with open('src/logic/features.json') as f:
    features = json.load(f)


app = Flask(__name__)
# Add find free socket method here


@app.route('api/ml-hotline', methods=['GET', 'POST'])
def send_prediction():
    '''
    The runtime engine (in this case Go) will Post an array as JSON
    to this endpoint and in this file we will use this array to
    use the ML model to predict with and then from this side
    Post again back to the same endpoint the prediction where
    Go can GET the new prediction and determine further action
    '''

    features  = get_new_features()
    features_np = np.array(features)

    prediction = model.predict(features_np)

    return prediction




# Run API server
if __name__ == '__main__':
    app.run()

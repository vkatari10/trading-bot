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

from flask import Flask, request, jsonify
import pickle
import json
import numpy as np
import src.ml.json.json_parser as jp # to get user config info
import sys


args = sys.argv     

config = None

if len(args) > 1:
    config = jp.UserConfig(args[-1]) # user config
else:
    raise ValueError("usage: ./contrade_cli mlapi <CONFIG_FILE_PATH>")


# Load in deciding ML model
file_path = "src/ml/models/decider/" + config.get_model_name()

with open(file_path, 'rb') as f:
    model = pickle.load(f)

# features_file = "config/" + os.getenv("FEATURE_CONFIG_FILE")
# with open(features_file) as f: MAY NOT BE NEEDED
#     features = json.load(f)

app = Flask(__name__)
# Add find free socket method here

final_pred = None

@app.route('/api/prediction', methods=['POST', 'GET'])
def send_prediction():
    '''
    This method allows the Go runtime engine to interact
    with the Machine Learning model in real time using
    a local server
    '''
    global final_pred

    if request.method == 'POST':
        data = request.get_json()

        features = []

        for i in range(len(data)):
            features.append(data[str(i)])
        
        # print(features)

        features_np = np.array(features).reshape(1, -1)

        prediction = model.predict(features_np)

        final_pred = int(prediction[0])

        return jsonify({"status": "received", "prediction": int(prediction[0])})

    elif request.method == 'GET':
        if final_pred is None:
            return jsonify({'status': 'none'}), 404
        return jsonify({'prediction': final_pred})
    
@app.route('/api/prediction_test', methods=['POST', 'GET'])
def send_prediction_test():
    '''
    This method is identical to the one above but for testing
    purposes without an ML model to ensure that the server
    structure works as intended
    '''
    global final_pred

    if request.method == 'POST':
        data = request.get_json()

        features = []

        for i in range(len(data)):
            features.append(data[str(i)])
        
        return jsonify({"status": "recieved", "prediction": int(1)})

    elif request.method == 'GET':
        if final_pred is None:
            return jsonify({'status': 'none'}), 404
        return jsonify({'prediction': final_pred})
        
# Run API server
if __name__ == '__main__':
    app.run(debug=False, port=5000)
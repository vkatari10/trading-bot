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

# Load in deciding ML model
with open('src/ml/models/decider/model.pkl', 'rb') as f:
    model = pickle.load(f)

with open('src/logic/features.json') as f:
    features = json.load(f)


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

        features_np = np.array(features).reshape(1, -1)

        prediction = model.predict(features_np)

        final_pred = int(prediction[0])

        return jsonify({"status": "recieved", "prediction": int(prediction[0])})

    elif request.method == 'GET':
        if final_pred is None:
            return jsonify({'status': 'none'}), 404
        return jsonify({'prediction': final_pred})

# Run API server
if __name__ == '__main__':
    app.run()

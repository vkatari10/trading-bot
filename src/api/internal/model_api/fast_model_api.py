'''
FastAPI implementation allowing Scikit-learn models
to communicate with Go

Author: Vikas Katari
Date: 07/07/2025
'''

from fastapi import FastAPI, BackgroundTasks, HTTPException, WebSocket
import uvicorn as uvi

# internal
import src.ml.json.json_parser as jp
import sys
import numpy as np
import pickle

# load in ML model based on CLI input

args = sys.argv

config = None

print(args)

if len(args) > 1:
    config = jp.UserConfig(args[-1])

    file_path = "src/ml/models/decider/" + config.get_model_name()

    with open(file_path, 'rb') as f:
        model = pickle.load(f)

elif len(args) == 1: # for tests
    pass
else:
    raise ValueError("usage: ./contrade_cli mlapi <CONFIG_FILE_PATH>")


# API Server

app = FastAPI()

@app.on_event("startup")
async def load_model():
    global model
    config = jp.UserConfig(args[-1])
    file_path = "src/ml/models/decider/" + config.get_model_name()
    print(file_path)
    with open(file_path, 'rb') as f:
        model = pickle.load(f)

def run_ml_prediction(data: dict):
    features = []

    #print(f"features -> {features}")

    for i in range(len(data)): # aligns features 
        features.append(data[str(i)])

    features_np = np.array(features).reshape(1, -1)

    prediction = model.predict(features_np)

    result = int(prediction[0])

    return result


@app.websocket("/results-ws/{job_id}")
async def results_websocket(websocket: WebSocket):
    await websocket.accept()

    while True:

        data = await websocket.receive_json()

        #print(f"CLIENT ===> {data}")

        result = run_ml_prediction(data) # int

        #print(f"SERVER SEND => {result}")

        await websocket.send_json({"result": result})


def test_ml_prediction(data: dict):

    features = []

    for i in range(len(data)): # aligns features 
        features.append(data[str(i)])

    # pretend we ran ML inference

    return 0

@app.websocket("/test-result")
async def test_websocket(websocket: WebSocket):

    await websocket.accept()

    data = await websocket.receive_json()

    result = test_ml_prediction(data)

    await websocket.send_json({"result": result})
    

if __name__ == "__main__":
    uvi.run(
        'src.api.internal.model_api.fast_model_api:app', reload=True,
        workers=len(config.get_live_stocks()) # 1:1 ratio of ticker to worker
    )
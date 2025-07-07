from fastapi import FastAPI, BackgroundTasks
import uuid

# internal
import src.ml.json.json_parser as jp
import sys
import numpy as np
import pickle
from typing import Dict, Any

# load in ML model based on CLI input
args = sys.argv

config = None

if len(args) > 1:
    config = jp.UserConfig(args[-1])
else:
    raise ValueError("usage: ./contrade_cli mlapi <CONFIG_FILE_PATH>")

file_path = "src/ml/models/decider" + config.get_model_name()

with open(file_path, 'rb') as f:
    model = pickle.load(f)

# job tracker (job id:data)
jobs: Dict[str, Any] = {}

# API Server

app = FastAPI()

def run_ml_prediction(job_id: str, data):
    features = []
    for i in range(len(data)): # aligns features 
        features.append(data[str(i)])

    features_np = np.array(features).reshape(1, -1)

    prediction = model.predict(features_np)

    result = int(prediction[0])

    jobs[job_id]["result"] = result
    jobs[job_id]["status"] = "done"

@app.post("/predict")
async def predict(data: dict, background_tasks: BackgroundTasks):
    job_id = data['job_number']
    jobs[job_id] = {"status": "pending", "result": None}
    background_tasks.add_task(run_ml_prediction, job_id, data)
    return {"status": "200"}

@app.get("/results/{job_id}")
def get_result(job_id: str):
    job = jobs.get(job_id)
    if not job:
        return {"error": "job not found"}
    if job["status"] != "done":
        return {"status": job["status"]}
    return {"status": "done", "result": job["result"]}

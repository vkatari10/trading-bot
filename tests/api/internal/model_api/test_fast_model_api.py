'''
Tests for ML API Server on Python's Side

Author: Vikas Katari
Date: 06/17/2025
'''
from fastapi.testclient import TestClient
from src.api.internal.model_api.fast_model_api import app

client = TestClient(app)

def test_ml_pred():
    with client.websocket_connect("/test-result") as websocket:
        dummy_data = {
            "0": 1.9,
            "1": 3.4,
            "2": 6.5
        }
        websocket.send_json(dummy_data)
        response = websocket.receive_json()

        assert "result" in response
        assert response["result"] == 0
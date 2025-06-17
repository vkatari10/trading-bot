'''
Tests for ML API Server on Python's Side

Author: Vikas Katari
Date: 06/17/2025
'''
from src.api.internal.model_api.model_api import app


def test_ping():
    client = app.test_client()       # create a test client
    response = client.get('/api/prediction')   # simulate GET /ping request

    assert response.status_code == 200

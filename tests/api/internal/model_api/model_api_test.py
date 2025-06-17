'''
Tests for ML API Server on Python's Side

Author: Vikas Katari
Date: 06/17/2025
'''
from src.api.internal.model_api.model_api import app


def test_ml_server():
    '''
    Tests the ML API server with a dummy endpoint to ensure
    JSONs can be recivied as intended
    '''
    client = app.test_client()      
    response_good = client.post(
        '/api/prediction_test', 
        json={"0": 1.23, "1": 4.56})   

    response_none = client.post(
        '/api/prediction_test',
    )

    assert response_good.status_code == 200 # 200 for proper input
    assert response_none.status_code == 415 # 415 for no input
    

    

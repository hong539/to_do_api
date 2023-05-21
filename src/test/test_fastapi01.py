from fastapi import FastAPI
import httpx
import pytest


app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Apple"}

import pytest
from httpx import AsyncClient

@pytest.mark.anyio()
async def test_root():
    async with httpx.AsyncClient(app=app, base_url="https://test") as client:
        response = await client.get("/")
    
    
    assert response.status_code == 200
    assert response.json() == {"message": "Apple"}
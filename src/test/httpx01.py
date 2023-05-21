import httpx
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
async def read_root():
    async with httpx.AsyncClient() as client:
        response = await client.get("https://www.example.com")
        return response.text
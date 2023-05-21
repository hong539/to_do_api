import asyncio
import httpx

async def main():
    async with httpx.AsyncClient() as client:
        response = await client.get('https://www.example.com/')
        print(response)


asyncio.run(main())
#python httpx01.py
#expected output: <Response [200 OK]>
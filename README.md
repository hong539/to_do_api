# to_do_api
Make a to_do backend api with Unkown lanugases. lol

## develope

```shell
poetry add fastapi
poetry add "uvicorn[standard]"
poetry add httpx

#Activating the virtual environment
poetry shell

#show commands help for uvicorn
uvicorn --help

#run server with uvicorn
uvicorn main:app --reload

#test
uvicorn httpx01:app --reload

#check URL
http://127.0.0.1:8000/docs
http://127.0.0.1:8000/redoc
```
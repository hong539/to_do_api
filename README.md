# to_do_api
Make a to_do backend api with Unkown lanugases. lol

## develope

```shell
poetry add fastapi
poetry add "uvicorn[standard]"

#Activating the virtual environment
poetry shell

#test
uvicorn main:app --reload

#check URL
http://127.0.0.1:8000/docs
http://127.0.0.1:8000/redoc
```
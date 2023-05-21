# to_do_api
Make a to_do backend api with Unkown lanugases. lol

## development

* HTTP/1.1
* HTTP/2
* HTTP/3

```shell
poetry add fastapi
poetry add "uvicorn[standard]"
poetry add httpx
poetry add flask
poetry add pytest

#Activating the virtual environment
poetry shell

#with uvicorn
#show commands help for uvicorn
uvicorn --help
#run server with uvicorn
uvicorn main:app --reload
uvicorn httpx01:app --reload

#with flask
flask --app test_flask01 run

#check URL
http://127.0.0.1:8000/docs
http://127.0.0.1:8000/redoc
```
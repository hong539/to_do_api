# to_do_api
Make a to_do backend api with Unkown lanugases. lol

## prerequisites

* Python == 3.8.16
* [pyenv+poetry](https://github.com/hong539/setup_dev_environment/blob/main/programing_languages/python/python.md)
* [diagrams](https://github.com/mingrammer/diagrams)
* [Graphviz](https://gitlab.com/graphviz/graphviz)
* flask
* qrcode
* Project dependcy detialls will be in pyproject.toml/poetry.lock

## development

* HTTP/1.1
* HTTP/2
* HTTP/3

```shell
poetry add qrcode
poetry add fastapi
poetry add "uvicorn[standard]"
poetry add httpx
poetry add flask
poetry add pytest
poetry add diagrams

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

python main.py
#test
curl -X POST -H "Content-Type: application/json" -d '{"text": "Hello, world!", "color": "#000000", "bgcolor": "#FFFFFF"}' http://localhost:5000/qrcode > qrcode.jpg
```
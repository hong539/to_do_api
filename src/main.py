from flask import Flask
import httpx


app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello, World!</p>"


with httpx.Client(app=app, base_url="https://testserver") as client:
    r = client.get("/")
    print(r.text)
    assert r.status_code == 200
    assert r.text == "<p>Hello, World!</p>"
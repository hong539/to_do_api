# to_do_api
Make a to_do api with golang and ChatGPT.


## Test

```shell
#Run with Go
go run main.go

#Test GetTodos
curl http://localhost:8000/todos

#Test GetTodo
curl http://localhost:8000/todos/{id}

#Test AddTodo
curl -X POST -H "Content-Type: application/json" -d '{"id":"001", "title":"play games", "completed":"Not yet"}' http://localhost:8000/todos

#Test UpdateTodo
curl -X PUT -H "Content-Type: application/json" -d '{"id":"001", "title":"play games", "completed":"true"}' http://localhost:8000/todos/{id}

#Test DeleteTodo
curl -X DELETE -H "Content-Type: application/json" -d '{"id":"001", "title":"play games", "completed":"true"}' http://localhost:8000/todos/{id}
```
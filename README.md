# Radix Demo App

Run the applications:
```shell
go run api/.
go run web/.
```

Test the responses:
```http request

### Get index from Web server
GET http://localhost:8000/

### Get test-api from Web server
GET http://localhost:8000/test-api

### Get api from API server
GET http://localhost:8001/api

```

# Radix Demo App

A simple web and api app as a radix quick start tutorial

## Topics:
- Registering a Application in Radix
- Github Deploykey and Webhooks
- Variables and Secrets
- Public and Internal communication
- Bonus: Monitoring

## Running locally
Run the applications:
```shell
go run ./api
go run ./web
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

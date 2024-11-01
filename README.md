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

# Test Web
curl http://localhost:8000
curl http://localhost:8000/test-api

# Test API
GET http://localhost:8001/api
```

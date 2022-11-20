#!/bin/sh
APP_TIMEOUT="2" \
DEBUG="true" \
APP_PORT="localhost:8001" \
GRPC_PORT="localhost:8002" \
USER_PORT="localhost:5002" \
AUTH_PORT="localhost:5003" \
go run services/api/main.go

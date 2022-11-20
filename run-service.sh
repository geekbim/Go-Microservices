#!/bin/sh
export $(cat ./environment/$1/.env | xargs)
go run services/$1/cmd/main.go
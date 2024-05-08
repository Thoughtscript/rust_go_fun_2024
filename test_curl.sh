#!/usr/bin/env bash

# Go Mongo DB

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST --insecure "https://localhost:443/api/example/create?id=1&name=a&val=a" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X PUT --insecure "https://localhost:443/api/example/update?id=1&name=c&val=c" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET --insecure "https://localhost:443/api/examples" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET --insecure "https://localhost:443/api/example/one?id=1" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X DELETE --insecure "https://localhost:443/api/example/delete?id=1" &

# Rust Actix

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:8000/"  &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST "http://localhost:8000/echo" --data "echoing" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:8000/hey" &

# Rust Mongo DB

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST "http://localhost:8000/example/create?id=1&name=a&val=a" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X PUT "http://localhost:8000/example/update?id=1&name=c&val=c" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:8000/examples" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:8000/example/one?id=1" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X DELETE "http://localhost:8000/example/delete?id=1" &
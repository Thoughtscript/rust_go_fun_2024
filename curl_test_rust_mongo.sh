#!/usr/bin/env bash

# Rust Mongo DB

printf "\n\n======================================================================================\n======================== BEGINNING RUST MONGO DB LIVE TESTS ==========================\n======================================================================================\n\n" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST "http://localhost:8000/example/create?id=1&name=a&val=a" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X PUT "http://localhost:8000/example/update?id=1&name=c&val=c" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:8000/examples" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:8000/example/one?id=1" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X DELETE "http://localhost:8000/example/delete?id=1" &
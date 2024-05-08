#!/usr/bin/env bash

# Go Mongo DB

printf "\n\n======================================================================================\n======================== BEGINNING GO MONGO DB LIVE TESTS ============================\n======================================================================================\n\n" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST --insecure "https://localhost:443/api/example/create?id=1&name=a&val=a" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X PUT --insecure "https://localhost:443/api/example/update?id=1&name=c&val=c" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET --insecure "https://localhost:443/api/examples" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET --insecure "https://localhost:443/api/example/one?id=1" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X DELETE --insecure "https://localhost:443/api/example/delete?id=1" &


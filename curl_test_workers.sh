#!/usr/bin/env bash

# Go Worker Queue

printf "\n\n======================================================================================\n======================== BEGINNING GO WORKER QUEUE LIVE TESTS ========================\n======================================================================================\n\n" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST --insecure "https://localhost:443/api/create?user=test&password=test&cmd=a&scheduled=2024-05-08T15:04:05Z" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET --insecure "https://localhost:443/api/status?user=test&password=test&uuid=5e44e582-7af4-4751-9dd8-9385c25e1e99" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET --insecure "https://localhost:443/api/workers?user=test&password=test" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET --insecure "https://localhost:443/api/jobs?user=test&password=test&uuid=5e44e582-7af4-4751-9dd8-9385c25e1e99" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST --insecure "https://localhost:443/api/stop?user=test&password=test&uuid=5e44e582-7af4-4751-9dd8-9385c25e1e99" &


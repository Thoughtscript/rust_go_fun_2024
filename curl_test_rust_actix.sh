#!/usr/bin/env bash

# Rust Actix

printf "\n\n======================================================================================\n======================== BEGINNING RUST ACTIX LIVE TESTS =============================\n======================================================================================\n\n" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:8000/"  &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST "http://localhost:8000/echo" --data "echoing" &

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:8000/hey" &


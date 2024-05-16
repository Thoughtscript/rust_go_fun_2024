#!/usr/bin/env bash

bash curl_test_go_mongo.sh
sleep 5 & bash curl_test_rust_actix.sh
sleep 10 & bash curl_test_rust_mongo.sh
sleep 15 & bash curl_test_workers.sh
# rust_go_fun

[![](https://img.shields.io/badge/Go-1.22.2-blue.svg)](https://golang.org/pkg/)
[![](https://img.shields.io/badge/Rust-2021-brown.svg)](https://doc.rust-lang.org/cargo/reference/manifest.html)
[![](https://img.shields.io/badge/Mongo-db-green.svg)](https://www.mongodb.com/)

Run the following from the root dir:

```bash
docker-compose up

# If using Docker Compose Engine V2:
docker compose up
```

Endpoint testing is supplied through `bash`:

```bash
bash curl_test_go_mongo.sh
bash curl_test_rust_actix.sh
bash curl_test_rust_mongo.sh
bash curl_test_workers.sh
# bash curl_test_all.sh
```

## Go

1. https://localhost:443/public/
1. https://localhost:443/public/queryworker.html
1. https://localhost:443/public/querystatus.html
1. https://localhost:443/public/querystatuses.html
1. https://localhost:443/public/addworker.html
1. https://localhost:443/public/stopworker.html

Defaults:

* User and password for basic auth: `test`, `test`
* `a` and `b` are commands that are supplied in the request - they correspond to the bash scripts in [/go/bin](/go/bin)
* Simple password auth isn't enterprise-grade security (obviously) but is added to demonstrate basic auth checking

### Worker

1. `POST` https://localhost:443/api/create?user=test&password=test&cmd=a&scheduled=2024-05-08T15:04:05Z
1. `GET` https://localhost:443/api/status?user=test&password=test&uuid=5e44e582-7af4-4751-9dd8-9385c25e1e99
1. `GET` https://localhost:443/api/workers?user=test&password=test
1. `GET` https://localhost:443/api/jobs?user=test&password=test&uuid=5e44e582-7af4-4751-9dd8-9385c25e1e99
1. ` POST` https://localhost:443/api/stop?user=test&password=test&uuid=5e44e582-7af4-4751-9dd8-9385c25e1e99

### Example

Mongo DB:

1. `GET` https://localhost:443/api/examples
1. `GET` https://localhost:443/api/example/one?id=1
1. `DELETE` https://localhost:443/api/example/delete?id=1
1. `POST` https://localhost:443/api/example/create?id=1&name=a&val=a
1. `PUT` https://localhost:443/api/example/update?id=1&name=c&val=c

## Rust

Working my way through some Rust and Actix Server documentation.

1. http://localhost:8000/
1. http://localhost:8000/echo
1. http://localhost:8000/hey

### Example

1. `GET` http://localhost:8000/examples
1. `GET` http://localhost:8000/example/one?id=1
1. `DELETE` http://localhost:8000/example/delete?id=1
1. `POST` http://localhost:8000/example/create?id=1&name=a&val=a
1. `PUT` http://localhost:8000/example/update?id=1&name=c&val=c

## Resources and Links

1. https://doc.rust-lang.org/rust-by-example/index.html
2. https://ntex.rs/docs/application
3. https://hub.qovery.com/guides/tutorial/create-a-blazingly-fast-api-in-rust-part-1/#overview
4. https://actix.rs/docs/getting-started
5. https://github.com/tfogo/mongodb-go-tutorial/blob/master/main.go
6. https://pkg.go.dev/net/url#URL.Query
7. https://stackoverflow.com/questions/55564562/what-is-the-bson-syntax-for-set-in-updateone-for-the-official-mongo-go-driver
8. https://github.com/Mr-Malomz/actix-mongo-api/blob/main/src/repository/mongodb_repo.rs
9. https://crates.io/crates/mongodb

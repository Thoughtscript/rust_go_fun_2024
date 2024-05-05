# rust_go_fun

[![](https://img.shields.io/badge/Go-1.22.2-blue.svg)](https://golang.org/pkg/)
[![](https://img.shields.io/badge/Rust-2021-brown.svg)](https://doc.rust-lang.org/cargo/reference/manifest.html)
[![](https://img.shields.io/badge/Mongo-db-green.svg)](https://www.mongodb.com/)

## Go

1. https://localhost/public/
1. https://localhost/public/queryworker.html
1. https://localhost/public/querystatus.html
1. https://localhost/public/querystatuses.html
1. https://localhost/public/addworker.html
1. https://localhost/public/stopworker.html

### Example

1. https://localhost/api/examples
1. https://localhost/api/example/one?id=1
1. https://localhost/api/example/delete?id=1
1. https://localhost/api/example/create?id=1&name=a&val=a
1. https://localhost/api/example/update?id=1&name=c&val=c

Defaults:

* User and password for basic auth: `test`, `test`
* `a` and `b` are commands that are supplied in the request - they correspond to the bash scripts in [/go/bin](/go/bin)

## Rust

Working my way through some Rust and Actix Server documentation.

1. http://localhost:8000/
1. http://localhost:8000/echo
1. http://localhost:8000/hey

## Resources and Links

1. https://doc.rust-lang.org/rust-by-example/index.html
2. https://ntex.rs/docs/application
3. https://hub.qovery.com/guides/tutorial/create-a-blazingly-fast-api-in-rust-part-1/#overview
4. https://actix.rs/docs/getting-started
5. https://github.com/tfogo/mongodb-go-tutorial/blob/master/main.go
6. https://pkg.go.dev/net/url#URL.Query
7. https://stackoverflow.com/questions/55564562/what-is-the-bson-syntax-for-set-in-updateone-for-the-official-mongo-go-driver
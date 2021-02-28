# Go Gin Boilerplate

> A starter project with Golang, Gin and DynamoDB

[![Go Version][go-image]][go-url]
[![License][license-image]][license-url]

Golang Gin boilerplate with DynamoDB resource. Supports multiple configuration environments.

![](header.jpg)

### Boilerplate structure

```
.
├── Makefile
├── Procfile
├── README.md
├── config
│   ├── config.go
│   ├── dev.yaml
│   ├── prod.yaml
│   ├── stage.yaml
│   └── test.yaml
├── controllers
├── db
│   └── db.go
├── header.jpg
├── main.go
├── middlewares
│   └── auth.go
├── models
└── server
    ├── router.go
    └── server.go
```

## Installation

```sh
make deps
```

## Usage example

`curl http://localhost:8888/health`

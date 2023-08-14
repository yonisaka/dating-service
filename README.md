# Dating Service

## Project Summary

| Item                      | Description                                                                                                          |
|---------------------------|----------------------------------------------------------------------------------------------------------------------|
| Golang Version            | [1.20](https://golang.org/doc/go1.20)                                                                                |
| Object Storage            | [Cloudinary](https://cloudinary.com)                                                                                 |
| moq                       | [mockgen](https://github.com/golang/mock)                                                                            |
| Linter                    | [GolangCI-Lint](https://github.com/golangci/golangci-lint)                                                           |
| Testing                   | [testing](https://golang.org/pkg/testing) and [testify/assert](https://godoc.org/github.com/stretchr/testify/assert) |
| Application Architecture  | [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)                   |
| Directory Structure       | [Standard Go Project Layout](https://github.com/golang-standards/project-layout)                                     |
| CI (Lint, Test, Generate) | [GitHubActions](https://github.com/features/actions)                                                                 |
| Logger                    | [logrus](https://github.com/sirupsen/logrus)                                                                         |
| Database                  | [postgres](https://www.postgresql.org/)                                                                              |
| Database Migration        | [goose](https://github.com/pressly/goose)                                                                            |

## Sequence Diagram

![Sequence-Diagram](https://static.swimlanes.io/efeb3293a28ddc94d333c90b03d12476.png)

## Entity Relationship Diagram

<img src="https://github.com/yonisaka/dating-service/blob/main/docs/erd.png?raw=true"/>

## Installation

### 1. Set Up Golang Development Environment

See the following page to download and install Golang.

https://go.dev/doc/install

### 2. Install Development Utility Tools

You can install all tools for development and deployment for this service by running:

```sh
$ go mod download
```

```sh
$ make install
```

### 3. Setup Environment Variables

```sh
$ cp .env.example .env
```
**_note: for cloudinary env, will be provided by separate file_**

### 4. Run Database Migration

```sh
$ go run main.go db:migrate up
```

### 5. Run Application

```sh
$ go run main.go http:start
```
or
```sh
$ make http
```

### 6. Run Test

```sh
$ make test
```

### 7. Run Lint

```sh
$ make lint
```

## Project Structure

This project follow https://github.com/golang-standards/project-layout

However, for have a clear direction when working in this project, here are some small guide about each directory:

* [api](api): contains Protobuf files, generated protobuf, swagger, etc.
* [build](build): Docker file for the service, migration, etc.
* [cmd](cmd): main Go file for running the service, producer, consumer, etc.
* [config](config): configuration file for the service.
* [database](database): database migration files.
* [development](development): file to support development like docker-compose.
* [docs](docs): file about project documentations such as diagram, sequence diagram, etc.
* [internal](internal): internal code that can't be shared.
    * [internal/adapters/httphandler](internal/adapters/httphandler): adapter layer that serve into gRPC service.
    * [internal/di](internal/di): dependencies injection for connecting each layer.
    * [internal/middleware](internal/middleware): middleware for HTTP server.
    * [internal/consts](internal/consts): constant value that can be shared.
    * [internal/presentations](internal/presentations): presentation that struct for request and response.
    * [internal/infrastructure](internal/infrastructure): infrastructure layer that connect to database, object storage, etc.
    * [internal/entities](internal/entities): entity layer that struct for database table.
    * [internal/usecases](internal/usecases): business logic that connect to repository layer, HTTP client, etc.
* [pkg](pkg): package code that can be shared.
* [scripts](scripts): shell script, go script to help build or testing something.

## GitHub Actions CI

<img src="https://github.com/yonisaka/dating-service/blob/main/docs/ci.jpg?raw=true"/>

This project has GitHub Actions CI to do some automation such as:

* [lint](.github/workflows/lint.yml): check the code style.
* [test](.github/workflows/test.yml): run unit testing and uploaded code coverage artifact.
* [deployment](.github/workflows/deployment.yml): deploy to GCE instance, setup only because not provided by the client.
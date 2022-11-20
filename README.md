# Crox & TSurvey

## Projects Structure

please also read https://github.com/golang-standards/project-layout

Following is the proposed project structure when Telkomsel Engineer team wants to create Golang microservice /
application

```
<root>
  ├── bin
  |   ├── <service_name_executable>
  |   └── ...
  |
  ├── cmd
  |   ├── <service_name>_main.go
  |   └── ...
  |
  ├── deploy
  |   └── <service_name>
  |       ├── deployment.yml
  |       ├── Dockerfile
  |       └── service.yml
  |
  ├── docs
  |
  ├── domain
  │   ├── entity
  │   │   └── <domain_name>.go
  │   ├── valueobject
  │   │   └── <enum_name>.go
  │   ├── repository
  │   │   └── <interface_repository_name>.go
  │   ├── usecase
  │   │   └── <interface_usecase_name>.go
  │   ├── service
  │   │   └── <interface_service_name>.go
  |
  ├── internal
  |   |
  │   ├── mocks
  │   │   ├── <usecase_implementation>Mock.go
  │   │   ├── <service_implementation>Mock.go
  │   │   └── <repo_implementation>Mock.go
  │   ├── config
  │   |   ├── consul-template.config
  │   |   └── <service_name>.ctmpl
  |   |
  │   ├── delivery
  │   │   └── http
  |   |       ├── <domain_name>_init.go
  |   |       ├── <domain_name>_init_test.go
  |   |       ├── <domain_name>_handler.go
  |   |       └── <domain_name>_handler_test.go
  |   |
  |   ├── repository
  |   |   └── <repo_implementation>.go
  |   |
  |   └── usecase
  |       ├── <usecase_name>.go
  |       └── <usecase_name>_test.go
  |
  ├── vendor
  ├── README.txt
  └── MAKEFILE
```

| No   | Folder Name                                  | Purpose                                                                                                      |
| ---- |----------------------------------------------| ------------------------------------------------------------------------------------------------------------ |
| 1.   | `bin`                                        | Golang executable file per service name                                                                      |
| 2.   | `cmd`                                        | Golang main program per service name                                                                         |
| 3.   | `deploy`                                     | Configuration file for deployment per service name                                                           |
| 4.   | `docs`                                       | Repository documentation file                                                                                |
| 5.   | `domain/*/<domain_name>.go`                  | Golang file that holds definiton for data model, usecase interface, and repository interface for each domain |
| 6a.  | `internal/config`                            | Functions, configs used by microsevices                                                                      |
| 6b.  | `internal/delivery/http`                     | As per port adapter pattern, this the port definition that interfacing with outside world                    |
| 6c.  | `internal/repository/<repo_implmentation>.go` | Encapsulated Implementation of Repository Interface                                                          |
| 6d.  | `internal/usecase`                           | Golang file that holds business logic implementation                                                         |
| 6d.1 | `internal/mocks`                             | Mock file                                                                                                    |
| 7.   | `vendor`                                     | Golang external package dependencies                                                                         |
| 8.   | `README.md`                                  | Manual handbook to use the service                                                                           |
| 9.   | `MAKEFILE`                                   | Script definition to build, test, and deploy the application                                                 |

## Prerequisites

1. golang (v1.16.4 or above)
2. postgreSQL database
3. MongoDB database
4. Protobuf Compiler | https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3 (v3.17.3 or above)
5. Buf | https://docs.buf.build/installation
6. Golang Migrate | https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
7. Mockery | https://github.com/vektra/mockery (v2.10.6 or above)

## Installation

1. copy environment -> `cp env.example .env`
2. sync go modules -> `go mod tidy`
3. download required modules -> `go mod download`
4. download GRPC ecosystem modules -> \
   `go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway`\
   `go get google.golang.org/protobuf/cmd/protoc-gen-go`\
   `go get google.golang.org/grpc/cmd/protoc-gen-go-grpc`
5. install GRPC ecosystem modules -> \
   `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway`\
   `go install google.golang.org/protobuf/cmd/protoc-gen-go`\
   `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`

## Services

#### Migration Database

1. download and install golang-migrate -> https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
2. migrate up service (eg. service auth) -> \
   `migrate -source "file://services/auth/migrations/pgsql" -database "postgres://username:pwd@localhost:5432/your_db?sslmode=disable&search_path=auth" up`
3. create new migration (if needed) (eg. service auth) -> \
   `migrate create -ext sql -dir services/auth/migrations/pgsql create_my_table`

#### GRPC Api Gateway

1. run app -> go run services/api/main.go

#### Service Auth

1. run app -> go run services/auth/main.go
2. test -> go test ./.../auth

## Mock

#### Prerequisites

1. Prerequisites -> https://github.com/vektra/mockery

#### Generate Mock

1. go to your service directory (eg. auth) -> \
   `cd services/auth`
2. run generate command -> \
   `mockery --all --case=snake`

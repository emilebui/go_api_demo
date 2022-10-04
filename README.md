# GO API DEMO

## Introduction

- This is an demo of creating go api using grpc and grpc-gateway
  - This demo include:
    - APIs to CRUD repositories
    - APIs to scan repo and get results
    - API to CRUD scanning rules
- This provide both REST API and gRPC API in one code base
- You can use REST API at port 8082 and gRPC API at port 50051
- Tech stack: grpc, grpc-gateway, gorm, sqlite, ...

## How to Install/Run

### Installation
- Run: `go mod download`

### Run the app
- Run: `go run main.go`

### Test
- Run: `go test ./... -cover`

### Docker
- Build: `docker build -t go_api:dev .`
- Build and Run: `docker-compose up -d`

### API Documentations
- You can view `./swagger.yaml` or check this link:
  - https://app.swaggerhub.com/apis-docs/emilebui/go_demo/1.0.0
- For gRPC API, checkout [**proto/demo.proto**](./proto/demo.proto) to see the API contract

### Problems when running/installation
- If you encounter some problems with sqlite during running or installation process
  - Remove previous installed gcc package (if you use minGW)
  - Install this library then restart your machine: https://jmeubank.github.io/tdm-gcc/
  - If the issue isn't resolved, run the app in docker
    - Run `docker-compose up -d`

## Project Structure

- `./main.go`: to run the entire app
- `service`: contains all the logics to run the app
- `proto`: contains API contracts (You need to know about gRPC proto-files)
- `proto_gen`: code that is generated from proto-files (API contract)
- `models`: contains SQL Schema (ORM structure using GORM)
- `load_config`: contains modules to load configs
- `./config.yaml`: overall config of the app
- `./error_messages.json`: config for error messages format
- `helpers`: contains all the helper functions to write code
  - Include: error handlers, path finders, ...
- `client_test`: contain code to test grpc API

## Features
- **gRPC API and REST API**
  - provide both REST API and gRPC API in one code base
- **Config Management**
  - Load config from `config.yaml`
  - Provide singleton config to entire app
  - Can be overwritten when test
- **ORM SQL Database**
  - Creating and Migrating sql database using ORM
  - Provide singleton database interface to entire app
  - Can be overwritten when test
- **Input Validation**
  - Provide input validation when creating new repo
    - repo name cannot contain special characters
- **File extensions White List**
  - Provide a config to whitelist some file extensions while scanning source code
    - E.g: file with `.png` will be skipped during scan
  - Check `./config.yaml` to see the current whitelist
- **Docker Images**
  - You can build and run the app with docker images
  - check `Dockerfile` and `docker-compose.yml`
- **Error Message Format**
  - Provide a config to define error message format
  - Provide error message interface to return error message format from the config

## Working with source code

### Prerequisite

- Understanding of grpc and grpc-gatway, documents can be found here:
  - https://grpc.io/docs/
  - https://github.com/grpc-ecosystem/grpc-gateway
- Install these packages:
  - https://github.com/grpc-ecosystem/grpc-gateway
  - buf.build
  - https://github.com/envoyproxy/protoc-gen-validate

### GRPC Code Gen Cmd
- `cd proto; buf generate`

### ORM Data Structure
- ORM data structure can be found in [**demo_model.go**](./models/demo_model.go)
- Structure Details:
  - Repo: Table for repository
  - Rule: Table for scanning rule
    - `Rule.StringCompare`: vulnerable keyword to check while scanning
  - Result: Table for scanning result
  - Vulnerability: A vulnerability that is found while scanning
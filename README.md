# GO API

## GRPC Gen Code Cmd

run `cd proto; buf generate`

## Run command:

- Server:
    - `go run main.go`
- Client:
    - client_test: `go run client_test/client.go`
    - benchmark:
        - grpc_gateway: `go run client_test/benchmark/grpc_gateway/main.go`
        - pure_grpc: `go run client_test/benchmark/pure_grpc/main.go`
        - rest: `go run client_test/benchmark/rest/main.go`
    
## How to Test

- Run Server
- Then run client_test or benchmark
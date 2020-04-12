# Client retry

Demonstrates how to use client fault tolerance to retry

-[client.go] (client.go) client
-[FailureServer.go] (failureServer.go) error server
-[SuccessServer.go] (successServer.go) Successful server

## Run

Run FailureServer.go The server always returns an error

```bash
go run FailureServer.go
```

Open a new window and run SuccessServer.go The server returns success

```bash
go run FailureServer.go
```

Open a new window and run the client

```bash
go run client.go
```

See the logs printed by the client and each server

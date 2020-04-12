# Use JSON client

Run rpc service

```bash
cd ../rpc
go run rpc.go
2019/08/16 23:29:25 Transport [http] Listening on [::]: 62397
2019/08/16 23:29:25 Broker [http] Connected to [::]: 62398
2019/08/16 23:29:25 Registry [mdns] Registering node: go.micro.rpc.example-ff010ffb-262a-486e-b1b8-7609a4705a86
```

Run the client

```bash
go run main.go
rsp: RPC Call received your request John
```
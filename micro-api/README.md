# Micro API

This article describes how to use Micro's API gateway.

## Overview

The API is essentially a service gateway, which has the capabilities of dynamic routing and service discovery, and maps external requests to local microservices in HTTP to provide external services.

Through service discovery, in conjunction with the built-in namespace rules, the API can map request url resolution to service interfaces that match the namespace rules.

In the Micro system, the service will have its own namespace, and the default namespace of the API is go.micro.api. Generally, microservices deployed afterwards to provide specific interfaces, we will set their namespaces to go.micro.api.example in the Micro style by default, example is the specific service name, if you need to change this value, Specify --namespace = specify name command at startup.

## Handler

There are currently five processing methods for Micro API. We will talk about it below. We can set **API** to the specified type according to our needs.

| - | Type | Description
----|----|----
1 | rpc | Forward the request to the go-micro application via RPC, only receive GET and POST requests, GET forwards `RawQuery`, POST forwards `Body`
2 | api | Similar to rpc, but it will encapsulate the complete http header and send it down, without limiting the request method
3 | http或proxy | Using **API** as a reverse proxy is equivalent to deploying an ordinary web application after **API** and letting the outside world call the web service like an API interface
4 | web | similar as httpwebsocket
5 | event | Proxy event service type request
6 | meta | Default value, metadata, use one of the above processors through the `Endpoint` configuration in the code, the default RPC

- Focus on the difference between the two types of rpc and api. The difference is that rpc does not encapsulate the request header information and api will.
- meta, there is no such mode, just an extended use of api, rpc, proxy, web and other modes.
- `Endpoint` will be matched first in the route, so using rpc or api mode can also use this method to define flexible routes.
- The current version (V1) cannot support multiple handlers running concurrently

## Request mapping

### RPC/API type

Micro has a mechanism for mapping http request paths to services, and the mapping rules can be introduced through the following table

http route    |    backend service    |    port method
----    |    ----    |    ----
/foo/bar    |    go.micro.api.foo    |    Foo.Bar
/foo/bar/baz    |    go.micro.api.foo    |    Bar.Baz
/foo/bar/baz/cat    |    go.micro.api.foo.bar    |    Baz.Cat

The default is **go.micro.api**，as we mentioned above, we can `--namespace` to set namespace。

We can use the version on the path name and map it to the service name

Request path   |    Path name  |    request method
----    |    ----    |    ----
/foo/bar    |    go.micro.api.foo    |    Foo.Bar
/v1/foo/bar    |    go.micro.api.v1.foo    |    Foo.Bar
/v1/foo/bar/baz    |    go.micro.api.v1.foo    |    Bar.Baz
/v2/foo/bar    |    go.micro.api.v2.foo    |    Foo.Bar
/v2/foo/bar/baz    |    go.micro.api.v2.foo    |    Bar.Baz

As can be seen from the above mapping rules, in the **RPC/API** mode, the two parameters after the path will be combined into the Golang public method path name, and the rest will be prefixed with the namespace to form the service name. such as:

`/v1/foo/bar/baz`, in which the first letter of `bar/baz` is converted to the path of the `Bar.Baz` method; the remaining `/v1/foo/` is appended with the namespace prefix `go.micro.api` composition
`go.micro.api.v1.foo`.

### Proxy type

If we start **API** and pass the instruction `--handler = http`, then **API** will reverse the proxy request to the background service with API namespace.

For example:

Request path    |    Service    |    Backend service path
---    |    ---    |    ---
/greeter    |    go.micro.api.greeter    |    /greeter
/greeter/:name    |    go.micro.api.greeter    |    /greeter/:name

### Event type

When starting **API**, the instruction `--handler=event` is passed, then **API** will reverse the proxy request to the background event consumption service with API namespace.

For example (namespace is set to go.micro.evt):

Request path    |    Service    |    Method
---    |    ---    |    ---
/user/login    |    go.micro.evt.user    |    All public methods of the listener object (new (Event) in the example), and the method must have ctx and event parameters

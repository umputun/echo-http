# echo-http - Echo http service

Responds with json-formatted echo of the incoming request and with a predefined message.

## installation

- The binary can be installed directly with go install: `go install github.com/umputun/echo-http@latest`
- Can be downloaded from [releases](https://github.com/umputun/echo-http/releases)
- For MacOS user can be installed with brew: `brew install umputun/apps/echo-http`
- For docker there is a multi-arch docker container `ghcr.io/umputun/echo-http`


## usage

Send any http request to the server, and it will respond with a json-formatted echo with all the things it knows about the request.

```sh
http https://echo.umputun.com/something
```

```json
{
    "headers": {
        "Accept": "*/*",
        "Accept-Encoding": "gzip",
        "User-Agent": "HTTPie/2.4.0",
        "X-Forwarded-For": "12.12.12.12",
        "X-Forwarded-Host": "172.29.0.2:8080",
        "X-Origin-Host": "echo.umputun.com",
        "X-Real-Ip": "12.12.12.12"
    },
    "host": "172.29.0.2:8080",
    "message": "echo echo 123",
    "remote_addr": "172.29.0.3:37432",
    "request": "GET /something"
}
```

Alternatively, it can be run within cloudflare worker. In this case, the server will respond with the same echo. See [cf-worker.js](https://github.com/umputun/echo-http/blob/master/cf-worker.js) for details. The worker is available at `https://echo.umputun.dev/`.

```sh
http https://echo.umputun.dev/aaaa
```

```json
{
  "headers": {
    "accept": "application/json, */*;q=0.5",
    "accept-encoding": "gzip, br",
    "connection": "Keep-Alive",
    "content-length": "15",
    "content-type": "application/json",
    "host": "echo.umputun.dev",
    "user-agent": "HTTPie/3.2.2",
    "x-forwarded-proto": "https"
  },
  "host": "echo.umputun.dev",
  "message": "echo",
  "remote_addr": "50.123.44.200",
  "request": "POST /aaaa"
}
```

## Application options

```
  -l, --listen=  listen on host:port (default: 0.0.0.0:8080) [$LISTEN]
  -m, --message= response message (default: echo) [$MESSAGE]

Help Options:
  -h, --help     Show this help message
```

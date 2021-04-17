# echo-http - Echo http service

Responds with json-formatted echo of the incoming request and with a predefined message.

Can be install directly (`go get github.com/umputun/echo-http`) or as a multi-arch docker container `ghcr.io/umputun/echo-http`

`http https://echo.umputun.com/something`

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

```
Application Options:
  -l, --listen=  listen on host:port (default: 0.0.0.0:8080) [$LISTEN]
  -m, --message= response message (default: echo) [$MESSAGE]

Help Options:
  -h, --help     Show this help message
```

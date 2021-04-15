# echo-http - Echo http service

Responds with json-formatted echo of the incoming request and with a predefined message.


`http http 127.0.0.1:8080/some/test`

```json
{
    "headers": [
        "Accept-Encoding:gzip, deflate",
        "Accept:*/*",
        "Connection:keep-alive",
        "User-Agent:HTTPie/2.4.0"
    ],
    "host": "127.0.0.1:8080",
    "message": "echo",
    "remote_addr": "127.0.0.1:49821",
    "request": "GET /some/test"
}

```

```
Application Options:
  -l, --listen=  listen on host:port (default: 0.0.0.0:8080) [$LISTEN]
  -m, --message= response message (default: echo) [$MESSAGE]
      --dbg      debug mode [$DEBUG]

Help Options:
  -h, --help     Show this help message
```
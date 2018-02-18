# RedisClusterBenchmark

## Guide

Install
```shell
go get -u github.com/spf13/cobra/cobra
go get -u github.com/go-redis/redis
```

Build
```shell
go build main.go
```

Run
```shell
./main
```

```shell

Usage:
  redisbenchmark test [flags]

Flags:
  -C, --cluster          cluster mode on/off
  -c, --concurrent int   number of concurrent clients (default 500)
  -h, --help             help for test
  -i, --ip string        ip of redis server (default "127.0.0.1")
  -p, --port string      port of redis server (default "6379")
  -n, --request int      number of requests (default 200)
  -s, --status string    [start|print] start test now! (required)
  -v, --verbose          verbose output

```

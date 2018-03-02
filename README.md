# RedisClusterBenchmark

[![wercker status](https://app.wercker.com/status/48d0ef59213a4e0901294db0857ee5ac/s/master "wercker status")](https://app.wercker.com/project/byKey/48d0ef59213a4e0901294db0857ee5ac)
[![License](https://img.shields.io/badge/LICENSE-Apache2.0-ff69b4.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)

## Install
```shell
go get github.com/imzye/RedisClusterBenchmark
```

## Build
```shell
cd $GOPATH/src/github.com/imzye/RedisClusterBenchmark
go build redisbenchmark.go
```

## Run
```shell
./redisbenchmark
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

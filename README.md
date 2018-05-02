# rate-limiter

Simple rate-limiter middleware written in go.
I made a break from golang for 8 months - so this is a comeback :)

## Requirements

- Install golang: <https://golang.org/doc/install>
- Setup GOPATH: <https://github.com/golang/go/wiki/SettingGOPATH>
- Install go dep: <https://github.com/golang/dep>

## Buiding

- Clone this repository
- `dep ensure`
- `cd cmd`
- `go build -o limiter`

There is a Makefile available with `make all` option but notice I am tempering with GOPATH (sandboxing building) if you did not go through hell with GOPATH/GOROOT please do not use it - just build it manually. If you select this approach the executable will be created in `./.gopath~/bin/` meaning just run `./.gopath~/bin/limiter` binary.

## Running

I have implemented two strategies for rate limiting one in-memory and one in redis.
The advantages of redis should be obvious (horizontal scalling)
Please check the ./cfg/limiter.yml file to set the configuration for your run.

To run in-memory example:

- source `go run cmd/* rlmem --config ./cfg/limiter.yml`
- binary `./.gopath~/bin/limiter rlmem --config ./cfg/limiter.yml`

```shell
➜ curl localhost:8080/greet/me -v
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /greet/me HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< X-Ratelimit-Limit: 5
< X-Ratelimit-Remaining: 5
< X-Ratelimit-Reset: 2018-05-02 20:03:47.693777422 +0200 CEST
< Date: Wed, 02 May 2018 18:03:41 GMT
< Content-Length: 93
<
* Connection #0 to host localhost left intact
{"name":"limiter","description":"Simple rate limiter demonstration.","uptime":"2.997741413s"}%
```

To run redis example:

- source `go run cmd/*  --config ./cfg/limiter.yml`
- binary `./.gopath~/bin/limiter --config ./cfg/limiter.yml`

```shell
➜ curl localhost:8080/greet/me/redis -v
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /greet/me/redis HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< X-Ratelimit-Limit: 5
< X-Ratelimit-Remaining: 5
< X-Ratelimit-Reset: 2018-05-02 20:06:13.686774342 +0200 CEST
< Date: Wed, 02 May 2018 18:06:07 GMT
< Content-Length: 94
<
* Connection #0 to host localhost left intact
{"name":"limiter","description":"Simple rate limiter demonstration.","uptime":"18.883323352s"}%
```
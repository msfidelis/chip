<p><img alt="Version" src="./.github/images/logo.png" /></p>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-0-blue.svg?cacheSeconds=2592000" />
  <a href="/README.md" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
  <a href="LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://twitter.com/fidelissauro" target="_blank">
    <img alt="Twitter: fidelissauro" src="https://img.shields.io/twitter/follow/fidelissauro.svg?style=social" />
  </a>
</p>

> Docker container to make runtime and platforms tests easy :whale: :package: :rocket:

### 🏠 [Homepage](/)

### ✨ [Demo](/)

#### Available images for test 

* [v1]()
* [v2]()
* [latest]()

```sh
docker run -it -p 8080:8080 fidelissauro/chip:v1 
```


## Compile Go Binary (Without Docker)

```sh
go get -u
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```

## Run tests

```sh
go test
```

## Setup Development Environment

* Development environment uses [air project](https://github.com/cosmtrek/air) to execute live reload on `.go` files.

```sh
docker-compose up --force-recreate
```

## Build image

```sh
docker build -it chip
```

## Usage

```sh
docker run -it -p 8080:8080 msfidelis/chip:v1
```

## Swagger

check on http://localhost:8080/swagger/index.html


# Endpoints


## Healthcheck Endpoint

Common healthcheck, dummy mock

```sh
curl 0.0.0.0:8080/healthcheck -i

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 30 May 2020 22:43:11 GMT
Content-Length: 14

{"status":200}
```

## Healthcheck Endpoint (Error)

Simulate error on Healthcheck

```sh
curl 0.0.0.0:8080/healthcheck/error -i

HTTP/1.1 503 Service Unavailable
Content-Type: application/json; charset=utf-8
Date: Sat, 30 May 2020 22:44:35 GMT
Content-Length: 14

{"status":503}
```

## Healthcheck with Fault Injection (Random Mode)

Use this for fault injection, circuit breaker, self healing tests on your readiness probe

```sh
while true; do curl 0.0.0.0:8080/healthcheck/fault; echo;  done

{"status":503}
{"status":503}
{"status":200}
{"status":503}
{"status":503}
{"status":200}
{"status":503}
{"status":200}
{"status":200}
{"status":503}
{"status":503}
{"status":200}
{"status":200}
{"status":200}
{"status":200}
{"status":503}
{"status":200}
{"status":200}
{"status":200}
```

## Healthcheck with Fault Injection (Soft Mode)

Cause ocasional failure in your probe

```sh
while true; do curl 0.0.0.0:8080/healthcheck/fault/soft; echo;  done

{"status":200}
{"status":200}
{"status":200}
{"status":200}
{"status":200}
{"status":200}
{"status":200}
{"status":200}
{"status":200}
{"status":200}
{"status":503}
{"status":200}
{"status":200}
{"status":503}
```

## Liveness Probe OK

For Liveness tests

```sh
curl 0.0.0.0:8080/liveness -i
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 08 May 2021 14:00:05 GMT
Content-Length: 17

{"status":"Live"}
```

## Liveness Probe with intentional failure

```sh 
curl 0.0.0.0:8080/liveness/error -i
HTTP/1.1 503 Service Unavailable
Content-Type: application/json; charset=utf-8
Date: Sat, 08 May 2021 14:01:08 GMT
Content-Length: 17

{"status":"Dead"}
```

## Readiness

For readiness tests

```sh 
curl 0.0.0.0:8080/readiness -i
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 08 May 2021 14:02:02 GMT
Content-Length: 18

{"status":"Ready"}
```

* You can set the environment variable  `READINESS_PROBE_MOCK_TIME_IN_SECONDS` to customize your readiness probe in seconds for testing. Default is 5 seconds. 

## Liveness Probe with intentional failure

```sh
curl 0.0.0.0:8080/readiness/error -i
HTTP/1.1 503 Service Unavailable
Content-Type: application/json; charset=utf-8
Date: Sat, 08 May 2021 14:02:37 GMT
Content-Length: 22

{"status":"Not Ready"}
```

## Version

This endpoint return different values in accord to tag version, v1, v2, v1-blue, v1-green, v2-blue and v2-green. Ideal to tests deployment scenarios behavior, like rollout, canary, blue / green etc.

This variable can be customized using environment variable called `VERSION``

```sh
export VERSION=v3
```

```sh
curl 0.0.0.0:8080/version -i

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 31 May 2020 03:38:21 GMT
Content-Length: 16

{"version":"v1"}
```


## System Info 
Retrieve some system info. Use this to test memory, cpu limits and isolation. Host name for load balancing tests and etc.

```sh
curl 0.0.0.0:8080/system -i

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 31 May 2020 03:33:12 GMT
Content-Length: 76

{"hostname":"21672316d98d","cpus":2,"os":"","hypervisor":"bhyve","memory":0}
```

## Dump Environment Variables

```sh
curl http://0.0.0.0:8080/system/environment -i  

[
  "HOSTNAME=78339a8484d4",
  "HOME=/root",
  "ENVIRONMENT=dev",
  "PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
  "GOPATH=/go",
  "PWD=/go/src/chip",
  "GOLANG_VERSION=1.13.11"
]
```

## Reflection (In Progress)

Use this endpoint to retrieve request headers, body, querystrings, cookies, etc. Ideal to tests API Gateway, CDN, Proxys, Load Balancers transformations on request. Available for all HTTP methods

```sh
curl -X GET 0.0.0.0:8080/reflect -i

{
  "method": "GET",
  "params": "",
  "headers": {
    "Accept": [
      "*/*"
    ],
    "User-Agent": [
      "curl/7.64.1"
    ]
  },
  "cookies": [],
  "body": "",
  "path": "/reflection"
}
```

## Proxy Request 

Use this endpoint to proxy HTTP requests betweet chip and another endpoint 


```sh
curl --location --request POST 'http://0.0.0.0:8080/proxy' \                                                     ─╯
--header 'Content-Type: application/json' \
--data-raw '{
    "method": "GET",
    "host": "https://google.com.br/",
    "path": "/whoami",
    "headers": [
        {
            "name": "content-type",
            "value": "application/json"
        },
        {
            "name": "foo",
            "value": "bar"
        }
    ],
    "body": "{\"vai\": \"sim\"}"
}' 
```

```sh
curl -X POST "0.0.0.0:8080/reflection?id=1" -H "Header-Foo: Bar" -d '{"foo":"bar"}' | jq .

{
  "method": "POST",
  "params": "id=1",
  "headers": {
    "Accept": [
      "*/*"
    ],
    "Content-Length": [
      "13"
    ],
    "Content-Type": [
      "application/x-www-form-urlencoded"
    ],
    "Header-Foo": [
      "Bar"
    ],
    "User-Agent": [
      "curl/7.64.1"
    ]
  },
  "cookies": [],
  "body": "{\"foo\":\"bar\"}",
  "path": "/reflection"
}
```


## Load

Use this endpoint to consume some CPU resources. Ideal to test auto scale policies, isolation, monitoring and alerts behaviors.

> Danger

```sh 
curl -X GET 0.0.0.0:8080/burn/cpu -i

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 02 Jun 2020 04:42:24 GMT
Content-Length: 20

{"status":"On Fire"}
```

```sh 
curl -X GET 0.0.0.0:8080/burn/ram -i

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 02 Jun 2020 04:42:24 GMT
Content-Length: 20

{"status":"On Fire"}
```

## Check ACL and Connection to Ports

Check connection between container environment / namespace and services and another applications


```sh
curl -X GET 0.0.0.0:8080/ping/google.com/80 -i

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 20 Jul 2020 16:04:02 GMT
Content-Length: 71

{"host":"google.com","port":"80","protocol":"tcp","status":"connected"}
```

## Logging Events 

Sent a log of logs to stdout. Sent querystring `events` to customize the number of logs; Default `1000`.

```sh
curl -X GET "0.0.0.0:8080/logging?events=2000" -i 

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 28 Aug 2021 14:50:41 GMT
Content-Length: 61

{"status":200,"message":"2000 logging events sent to stdout"}
```

## Filesystem

List some directory contents 

```bash
curl -X POST "0.0.0.0:8080/filesystem/ls" -i -d '{"path": "./"}'

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 10 Jan 2025 11:52:26 GMT
Content-Length: 86

{"path":"./","files":[".dockerenv","dev","etc","lib","main","proc","sys","tmp","var"]}
```

Write file - base64 content

```bash
curl -X POST "0.0.0.0:8080/filesystem/write" -i -d '{"path": "./test", "content": "V3JpdGUgVGVzdAo="}';

HTTP/1.1 202 Accepted
Content-Type: application/json; charset=utf-8
Date: Fri, 10 Jan 2025 11:58:13 GMT
Content-Length: 42

{"message":"file written","path":"./test"}
```

Read files on filesystem 

```bash
curl -X POST "0.0.0.0:8080/filesystem/cat" -i -d '{"path": "./test"}';

HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Fri, 10 Jan 2025 11:58:44 GMT
Content-Length: 11

Write Test
```

Delete files on filesystem 

```bash
curl -X DELETE "0.0.0.0:8080/filesystem/delete" -i -d '{"path": "./test"}';

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 10 Jan 2025 11:59:01 GMT
Content-Length: 26

{"message":"file deleted"}
```

## Whoami? 

```sh
curl -X GET 0.0.0.0:8080/whoami -i 

HTTP/1.1 418 I'm a teapot
Content-Type: text/plain; charset=utf-8
Date: Sun, 23 May 2021 00:53:46 GMT
Content-Length: 85


             ;,'
     _o_    ;:;'
 ,-.'---`.__ ;
((j`=====',-'
 `-\     /
    `-=-'
```

## Author

👤 **Matheus Fidelis**

* Website: https://raj.ninja
* Twitter: [@fidelissauro](https://twitter.com/fidelissauro)
* Github: [@msfidelis](https://github.com/msfidelis)
* LinkedIn: [@msfidelis](https://linkedin.com/in/msfidelis)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](/issues). 

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2020 [Matheus Fidelis](https://github.com/msfidelis).<br />
This project is [MIT](LICENSE) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_

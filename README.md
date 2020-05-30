<h1 align="center">Welcome to Chip - Cloud Native &#34;Smart&#34; Dummy Mock 👋</h1>
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

## Install

```sh
go get -u
```

## Run tests

```sh
go test
```

## Usage

```sh
docker run -it -p 8080:8080 msfidelis/chip:v1
```

## Endpoints

### Healthcheck Endpoint

```sh
curl 0.0.0.0:8080/healthcheck -i

Common healthcheck, dummy mock

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 30 May 2020 22:43:11 GMT
Content-Length: 14

{"status":200}
```

### Healthcheck Endpoint (Error)

Simulate error on Healthcheck

```sh
curl 0.0.0.0:8080/healthcheck/error -i

HTTP/1.1 503 Service Unavailable
Content-Type: application/json; charset=utf-8
Date: Sat, 30 May 2020 22:44:35 GMT
Content-Length: 14

{"status":503}
```

### Healthcheck with Fault Injection

Use this for fault injection, circuit breaker, self healing tests

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


### System Info 
Retrieve some system info. Use this to test memory, cpu limits and isolation. Host name for load balancing tests and etc. 

```sh
curl 0.0.0.0:8080/system -i
```

### Reflection

Use this endpoint to retrieve request headers, body, querystrings, cookies, etc. Ideal to tests API Gateway, CDN, Proxys, Load Balancers transformations on request

```sh
curl -X GET 0.0.0.0:8080/reflect -i
```

```sh
curl -X POST 0.0.0.0:8080/reflect -i
```

```sh
curl -X PUT 0.0.0.0:8080/reflect -i
```

```sh
curl -X PATCH 0.0.0.0:8080/reflect -i
```

```sh
curl -X DELETE 0.0.0.0:8080/reflect -i
```

### Load

Use this endpoint to consume some CPU / Memory resources. Ideal to test auto scale policies, monitoring and alerts behaviors. 

```sh 
curl -X GET 0.0.0.0:8080/load -i
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
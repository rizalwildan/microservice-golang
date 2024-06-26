# Golang Microservice

> Notes: This microservice is not ready for production; it is intended solely for demonstration purposes.

Slide: [Google Slide](https://docs.google.com/presentation/d/1AQniCBgoW7SPgcQOVffojbuBElu21Muvl0SmFDTTWTc/edit?usp=sharing)

## Requirement
* Go version 1.20
* GoFiber
* Traefik
* Docker

## Architecture

![image](arch.png)

## How to Run

### Using Docker

```bash
$ docker compose up -d
```

### Local Machine

Change directory to each root projects

```bash
$ cd {service_name}
```

```bash
$ go mod tidy
```

```bash
$ make run
```

## API Docs

API Documentation use swagger, you can access from:

```bash
http://localhost:8081/v1/users/docs
http://localhost:8081/v1/products/docs
http://localhost:8081/v1/merchants/docs
```

or you can import postman collection inside `postman/` folder

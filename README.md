# go-programming-projects

## Introduction

This repository contains resources for learning Go Programming through building projects.

To run the database services:

```
[x86 machines]
docker-compose -p go-programming-projects -f deployment/docker-compose/docker-compose-x86.yml --profile databases up -d
```

Login to pgadmin4
```
localhost:5050
```

Create connection to existing postgres database inside docker network with docker service name
```
postgres
```
## Run Commands:
Ensure the go version installed is

```
1.21.3
```

Change directory into module directory to init, tidy, build and run go service
```
go mod init github.com/XXX

go mod tidy

go build

go run cmd/<service-name>/main.go
```

lambda-yt-example - Invoke api 

```
aws lambda invoke --function-name lambda-yt-example --cli-binary-format raw-in-base64-out --payload '{\"What is your name?\": \"Jim\",\"How old are you?\": 33}' output.txt
```

go-fiber-crm-basic - Generate swagger json and be displayed by swagger/ui/swagger-initializer.js
```
swagger generate spec -o .\swagger\swagger.json

Navigate to http://localhost:3000/swagger/ui/index.html#/ to view your generated swagger ui documentation
```

## Tutorial Resources:
* **[freeCodeCamp](https://www.freecodecamp.org/news/learn-go-by-building-11-projects/)**
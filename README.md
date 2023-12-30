# go-programming-projects

## Introduction

This repository contains resources for learning Go Programming through building projects.

## Sql Databases

To run the sql-db-stack services:

```
[x86 machines]
docker-compose -p sql-db-stack -f deployment/docker-compose/docker-compose-x86.yml --profile sql-db up -d
```

Login to pgadmin4
```
localhost:5050
Email: pgadmin4@pgadmin.org
Password: admin
```

Create connection to existing postgres database inside docker network with docker service name
```
Name: sql-db-stack
Host: postgres
```

### Nosql Databases
To run the nosql-db-stack services:

```
[x86 machines]
docker-compose -p nosql-db-stack -f deployment/docker-compose/docker-compose-x86.yml --profile nosql-db up -d
```

Login to mongo-express
```
localhost:8081
User: admin
Password: pass
```

Mongo databases are created automatically when there is data inserted

## Run Commands:
Ensure the go version installed is

```
1.21.3
```

Change directory into module directory e.g go-bookstore for the following go commands
```
go mod init github.com/XXX

make build

make run

make test

make coverage

make clean
```

lambda-yt-example - Invoke api 

```
aws lambda invoke --function-name lambda-yt-example --cli-binary-format raw-in-base64-out --payload '{\"What is your name?\": \"Jim\",\"How old are you?\": 33}' output.txt
```

go-fiber-crm-basic - Generate swagger json and be displayed by swagger/ui/swagger-initializer.js
```
Navigate to http://localhost:3000/swagger/ui/index.html#/ to view your generated swagger ui documentation
```

## Tutorial Resources:
* **[freeCodeCamp](https://www.freecodecamp.org/news/learn-go-by-building-11-projects/)**
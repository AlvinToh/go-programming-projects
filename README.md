# Go Programming Projects

## Introduction

This repository provides resources for learning Go through project building.

## Setup

### SQL Databases

1. Run the sql-db-stack services:
    ```bash
    docker-compose -p sql-db-stack -f deployment/docker-compose/docker-compose-x86.yml --profile sql-db up -d
    ```
2. Login to pgadmin4 at `localhost:5050` (Email: `pgadmin4@pgadmin.org`, Password: `admin`)
3. Create connection to existing postgres database inside docker network with docker service name:
    - Name: `sql-db-stack`
    - Host: `postgres`

### NoSQL Databases

1. Run the nosql-db-stack services:
    ```bash
    docker-compose -p nosql-db-stack -f deployment/docker-compose/docker-compose-x86.yml --profile nosql-db up -d
    ```
2. Login to mongo-express at `localhost:8081` (User: `admin`, Password: `pass`)

Note: Mongo databases are created automatically when there is data inserted.

## Usage

Ensure Go version `1.21.3` is installed. 

1. Create a new service module directory:
    ```bash
    go mod init github.com/XXX
    ```
2. Clone `go-bookstore` or `go-fiber-crm-basic` as a template.
3. Update `application.yml` and `application-deploy.yml` under `service-name/resource`.

### Commands

- `make build`
- `make run`
- `make test`
- `make coverage`
- `make clean`
- `make deploy`
- `make destroy`

### Lambda YT Example

Invoke API: 

```bash
aws lambda invoke --function-name lambda-yt-example --cli-binary-format raw-in-base64-out --payload '{\"What is your name?\": \"Jim\",\"How old are you?\": 33}' output.txt
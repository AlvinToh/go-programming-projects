# Go Programming Projects

## Introduction

This repository provides resources for learning Go through project building.

## Setup

### Windows Dependency Setup

1. **Go**
   - Download Go version `1.21.3` from the [official download page](https://golang.org/dl/).
   - Install Go and add `C:\Go\bin` to your PATH:

     ```powershell
     $env:Path += ";C:\Go\bin"
     go version
     ```

2. **Docker**
   - Download Docker Desktop from the [official download page](https://www.docker.com/products/docker-desktop).
   - Install Docker and verify the installation:

     ```powershell
     docker version
     ```

3. **Chocolatey**
   - Install Chocolatey:

     ```powershell
     Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
     ```

4. **Make**
   - Install Make using Chocolatey:

     ```powershell
     choco install make
     ```

5. **Go-Swagger**
   - Install `go-swagger`:

     ```powershell
     go install github.com/go-swagger/go-swagger/cmd/swagger@latest
     $env:Path += ";$(go env GOPATH)\bin"
     swagger --version
     ```

Remember to reopen PowerShell after each installation to ensure changes take effect.

### SQL Databases

1. Run the SQL DB stack services:

    ```bash
    docker-compose -p sql-db-stack -f deployment/docker-compose/docker-compose-x86.yml --profile sql-db up -d
    ```

2. Login to pgAdmin4 at `localhost:5050` (Email: `pgadmin4@pgadmin.org`, Password: `admin`).
3. Create a connection to the existing Postgres database inside the Docker network with the Docker service name:
    - Name: `sql-db-stack`
    - Host: `postgres`

### NoSQL Databases

1. Run the NoSQL DB stack services:

    ```bash
    docker-compose -p nosql-db-stack -f deployment/docker-compose/docker-compose-x86.yml --profile nosql-db up -d
    ```

2. Login to Mongo Express at `localhost:8081` (User: `admin`, Password: `password`).

Note: Mongo databases are created automatically when data is inserted.

## Usage

Ensure Go version `1.21.3` is installed. 

1. Create a new service module directory:

    ```bash
    go mod init github.com/your_username/your_repository
    ```

2. Clone one of the following service directories to use as a Go service template:
    - `go-bookstore`
    - `go-fiber-crm-basic`

3. Update `application.yml` and `application-deploy.yml` under `service-name/resource`.

### Commands

You can use the following commands under `service-name/Makefile`:

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
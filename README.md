# go-programming-projects

## Introduction

This repository contains resources for learning Go Programming through building projects.

To run the database services:

```
[x86 machines]
docker-compose -p go-programming-projects -f deployment/docker-compose/docker-compose-x86.yml --profile databases up -d
```

## Run Commands:
* lambda-yt-example
```
aws lambda invoke --function-name lambda-yt-example --cli-binary-format raw-in-base64-out --payload '{\"What is your name?\": \"Jim\",\"How old are you?\": 33}' output.txt
```

## Tutorial Resources:
* **[freeCodeCamp](https://www.freecodecamp.org/news/learn-go-by-building-11-projects/)**
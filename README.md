# week-17-rest-api

Assignment Week 17 Rest Api

## Description

This is a simple REST API that allows you to create, read, update, and delete users. The API is built using Go, and MongoDB.

## Project environment variables

Environment variables are used to configure the application. The following environment variables are required:

```bash
  LISTEN_ADDR=":XXXX"
  MONGO_DB_NAME="your-db-name"
  MONGO_DB_TEST_NAME="your-db-test-name"
  MONGO_URI="mongodb://..."
  MONGO_TEST_URI="mongodb://..."
```

### Mongodb

Documentation

```bash
https://mongodb.com/docs/drivers/go/current-quickstart
```

Installing mongodb client

```bash
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber

Documentation

```bash
https://gofiber.io
```

Installing gofiber

``` bash
go get github.com/gofiber/fiber/v2
```

## Docker

### Installing mongodb as a Docker container

```bash
docker run --name mongodb -d -p 27017:27017 mongo:latest
```

## Instructions

![image](https://assets.skool.com/f/315feac117b14c21bd46cb1ea5854e69/0ff433a96995412ea326cc698bf062f41706d470069f4bfe9a497c43bb635982)
![image](https://assets.skool.com/f/315feac117b14c21bd46cb1ea5854e69/c6232081b0684be1aabc7cf38079d39df37c88a1e7ce48d9b03496aae617e959)
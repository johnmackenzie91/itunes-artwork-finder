# Itunes Artwork Proxy API

[![CircleCI](https://circleci.com/gh/johnmackenzie91/itunes-artwork-finder/tree/master.svg?style=svg)](https://circleci.com/gh/johnmackenzie91/itunes-artwork-finder/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/johnmackenzie91/itunes-artwork-finder/badge.svg?branch=master)](https://coveralls.io/github/johnmackenzie91/itunes-artwork-finder?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnmackenzie91/itunes-artwork-finder)](https://goreportcard.com/report/github.com/johnmackenzie91/itunes-artwork-finder)

# What is this?
This is a golang api that calls and caches the Itunes artwork api.

## How do I use it?
The recommended way to run this service is to pull from [Dockerhub](https://hub.docker.com/repository/docker/johnmackenzie91/itunes-artwork-proxy-api).
This can be accomplished by;
```shell
docker run -p 8080:80 johnmackenzie91/itunes-artwork-proxy-api:latest
```

You can also pull down the source code from github and run that way.

```shell script
$ git clone https://github.com/johnmackenzie91/itunes-artwork-finder.git
$ cd ./johnmackenzie91/itunes-artwork-finder.git
$ make start logs
```

## Which endpoints are currently provided.
API documentation can be found by visiting http://0.0.0.0:8080/docs once the app is up and running.

# Itunes Artwork Proxy API

[![CircleCI](https://circleci.com/gh/johnmackenzie91/itunes-artwork-finder/tree/master.svg?style=svg&circle-token=<YOUR-TOKEN>)](https://circleci.com/gh/johnmackenzie91/itunes-artwork-finder/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/johnmackenzie91/itunes-artwork-finder/badge.svg?branch=<YOUR-HEAD-BRANCH>)](https://coveralls.io/github/johnmackenzie91/itunes-artwork-finder?branch=master

## What is this?
This is a golang api that calls and caches the amazon artwork api.

## How do I use it?
There are two different ways of using this codebase.
You can either pull the docker image straight from docker hub, or clone this codebase and run `make start logs`.

#### Docker Image
You can pull down and run the docker image via;
```shell script
command for docker
```

#### Clone the repo
You can clone this repo down and run locally with;

```shell script
    git clone ...
    cd johnmackenzie91/itunes-artwork-proxy-api
    make start logs
```
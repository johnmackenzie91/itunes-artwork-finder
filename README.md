# Itunes Artwork Proxy API

[![CircleCI](https://circleci.com/gh/johnmackenzie91/itunes-artwork-finder/tree/master.svg?style=svg&circle-token=<YOUR-TOKEN>)](https://circleci.com/gh/johnmackenzie91/itunes-artwork-finder/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/johnmackenzie91/itunes-artwork-finder/badge.svg?branch=<YOUR-HEAD-BRANCH>)](https://coveralls.io/github/johnmackenzie91/itunes-artwork-finder?branch=master

## What is this?
This is a golang api that calls and caches the Itunes artwork api.

## How do I use it?
The best way to begin using the api is to clone the codebase.

```shell script
git clone https://github.com/johnmackenzie91/itunes-artwork-finder.git
```

Run in docker in one of two ways;

Via Make
```shell script
make start logs
```

or straight via docker
```shell script
docker run -p 5678:80 itunes-proxy
```
# Itunes Artwork Proxy API

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
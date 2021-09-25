# Artwork Package

## What Is This Package used for?
This package contains the functionality to fetch artwork form datasources (itunes only atm).

## How is this done?
This package uses the adapter pattern to allow multiple data sources to conform to a common api;

```go
type Adapter func(ctx context.Context, term, country, entity string) (SearchResponse, error)
```

## Currently Supporter data sources

```go
artwork.Itunes("https://itunes.apple.com/search", logger)
```
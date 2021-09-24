# Itunes Package

## What Is This Package used for?
This package allows us to query the itunes search endpoint for artwork "https://itunes.apple.com/search".

## How is this done?
For configuration please see opts.go file.

```go
	cli, err := itunesCli.New(
        itunesCli.SetDomain(endpoint),
        itunesCli.WithLogger(logger),
    )
    if err != nil {
        return nil, err
    }
	cli.Search(ctx, "some artist - some album", "gb", "album")
```

### Generate test fixtures
Run generate_test_fixtures.sh to update the test fixtures which the unit tests rely on.
Be aware this could cause unit tests to fail. Please fix accordingly.
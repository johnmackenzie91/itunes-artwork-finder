FROM golang:1.16
RUN go get github.com/githubnemo/CompileDaemon
RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /src
COPY . .

RUN go build -gcflags="all=-N -l" -o /app
ENTRYPOINT dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec /app
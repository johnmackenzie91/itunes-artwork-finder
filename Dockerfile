# start a golang base image, version 1.8
FROM golang:1.14 as builder

WORKDIR /go/src

COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build .

FROM alpine
# Copy our static executable.
COPY --from=builder /go/src/itunes-artwork-proxy-api /bin/itunes-artwork-proxy-api

RUN ls -lah /bin/itunes-artwork-proxy-api

EXPOSE 80
CMD ["/bin/itunes-artwork-proxy-api"]

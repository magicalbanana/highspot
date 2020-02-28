FROM golang:1.13.8-alpine3.11 as builder

RUN mkdir -p /go/src/github.com/magicalbanana/highspot/

WORKDIR /go/src/github.com/magicalbanana/highspot/

COPY . .

RUN apk add --update --no-cache alpine-sdk git

RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -mod vendor -v -a -installsuffix cgo -o mixtape-cmd \
    main.go

# actual container
FROM alpine:3.11

RUN apk add --update --no-cache bash ca-certificates

RUN mkdir -p /app

WORKDIR /app

RUN mkdir -p testdata

COPY --from=builder /go/src/github.com/magicalbanana/highspot/mixtape-cmd .

CMD ["./mixtape-cmd"]

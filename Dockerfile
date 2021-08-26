FROM golang:1.6.0 as builder

WORKDIR /go/src/github.com/oyamoh-brian/tv-service-database

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build  -o tv-service-database -a -installsuffix cgo main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/oyamoh-brian/tv-service-database/tv-service-database .

CMD ["./consignment-service"]
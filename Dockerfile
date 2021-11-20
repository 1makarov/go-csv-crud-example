FROM golang:1.17 AS builder

RUN go version

COPY . /github.com/1makarov/go-csv-crud-example/
WORKDIR /github.com/1makarov/go-csv-crud-example/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/1makarov/go-csv-crud-example/.bin/app .

CMD ["./app"]
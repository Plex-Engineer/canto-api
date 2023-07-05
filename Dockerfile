FROM golang:1.20

WORKDIR /go/src/canto-api

COPY . .

RUN go build -o build/canto-api main.go

CMD ["./build/canto-api"]